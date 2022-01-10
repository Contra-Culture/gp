package gp

type (
	optionalParser struct {
		tags   []string
		option Parser
	}
)

func Optional(op Parser) (p Parser) {
	p = &optionalParser{
		option: op,
	}
	return
}
func (p *optionalParser) Parse(rs *GPRuneScanner) (node *Node, err error) {
	node, err = p.option.Parse(rs)
	err = nil
	return
}
