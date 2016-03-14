package lang

func Equal(v, w interface{}) bool {
	if v, ok := v.(Equatable); ok {
		return v.Equals(w)
	}

	return v == w
}

func EqualDeep(v, w interface{}) bool {
	if v, ok := v.(EquatableDeep); ok {
		return v.EqualsDeep(w)
	}

	return v == w
}

type Equatable interface {
	Equals(v interface{}) bool
}

type EquatableDeep interface {
	EqualsDeep(v interface{}) bool
}
