package gp

func Optional(r interface{}) interface{} {
	switch _r := r.(type) {
	case symbolRule:
		_r.optional = true
	case stringRule:
		_r.optional = true
	case sequenceRule:
		_r.optional = true
	case repeatRule:
		_r.optional = true
	case variantRule:
		_r.optional = true
	case rangeRule:
		_r.optional = true
	case anyOfRunesRule:
		_r.optional = true
	case runeExceptRule:
		_r.optional = true
	default:
		panic("not a rule") // should not occur
	}
	return r
}
