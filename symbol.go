package gp

import "fmt"

type (
	symbolParser struct {
		tags []string
		r    rune
	}
)

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
