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
