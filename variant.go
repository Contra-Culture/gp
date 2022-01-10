package gp

import "errors"

type (
	variantParser struct {
		tags     []string
		variants []Parser
	}
)

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
