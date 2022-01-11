package gp

type (
	anyOfRunesRule struct {
		optional bool
		tags     []string
		runes    []rune
	}
)

func AnyOfRunes(rs ...rune) interface{} {
	return anyOfRunesRule{
		runes: rs,
	}
}
