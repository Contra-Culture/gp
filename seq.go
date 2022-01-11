package gp

type (
	sequenceRule struct {
		optional bool
		tags     []string
		children []interface{}
	}
)

func Seq(ps ...interface{}) interface{} {
	return sequenceRule{
		children: ps,
	}
}
