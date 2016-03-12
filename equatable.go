package lang

func Equal(v, w interface{}) bool {
	if v, ok := v.(Equatable); ok {
		return v.Equals(w)
	}

	return v == w
}

type Equatable interface {
	Equals(v interface{}) bool
}
