package gp

type (
	proxyRule struct {
		optional bool
		name     string
	}
)

func Use(n string) interface{} {
	return proxyRule{
		name: n,
	}
}
