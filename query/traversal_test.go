package query_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/storage"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type FakeStorage struct {
	vertex *vertices.Vertex
}

func (s FakeStorage) Fetch() func(string) (*vertices.Vertex, error) {
	return func(string) (*vertices.Vertex, error) {
		return s.vertex, nil
	}
}

func (s FakeStorage) ForEach() enumerables.Iterator {
	return func() (item interface{}, ok bool) {

		return nil, false
	}
}

func NewFakeStorage(v *vertices.Vertex) storage.Storage {
	return &FakeStorage{vertex: v}
}
func Test_Traversal_Travers(t *testing.T) {

	vertex, _ := vertices.NewVertex()
	vertex.SetLabel("foo")

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	traversal := query.NewTraversal(NewFakeStorage(vertex))

	frontier := query.Frontier{}
	fv := &query.FrontierVertex{Vertex: vertex}
	frontier = frontier.Append([]*query.FrontierVertex{fv}, 0)

	state := false

	path, _ := NewPath()

	edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}

	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return true
		}
	}

	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) (string, bool) {
			return "", false
		}
	}
	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	path.SetNext(vertexPath)

	iteratorFrontier := traversal.Travers(func() (item *query.Frontier, ok bool) {
		state = expressions.XORSwap(state)
		return &frontier, state
	}, path)

	results := ToVertices(iteratorFrontier)

	if len(results) != 1 {
		t.Errorf("Failed to match")
	}
}

func ToVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for frontier, ok := i(); ok; frontier, ok = i() {
		results = append(results, *frontier.Peek())
	}
	return results
}
