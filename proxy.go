package gp

type (
	proxyParser string
)

func (p proxyParser) Parse(rs *GPRuneScanner) (*Node, error) {
	return (func() Parser)(p)().Parse(rs)
}
