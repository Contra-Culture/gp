package gp

import (
	"errors"
	"fmt"

	"github.com/Contra-Culture/report"
)

type (
	ASTNode struct {
		parser   Parser
		parsed   []rune
		children []*ASTNode
	}
	Parser interface {
		Parse(*GPRuneScanner) (*ASTNode, error)
	}
	symbolParser struct {
		r rune
	}
	stringParser struct {
		str string
	}
	sequenceParser struct {
		children []Parser
	}
	repeatParser struct {
		repeatable Parser
	}
	variantParser struct {
		variants []Parser
	}
	optionalParser struct {
		option Parser
	}
	digitParser         struct{}
	lowAlphaParser      struct{}
	highAlphaParser     struct{}
	anyOneOfRunesParser struct {
		runes []rune
	}
	runeExceptParser struct {
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
func (p proxyParser) Parse(rs *GPRuneScanner) (*ASTNode, error) {
	return (func() Parser)(p)().Parse(rs)
}
func (u *univParser) Parse(rs *GPRuneScanner) (*ASTNode, error) {
	return u.parsers[TOP_NAME].Parse(rs)
}
func RuneExcept(rs ...rune) (p Parser) {
	return runeExceptParser{
		exceptions: rs,
	}
}
func (p runeExceptParser) Parse(rs *GPRuneScanner) (n *ASTNode, err error) {
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
	n = &ASTNode{
		parser: p,
		parsed: []rune{r},
	}
	return
}
func Alpha() (p Parser) {
	return Variant(HighAlpha(), LowAlpha())
}
func AlphaDigit() (p Parser) {
	return Variant(Alpha(), Digit())
}
func HighAlpha() (p Parser) {
	return highAlphaParser{}
}
func (p highAlphaParser) Parse(rs *GPRuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 65 || i > 90 {
		_, err = rs.Unread(1)
		if err != nil {
			panic(err)
		}
		err = fmt.Errorf("wrong rune `%#U`", r)
		return
	}
	n = &ASTNode{
		parser: p,
		parsed: []rune{r},
	}
	return
}
func LowAlpha() (p Parser) {
	return lowAlphaParser{}
}
func (p lowAlphaParser) Parse(rs *GPRuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 97 || i > 122 {
		_, err = rs.Unread(1)
		if err != nil {
			panic(err)
		}
		err = fmt.Errorf("wrong rune `%#U`", r)
		return
	}
	n = &ASTNode{
		parser: p,
		parsed: []rune{r},
	}
	return
}
func Digit() (p Parser) {
	return digitParser{}
}
func (p digitParser) Parse(rs *GPRuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 48 || i > 58 {
		_, err = rs.Unread(1)
		if err != nil {
			panic(err)
		}
		err = fmt.Errorf("wrong rune `%#U`", r)
		return
	}
	n = &ASTNode{
		parser: p,
		parsed: []rune{r},
	}
	return
}
func (n *ASTNode) Parsed() []rune {
	return n.parsed
}
func (n *ASTNode) Children() []*ASTNode {
	return n.children
}
func Symbol(r rune) (p Parser) {
	p = &symbolParser{
		r: r,
	}
	return
}
func (p *symbolParser) Parse(rs *GPRuneScanner) (node *ASTNode, err error) {
	r, _, err := rs.ReadOne()
	if err != nil {
		return
	}
	if r == p.r {
		node = &ASTNode{
			parser: p,
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
func (p *sequenceParser) Parse(rs *GPRuneScanner) (node *ASTNode, err error) {
	node = &ASTNode{
		parser: p,
	}
	var childNode *ASTNode
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
		node.parsed = append(node.parsed, childNode.parsed...)
		node.children = append(node.children, childNode)
	}
	return
}
func AnyOneOfRunes(rs ...rune) (p Parser) {
	return &anyOneOfRunesParser{
		runes: rs,
	}
}
func (p *anyOneOfRunesParser) Parse(rs *GPRuneScanner) (n *ASTNode, err error) {
	n = &ASTNode{
		parser: p,
	}
	var r rune
	r, _, err = rs.ReadOne()
	if err != nil {
		panic(r)
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
	err = fmt.Errorf("unexpected utf-8 rune: `%#U`", r)
	return
}
func String(s string) (p Parser) {
	p = &stringParser{
		str: s,
	}
	return
}
func (p *stringParser) Parse(rr *GPRuneScanner) (node *ASTNode, err error) {
	var r rune
	node = &ASTNode{
		parser: p,
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
func (p *repeatParser) Parse(rr *GPRuneScanner) (node *ASTNode, err error) {
	var child *ASTNode
	node = &ASTNode{
		parser: p,
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
func (p *variantParser) Parse(rs *GPRuneScanner) (node *ASTNode, err error) {
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
func (p *optionalParser) Parse(rs *GPRuneScanner) (node *ASTNode, err error) {
	node = &ASTNode{
		parser: p,
	}
	child, err := p.option.Parse(rs)
	if err != nil {
		err = nil
		return
	}
	node.children = append(node.children, child)
	node.parsed = append(node.parsed, child.parsed...)
	return
}
