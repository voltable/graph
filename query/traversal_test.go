package query_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/storage"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type FakeStorage struct {
	vertices map[string]vertices.Vertex
	f        map[string]vertices.Vertex

	keys []string
}

func (se FakeStorage) Fetch(ID string) (*vertices.Vertex, error) {
	if v, ok := se.f[ID]; ok {
		return &v, nil
	} else {
		return nil, errors.New("Not found")
	}
}

func (se FakeStorage) ForEach() enumerables.Iterator {
	position := 0
	length := len(se.keys)
	return func() (item interface{}, ok bool) {
		if position < length {
			key := se.keys[position]
			v := se.vertices[key]
			position = position + 1
			return &v, true
		}
		return nil, false
	}
}

func NewFakeStorage(vv []*vertices.Vertex, ff []*vertices.Vertex) storage.Storage {
	se := &FakeStorage{vertices: make(map[string]vertices.Vertex), f: make(map[string]vertices.Vertex)}
	for _, v := range ff {
		se.f[v.ID()] = *v
	}
	for _, v := range vv {
		se.vertices[v.ID()] = *v
		se.keys = append(se.keys, v.ID())
	}
	return se
}
func Test_Traversal_Travers(t *testing.T) {
	tests := []struct {
		b        []bool
		v        func() ([]*vertices.Vertex, []*vertices.Vertex)
		expected int
	}{
		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match All
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)

				arrFind = append(arrFind, vertexDirection)

				return arr, arrFind
			},
			expected: 1,
		},
		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match all
				arr = append(arr, true) // Match first edge
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)

				arrFind = append(arrFind, vertexDirection)

				return arr, arrFind
			},
			expected: 2,
		},
		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true)
				arr = append(arr, false)
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)

				arrFind = append(arrFind, vertexDirection)

				return arr, arrFind
			},
			expected: 0,
		},
		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match all
				arr = append(arr, true) // Match first edge
				arr = append(arr, true) // Match second node to first edge
				arr = append(arr, true) // Match
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)
				arrFind = append(arrFind, vertexDirection)

				vertexDirection2, _ := vertices.NewVertex()
				vertexDirection.AddDirectedEdge(vertexDirection2)
				arrFind = append(arrFind, vertexDirection2)

				return arr, arrFind
			},
			expected: 3,
		},

		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match all
				arr = append(arr, true) // Match first edge
				arr = append(arr, true) // Match second node to first edge
				arr = append(arr, true) // Match second node to second edge
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertex2, _ := vertices.NewVertex()
				vertex2.SetLabel("foo")
				arr = append(arr, vertex2)
				arrFind = append(arrFind, vertex2)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)

				arrFind = append(arrFind, vertexDirection)

				vertexDirection2, _ := vertices.NewVertex()
				vertexDirection.AddDirectedEdge(vertexDirection2)
				arrFind = append(arrFind, vertexDirection2)

				return arr, arrFind
			},
			expected: 3,
		},

		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match all
				arr = append(arr, true) // Match first edge
				arr = append(arr, true) // Match second node to first edge
				arr = append(arr, true) // Match second node to second edge
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertex2, _ := vertices.NewVertex()
				vertex2.SetLabel("foo")
				arr = append(arr, vertex2)
				arrFind = append(arrFind, vertex2)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)
				vertex2.AddDirectedEdge(vertexDirection)

				arrFind = append(arrFind, vertexDirection)

				vertexDirection2, _ := vertices.NewVertex()
				vertexDirection.AddDirectedEdge(vertexDirection2)
				arrFind = append(arrFind, vertexDirection2)

				return arr, arrFind
			},
			expected: 3,
		},

		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true) // Match all
				arr = append(arr, true) // Match first edge
				arr = append(arr, true) // Match second node to first edge
				arr = append(arr, true) // Match second node to second edge
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertex2, _ := vertices.NewVertex()
				vertex2.SetLabel("foo2")
				arr = append(arr, vertex2)
				arrFind = append(arrFind, vertex2)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)
				arrFind = append(arrFind, vertexDirection)

				vertex2Direction, _ := vertices.NewVertex()
				vertex2.AddDirectedEdge(vertex2Direction)

				arrFind = append(arrFind, vertex2Direction)

				vertexDirection2, _ := vertices.NewVertex()
				vertexDirection.AddDirectedEdge(vertexDirection2)
				arrFind = append(arrFind, vertexDirection2)

				vertex2Direction2, _ := vertices.NewVertex()
				vertex2Direction.AddDirectedEdge(vertex2Direction2)
				arrFind = append(arrFind, vertex2Direction2)

				return arr, arrFind
			},
			expected: 6,
		},
		{
			b: func() []bool {
				arr := make([]bool, 0, 0)
				arr = append(arr, true)  // Match all
				arr = append(arr, true)  // Match first edge
				arr = append(arr, true)  // Match second node to first edge
				arr = append(arr, false) // Fail to Match second node to second edge
				return arr
			}(),
			v: func() ([]*vertices.Vertex, []*vertices.Vertex) {
				arr := make([]*vertices.Vertex, 0, 0)
				arrFind := make([]*vertices.Vertex, 0, 0)

				vertex, _ := vertices.NewVertex()
				vertex.SetLabel("foo")
				arr = append(arr, vertex)
				arrFind = append(arrFind, vertex)

				vertex2, _ := vertices.NewVertex()
				vertex2.SetLabel("foo2")
				arr = append(arr, vertex2)
				arrFind = append(arrFind, vertex2)

				vertexDirection, _ := vertices.NewVertex()
				vertex.AddDirectedEdge(vertexDirection)
				arrFind = append(arrFind, vertexDirection)

				vertex2Direction, _ := vertices.NewVertex()
				vertex2.AddDirectedEdge(vertex2Direction)

				arrFind = append(arrFind, vertex2Direction)

				vertexDirection2, _ := vertices.NewVertex()
				vertexDirection.AddDirectedEdge(vertexDirection2)
				arrFind = append(arrFind, vertexDirection2)

				vertex2Direction2, _ := vertices.NewVertex()
				vertex2Direction.AddDirectedEdge(vertex2Direction2)
				arrFind = append(arrFind, vertex2Direction2)

				return arr, arrFind
			},
			expected: 0,
		},
	}
	for i, tt := range tests {
		arr, find := tt.v()
		traversal := query.NewTraversal(NewFakeStorage(arr, find))
		iteratorFrontier, _ := traversal.Travers(BuildIterator(arr), BuildPath(tt.b))

		results := ToVertices(iteratorFrontier)
		fmt.Printf("%d", len(results))
		if len(results) != tt.expected {
			t.Errorf("%d. Failed to match expect %+v got %+v", i, tt.expected, len(results))
		}
	}
}

func BuildIterator(arr []*vertices.Vertex) func() (item interface{}, ok bool) {
	position := 0
	length := len(arr)
	return func() (item interface{}, ok bool) {
		if position < length {
			v := arr[position]
			position++
			return v, true
		}
		return nil, false
	}
}

func BuildPath(arr []bool) query.Path {
	path, _ := NewPath()
	last := path
	for i, a := range arr {
		if i%2 == 0 {
			vertexPath := &query.PredicateVertexPath{PredicateVertex: ToPredicateVertex(a)(nil)}
			last.SetNext(vertexPath)
			last = vertexPath
		}
		if i%2 == 1 {
			edgePath := &query.PredicateEdgePath{PredicateEdge: ToPredicateEdge(a)(nil)}
			last.SetNext(edgePath)
			last = edgePath
		}
	}

	return path
}

func ToPredicateVertex(b bool) func(*ir.VertexPatn) query.PredicateVertex {
	toPredicateVertex := func(*ir.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) (string, query.Traverse) {
			if b {
				return "", query.Matched
			} else {
				return "", query.Failed

			}
		}
	}

	return toPredicateVertex
}

func ToPredicateEdge(b bool) func(patn *ir.EdgePatn) query.PredicateEdge {
	toPredicateEdge := func(patn *ir.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge, depth uint) (string, query.Traverse) {
			if b {
				return "", query.Matched
			} else {
				return "", query.Failed

			}
		}
	}
	return toPredicateEdge
}

func ToVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for frontier, ok := i(); ok != query.Failed; frontier, ok = i() {
		if frontier.Len() > 0 {
			vertices, _, _ := frontier.Pop()
			for _, v := range vertices {
				results = append(results, v)
			}
		}
	}
	return results
}
