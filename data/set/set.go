package set

import "github.com/willfaught/lang/data/maps"

type Set interface {
	Add(v interface{})

	Contains(v interface{}) bool

	Elements() []interface{}

	Len() int

	Remove(v interface{})
}

var _ Set = set{}

func Difference(s, t Set) {
	for _, e := range t.Elements() {
		if s.Contains(e) {
			s.Remove(e)
		}
	}
}

func Intersect(s, t Set) {
	for _, e := range s.Elements() {
		if !t.Contains(e) {
			s.Remove(e)
		}
	}

	for _, e := range t.Elements() {
		if !s.Contains(e) {
			t.Remove(e)
		}
	}
}

func Subset(s, t Set) bool {
	for _, e := range s.Elements() {
		if !t.Contains(e) {
			return false
		}
	}

	return true
}

func Union(s, t Set) {
	for _, e := range t.Elements() {
		s.Add(e)
	}
}

type set struct {
	m maps.Map
}

func NewSet() Set {
	return set{m: maps.MapInterfaceStruct{}}
}

func NewSetFor(m maps.Map) Set {
	return set{m: m}
}

func (s set) Add(v interface{}) {
	s.m.Set(v, struct{}{})
}

func (s set) Contains(v interface{}) bool {
	var _, ok = s.m.GetCheck(v)

	return ok
}

func (s set) Elements() []interface{} {
	return s.m.Keys()
}

func (s set) Len() int {
	return s.m.Len()
}

func (s set) Remove(v interface{}) {
	s.m.Delete(v)
}
