package gp

type (
	variantRule struct {
		tags     []string
		variants []interface{}
	}
)

func Variant(variants ...interface{}) interface{} {
	return &variantRule{
		variants: variants,
	}
}
