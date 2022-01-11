package gp

type (
	rangeRule struct {
		optional bool
		tags     []string
		min, max rune
	}
)

func Range(min, max rune) interface{} {
	return rangeRule{
		min: min,
		max: max,
	}
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
