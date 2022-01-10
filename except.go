package gp

import "fmt"

type (
	runeExceptParser struct {
		tags       []string
		exceptions []rune
	}
)

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
