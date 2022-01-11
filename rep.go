package gp

type (
	repeatRule struct {
		optional   bool
		tags       []string
		repeatable interface{}
	}
)

func Repeat(r interface{}) interface{} {
	return repeatRule{
		repeatable: r,
	}
}
