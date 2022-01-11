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
		panic("not a parser") // should not occur
	}
	return r
}
func Digits() (rune, rune) {
	return 48, 58
}
func LowASCIIAlphabet() (rune, rune) {
	return 97, 122
}
func HighASCIIAlphabet() (rune, rune) {
	return 65, 90
}
func (p *Parser) Parse(rs *GPRuneScanner) (n *Node, err error) {
	return
}
