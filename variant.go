package gp

type (
	variantRule struct {
		optional bool
		tags     []string
		variants []interface{}
	}
)

func Variant(variants ...interface{}) interface{} {
	return &variantRule{
		variants: variants,
	}
}
