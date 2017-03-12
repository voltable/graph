package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath represents the Vertex part of a Path
type VertexPath struct {
	Iterate func() Iterator
	next    Path
	fetch   func(string) (*vertices.Vertex, error)
}

// Match returns all Verteces matching the predicate
func (t *VertexPath) Match(predicate func(v *vertices.Vertex) bool) *EdgePath {
	if predicate == nil {
		predicate = AllVertices()
	}

	return &EdgePath{
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if v, is := item.(*vertices.Vertex); is {
						if predicate(v) {
							for _, e := range v.Edges() {
								return e, true
							}
						}
					}
				}
				return
			}
		},
	}
}

func (t *VertexPath) MatchAll() *EdgePath {
	return t.Match(nil)
}

// ToSlice iterates over a collection and saves the results in the slice
func (t *VertexPath) ToSlice(predicate func(v *vertices.Vertex) bool) []*vertices.Vertex {
	if predicate == nil {
		predicate = AllVertices()
	}

	slice := []*vertices.Vertex{}
	next := t.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if e, is := item.(*vertices.Vertex); is {
			if predicate(e) {
				slice = append(slice, e)
			}
		}
	}
	return slice
}

func (t *VertexPath) ToSliceAll() []*vertices.Vertex {
	return t.ToSlice(nil)
}

// AllVertices matches all Vertexes.
func AllVertices() func(v *vertices.Vertex) bool {
	return func(v *vertices.Vertex) bool {
		return true
	}
}
