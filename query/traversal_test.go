package query_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Traversal_Travers(t *testing.T) {

	vertex, _ := vertices.NewVertex()
	vertex.SetLabel("foo")

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	fetch := func(string) (*vertices.Vertex, error) {
		return vertex, nil
	}
	traversal := query.NewTraversal(fetch)

	frontier := query.Frontier{}
	frontier = frontier.Append([]*vertices.Vertex{vertex}, 0)

	state := false
	it := func() (item interface{}, ok bool) {
		state = XOR(state)
		return frontier, state
	}

	path := &query.Root{}

	edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}

	b := true
	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return b
		}
	}

	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) bool {
			return b
		}
	}
	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	path.SetNext(vertexPath)

	q := query.NewQuery(path, "")
	traversal.Travers(func() query.Iterator {
		return it
	}, q)

	if len(q.Results) != 1 {
		t.Errorf("Failed to match")
	}
}
