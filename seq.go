package gp

import "fmt"

type (
	sequenceParser struct {
		tags     []string
		children []Parser
	}
)

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
