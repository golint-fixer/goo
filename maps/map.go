package maps

type Map interface {
	Delete(k interface{})

	Get(k interface{}) interface{}

	GetCheck(k interface{}) (interface{}, bool)

	KeyValues() [][2]interface{}

	Keys() []interface{}

	Len() int

	Make(c int) Map

	Set(k, v interface{})
}
