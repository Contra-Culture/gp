package gp

type (
	Parser struct {
		syntax *syntax
	}
)

const TOP_NAME = ""

func New(cfg func(*SyntaxCfgr)) (p *Parser, err error) {
	uc := &SyntaxCfgr{
		syntax: &syntax{
			rules: map[string]interface{}{},
		},
	}
	cfg(uc)
	ok := uc.check()
	if !ok {
		return
	}
	u := uc.syntax
	p = &Parser{
		syntax: u,
	}
	return
}
func (p *Parser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	u := p.syntax
	tr := u.rules[u.top]
	return
}
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
