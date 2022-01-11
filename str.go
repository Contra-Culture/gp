package gp

type (
	stringRule struct {
		optional bool
		tags     []string
		str      string
	}
)

func String(s string) interface{} {
	return stringRule{
		str: s,
	}
}
