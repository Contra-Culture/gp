package gp

import "fmt"

type (
	anyOneOfRunesParser struct {
		tags  []string
		runes []rune
	}
)

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
