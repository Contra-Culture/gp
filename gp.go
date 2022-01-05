package gp

import (
	"fmt"
	"io"
)

type (
	ASTNode struct {
		parser   Parser
		parsed   []rune
		children []*ASTNode
	}
	Parser interface {
		Name() string
		Kind() string
		Parse(io.RuneScanner) (*ASTNode, error)
	}
	symbolParser struct {
		r    rune
		name string
	}
	stringParser struct {
		str  string
		name string
	}
	sequenceParser struct {
		children []Parser
		name     string
	}
	repeatParser struct {
		repeatable Parser
		name       string
	}
	variantParser struct {
		variants []Parser
		name     string
	}
	optionalParser struct {
		option Parser
	}
	digitParser         struct{}
	lowAlphaParser      struct{}
	highAlphaParser     struct{}
	anyOneOfRunesParser struct {
		runes []rune
		name  string
	}
)

func Alpha() (p Parser) {
	return Variant("alpha", HighAlpha(), LowAlpha())
}
func AlphaDigit() (p Parser) {
	return Variant("alphaDigit", Alpha(), Digit())
}
func HighAlpha() (p Parser) {
	return highAlphaParser{}
}
func (_ highAlphaParser) Name() string {
	return "highAlpha"
}
func (_ highAlphaParser) Kind() string {
	return "variant"
}
func (p highAlphaParser) Parse(rs io.RuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadRune()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 65 || i > 90 {
		err = rs.UnreadRune()
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
func (_ lowAlphaParser) Name() string {
	return "lowAlpha"
}
func (_ lowAlphaParser) Kind() string {
	return "variant"
}
func (p lowAlphaParser) Parse(rs io.RuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadRune()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 97 || i > 122 {
		err = rs.UnreadRune()
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
func (_ digitParser) Name() string {
	return "digit"
}
func (_ digitParser) Kind() string {
	return "variant"
}
func (p digitParser) Parse(rs io.RuneScanner) (n *ASTNode, err error) {
	r, _, err := rs.ReadRune()
	if err != nil {
		return
	}
	i := int32(r)
	if i < 48 || i > 58 {
		err = rs.UnreadRune()
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
func (n *ASTNode) ParserName() string {
	return n.parser.Name()
}
func (n *ASTNode) ParserKind() string {
	return n.parser.Kind()
}
func (n *ASTNode) Children() []*ASTNode {
	return n.children
}
func Symbol(n string, r rune) (p Parser) {
	p = &symbolParser{
		r:    r,
		name: n,
	}
	return
}
func (p *symbolParser) Name() string {
	return p.name
}
func (p *symbolParser) Kind() string {
	return "symbol"
}
func (p *symbolParser) Parse(rr io.RuneScanner) (node *ASTNode, err error) {
	expected := rune(p.r)
	r, _, err := rr.ReadRune()
	if err != nil {
		return
	}
	if r != expected {
		err = rr.UnreadRune()
		if err != nil {
			panic(err) // should not occur
		}
		err = fmt.Errorf("unexpected utf-8 rune, expected: `%#U`, got: `%#U`", expected, r)
		return
	}
	node = &ASTNode{
		parser: p,
		parsed: []rune{r},
	}
	return
}
func Seq(n string, ps ...Parser) (p Parser) {
	p = &sequenceParser{
		children: ps,
		name:     n,
	}
	return
}
func (p *sequenceParser) Name() string {
	return p.name
}
func (p *sequenceParser) Kind() string {
	return "sequence"
}
func (p *sequenceParser) Parse(rs io.RuneScanner) (node *ASTNode, err error) {
	node = &ASTNode{
		parser: p,
	}
	var childNode *ASTNode
	for _, childParser := range p.children {
		childNode, err = childParser.Parse(rs)
		if err != nil {
			for range node.parsed {
				err = rs.UnreadRune()
				if err != nil {
					panic(err) // should not occur
				}
			}
			err = fmt.Errorf("children parser `%s` failed: %w", childParser.Name(), err)
			return
		}
		node.parsed = append(node.parsed, childNode.parsed...)
		node.children = append(node.children, childNode)
	}
	return
}
func AnyOneOfRunes(n string, rs ...rune) (p Parser) {
	return &anyOneOfRunesParser{
		runes: rs,
		name:  n,
	}
}
func (p *anyOneOfRunesParser) Name() string {
	return p.name
}
func (p *anyOneOfRunesParser) Kind() string {
	return "variant"
}
func (p *anyOneOfRunesParser) Parse(rs io.RuneScanner) (n *ASTNode, err error) {
	n = &ASTNode{
		parser: p,
	}
	var r rune
	r, _, err = rs.ReadRune()
	if err != nil {
		panic(r)
	}
	for _, expected := range p.runes {
		if expected == r {
			n.parsed = append(n.parsed, r)
			return
		}

	}
	err = rs.UnreadRune()
	if err != nil {
		panic(err)
	}
	err = fmt.Errorf("unexpected utf-8 rune: `%#U`", r)
	return
}
func String(s string) (p Parser) {
	p = &stringParser{
		str:  s,
		name: fmt.Sprintf("\"%s\"", s),
	}
	return
}
func (p *stringParser) Name() string {
	return p.name
}
func (p *stringParser) Kind() string {
	return "sequence"
}
func (p *stringParser) Parse(rr io.RuneScanner) (node *ASTNode, err error) {
	var r rune
	node = &ASTNode{
		parser: p,
	}
	for i, expected := range p.str {
		r, _, err = rr.ReadRune()
		if err != nil {
			for ; i != 0; i-- {
				err = rr.UnreadRune()
				if err != nil {
					panic(err) // should not occur
				}
			}
			return
		}
		if r != expected {
			err = fmt.Errorf("unexpected utf-8 rune, expected: `%#U`, got: `%#U`", expected, r)
			return
		}
		node.parsed = append(node.parsed, r)
	}
	return
}
func Repeat(n string, rp Parser) (p Parser) {
	p = &repeatParser{
		repeatable: rp,
		name:       n,
	}
	return
}
func (p *repeatParser) Name() string {
	return p.name
}
func (p *repeatParser) Kind() string {
	return "repeat"
}
func (p *repeatParser) Parse(rr io.RuneScanner) (node *ASTNode, err error) {
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
func Variant(n string, variants ...Parser) (p Parser) {
	p = &variantParser{
		variants: variants,
		name:     n,
	}
	return
}
func (p *variantParser) Name() string {
	return p.name
}
func (p *variantParser) Kind() string {
	return "variant"
}
func (p *variantParser) Parse(rs io.RuneScanner) (node *ASTNode, err error) {
	node = &ASTNode{
		parser: p,
	}
	var child *ASTNode
	for _, v := range p.variants {
		child, err = v.Parse(rs)
		if err != nil {
			continue
		}
		node.parsed = append(node.parsed, child.parsed...)
		node.children = append(node.children, child)
		return
	}
	err = fmt.Errorf("no `%s` variant parsed", p.name)
	return
}
func Optional(op Parser) (p Parser) {
	p = &optionalParser{
		option: op,
	}
	return
}
func (p *optionalParser) Name() string {
	return fmt.Sprintf("[%s]", p.option.Name())
}
func (p *optionalParser) Kind() string {
	return "optional"
}
func (p *optionalParser) Parse(rs io.RuneScanner) (node *ASTNode, err error) {
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
