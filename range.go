package gp

import "fmt"

type (
	rangeParser struct {
		tags     []string
		min, max rune
	}
)

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
