package gp

import (
	"errors"
	"fmt"

	"github.com/Contra-Culture/report"
)

type (
	Parser interface {
		Parse(*GPRuneScanner) (*Node, error)
	}
	symbolParser struct {
		tags []string
		r    rune
	}
	stringParser struct {
		tags []string
		str  string
	}
	sequenceParser struct {
		tags     []string
		children []Parser
	}
	repeatParser struct {
		tags       []string
		repeatable Parser
	}
	variantParser struct {
		tags     []string
		variants []Parser
	}
	optionalParser struct {
		tags   []string
		option Parser
	}
	rangeParser struct {
		tags     []string
		min, max rune
	}
	anyOneOfRunesParser struct {
		tags  []string
		runes []rune
	}
	runeExceptParser struct {
		tags       []string
		exceptions []rune
	}
	proxyParser func() Parser
	univParser  struct {
		parsers map[string]Parser
	}
	UnivCfgr struct {
		univ         *univParser
		namesToCheck []string
		report       report.Node
	}
)

const TOP_NAME = ""

func T(p Parser, tags ...string) Parser {
	switch _p := p.(type) {
	case *symbolParser:
		_p.tags = tags
	case *stringParser:
		_p.tags = tags
	case *sequenceParser:
		_p.tags = tags
	case *repeatParser:
		_p.tags = tags
	case *variantParser:
		_p.tags = tags
	case *optionalParser:
		_p.tags = tags
	case *rangeParser:
		_p.tags = tags
	case *anyOneOfRunesParser:
		_p.tags = tags
	case *runeExceptParser:
		_p.tags = tags
	default:
		panic("not a parser") // should not occur
	}
	return p
}
func New(cfg func(*UnivCfgr)) (p Parser, err error) {
	uc := &UnivCfgr{
		univ: &univParser{
			parsers: map[string]Parser{},
		},
	}
	cfg(uc)
	ok := uc.check()
	if !ok {
		return
	}
	p = uc.univ
	return
}
func (c *UnivCfgr) check() (ok bool) {
	_, ok = c.univ.parsers[TOP_NAME]
	if !ok {
		c.report.Error("top-level parser is not specified")
		return false
	}
outer:
	for _, nameToCheck := range c.namesToCheck {
		for name := range c.univ.parsers {
			if nameToCheck == name {
				continue outer
			}
		}
		c.report.Error("wrong parser name \"%s\"", nameToCheck)
		return false
	}
	return true
}
func (c *UnivCfgr) Top(p Parser) {
	_, exists := c.univ.parsers[TOP_NAME]
	if exists {
		c.report.Error("top parser already specified")
		return
	}
	c.univ.parsers[TOP_NAME] = p
}
func (c *UnivCfgr) Define(n string, p Parser) {
	_, exists := c.univ.parsers[n]
	if exists {
		c.report.Error("parser \"%s\" already specified", n)
		return
	}
	c.univ.parsers[n] = p
}
func (c *UnivCfgr) Get(n string) Parser {
	c.namesToCheck = append(c.namesToCheck, n)
	u := c.univ
	return proxyParser(
		func() Parser {
			return u.parsers[n]
		})
}
func (p proxyParser) Parse(rs *GPRuneScanner) (*Node, error) {
	return (func() Parser)(p)().Parse(rs)
}
func (u *univParser) Parse(rs *GPRuneScanner) (*Node, error) {
	return u.parsers[TOP_NAME].Parse(rs)
}
func RuneExcept(rs ...rune) (p Parser) {
	return runeExceptParser{
		exceptions: rs,
	}
}
func (p runeExceptParser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		panic(err)
	}
	for _, notExpected := range p.exceptions {
		if r == notExpected {
			err = fmt.Errorf("wrong rune `%#U`", r)
			return
		}
	}
	n = &Node{
		tags:   p.tags,
		parsed: []rune{r},
	}
	return
}
func Digits() (rune, rune) {
	return 48, 58
}
func LowASCIIAlphabet() (rune, rune) {
	return 97, 122
}
func HighASCIIAlphabet() (rune, rune) {
	return 65, 90
}

func Range(min, max rune) (p Parser) {
	return rangeParser{
		min: min,
		max: max,
	}
}
func (p rangeParser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	i := int32(r)
	if i < p.min || i > p.max {
		_, err = rs.Unread(1)
		if err != nil {
			panic(err)
		}
		err = fmt.Errorf("wrong rune `%#U`", r)
		return
	}
	n = &Node{
		parsed: []rune{r},
		tags:   p.tags,
	}
	return
}
func Symbol(r rune) (p Parser) {
	p = &symbolParser{
		r: r,
	}
	return
}
func (p *symbolParser) Parse(rs *GPRuneScanner) (node *Node, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	if r == p.r {
		node = &Node{
			tags:   p.tags,
			parsed: []rune{r},
		}
		return
	}
	_, err = rs.Unread(1)
	if err != nil {
		panic(err) // should not occur
	}
	err = fmt.Errorf("unexpected utf-8 rune, expected: `%#U`, got: `%#U`", p.r, r)
	return
}
func Seq(ps ...Parser) (p Parser) {
	p = &sequenceParser{
		children: ps,
	}
	return
}
func (p *sequenceParser) Parse(rs *GPRuneScanner) (node *Node, err error) {
	node = &Node{
		tags:   p.tags,
		parsed: []rune{},
	}
	var childNode *Node
	for _, childParser := range p.children {
		childNode, err = childParser.Parse(rs)
		originErr := err
		if err != nil {
			if len(node.parsed) > 0 {
				_, err = rs.Unread(len(node.parsed))
				if err != nil {
					panic(err) // should not occur
				}
			}
			err = fmt.Errorf("parser failed: %w", originErr)
			return
		}
		if childNode != nil { // not optional node
			node.parsed = append(node.parsed, childNode.parsed...)
			node.children = append(node.children, childNode)
		}
	}
	return
}
func AnyOneOfRunes(rs ...rune) (p Parser) {
	return &anyOneOfRunesParser{
		runes: rs,
	}
}
func (p *anyOneOfRunesParser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	n = &Node{
		tags:   p.tags,
		parsed: []rune{},
	}
	var r rune
	r, _, err = rs.ReadOne()
	if err != nil {
		if err == ErrNoRuneToRead {
			return nil, nil
		}
		panic(err)
	}
	for _, expected := range p.runes {
		if expected == r {
			n.parsed = append(n.parsed, r)
			return
		}
	}
	_, err = rs.Unread(1)
	if err != nil {
		panic(err)
	}
	err = fmt.Errorf("unexpected utf-8 rune: `%#U`, expected: %#v", r, p.runes)
	return
}
func String(s string) (p Parser) {
	p = &stringParser{
		str: s,
	}
	return
}
func (p *stringParser) Parse(rr *GPRuneScanner) (node *Node, err error) {
	var r rune
	node = &Node{
		tags:   p.tags,
		parsed: []rune{},
	}
	for _, expected := range p.str {
		r, _, err = rr.ReadOne()
		if err != nil {
			if err == ErrNoRuneToRead {
				_, err = rr.Unread(len(node.parsed) + 1)
				if err != nil {
					panic(err) // should not occur
				}
				return
			}
			panic(err) // should not occur
		}
		if r != expected {
			_, err = rr.Unread(len(node.parsed) + 1)
			if err != nil {
				panic(err) // should not occur
			}
			err = fmt.Errorf("unexpected utf-8 rune, expected: `%#U`, got: `%#U`", expected, r)
			return
		}
		node.parsed = append(node.parsed, r)
	}
	return
}
func Repeat(rp Parser) (p Parser) {
	p = &repeatParser{
		repeatable: rp,
	}
	return
}
func (p *repeatParser) Parse(rr *GPRuneScanner) (node *Node, err error) {
	var child *Node
	node = &Node{
		tags:   p.tags,
		parsed: []rune{},
	}
	for {
		child, err = p.repeatable.Parse(rr)
		if err != nil {
			err = nil
			return
		}
		node.parsed = append(node.parsed, child.parsed...)
		node.children = append(node.children, child)
	}
}
func Variant(variants ...Parser) (p Parser) {
	p = &variantParser{
		variants: variants,
	}
	return
}
func (p *variantParser) Parse(rs *GPRuneScanner) (node *Node, err error) {
	for _, v := range p.variants {
		node, err = v.Parse(rs)
		if err != nil {
			err = nil
			continue
		}
		return
	}
	err = errors.New("no variant parsed")
	return
}
func Optional(op Parser) (p Parser) {
	p = &optionalParser{
		option: op,
	}
	return
}
func (p *optionalParser) Parse(rs *GPRuneScanner) (node *Node, err error) {
	node, err = p.option.Parse(rs)
	err = nil
	return
}
