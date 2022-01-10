package gp

import "fmt"

type (
	stringParser struct {
		tags []string
		str  string
	}
)

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
