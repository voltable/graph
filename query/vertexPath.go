package query

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/query"
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
					if frontier, is := item.(query.Frontier); is {
						sort.Sort(frontier)
						path := frontier.peek()
						v := path.Vertices[len(p.Vertices)-1]
						if predicate(v) {
							return frontier, true
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

// MatchToSlice iterates over a collection and saves the results in the slice
func (t *VertexPath) MatchToSlice(predicate func(v *vertices.Vertex) bool) []*vertices.Vertex {
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

func (t *VertexPath) ToSlice() []*vertices.Vertex {
	return t.MatchToSlice(nil)
}

// AllVertices matches all Vertexes.
func AllVertices() func(v *vertices.Vertex) bool {
	return func(v *vertices.Vertex) bool {
		return true
	}
}
