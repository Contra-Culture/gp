package gp

type (
	Parser struct {
		universe universe
	}
)

const TOP_NAME = ""

func T(r interface{}, tags ...string) interface{} {
	switch _r := r.(type) {
	case symbolRule:
		_r.tags = tags
	case stringRule:
		_r.tags = tags
	case sequenceRule:
		_r.tags = tags
	case repeatRule:
		_r.tags = tags
	case variantRule:
		_r.tags = tags
	case rangeRule:
		_r.tags = tags
	case anyOfRunesRule:
		_r.tags = tags
	case runeExceptRule:
		_r.tags = tags
	default:
		panic("not a rule") // should not occur
	}
	return r
}

func (p *Parser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	return
}
