package query_test

import (
	"errors"
	"testing"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/storage"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type EdgeFakeStorage struct {
	vertices map[string]vertices.Vertex
	keys     []string
}

func (se EdgeFakeStorage) Fetch(ID string) (*vertices.Vertex, error) {
	if v, ok := se.vertices[ID]; ok {
		return &v, nil
	} else {
		return nil, errors.New("Not found")
	}
}

func (se EdgeFakeStorage) ForEach() enumerables.Iterator {
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

func NewEdgeFakeStorage(c ...*vertices.Vertex) storage.Storage {
	se := &FakeStorage{vertices: make(map[string]vertices.Vertex)}
	for _, v := range c {
		se.vertices[v.ID()] = *v
		se.keys = append(se.keys, v.ID())
	}
	return se
}

func Test_MatchEdge(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	frontier := query.Frontier{}
	fv := &query.FrontierVertex{Vertex: vertex}
	frontier = frontier.Append([]*query.FrontierVertex{fv}, 0)

	p := query.NewEdgePath(func() (item *query.Frontier, ok query.Traverse) {
		state = expressions.XORSwap(state)
		if state {
			return &frontier, query.Visiting
		}
		return &frontier, query.Failed

	}, NewEdgeFakeStorage(vertex))

	matches := p.Relationship(func(v *vertices.Edge, depth uint) (string, query.Traverse) {
		return "", query.Matched
	})

	matches.Iterate()
}
