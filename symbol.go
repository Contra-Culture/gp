package gp

type (
	symbolRule struct {
		optional bool
		tags     []string
		r        rune
	}
)

func Symbol(r rune) interface{} {
	return symbolRule{
		r: r,
	}
}
