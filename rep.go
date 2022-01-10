package gp

type (
	repeatParser struct {
		tags       []string
		repeatable Parser
	}
)

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
