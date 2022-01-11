package gp

type (
	symbolRule struct {
		tags []string
		r    rune
	}
)

func Symbol(r rune) interface{} {
	return symbolRule{
		r: r,
	}
}
