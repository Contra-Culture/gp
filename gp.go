package gp

const TOP_NAME = ""

func T(p Parser, tags ...string) Parser {
	switch _p := p.(type) {
	case *symbolParser:
		_p.tags = tags
	case *stringParser:
		_p.tags = tags
	case *sequenceParser:
		_p.tags = tags
	case *repeatParser:
		_p.tags = tags
	case *variantParser:
		_p.tags = tags
	case *optionalParser:
		_p.tags = tags
	case *rangeParser:
		_p.tags = tags
	case *anyOneOfRunesParser:
		_p.tags = tags
	case *runeExceptParser:
		_p.tags = tags
	default:
		panic("not a parser") // should not occur
	}
	return p
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
