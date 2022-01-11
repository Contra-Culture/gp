package gp

type (
	runeExceptRule struct {
		optional   bool
		tags       []string
		exceptions []rune
	}
)

func RuneExcept(rs ...rune) interface{} {
	return runeExceptRule{
		exceptions: rs,
	}
}
