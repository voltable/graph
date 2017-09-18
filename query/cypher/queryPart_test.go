package cypher_test

import (
	"testing"
)

func Test_ToQueryPath(t *testing.T) {
	// edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	// vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}
	// var b bool

	// toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
	// 	return func(v *vertices.Vertex) bool {
	// 		return b
	// 	}
	// }

	// toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
	// 	return func(e *vertices.Edge) bool {
	// 		return b
	// 	}
	// }

	// want := &cypher.Root{}
	// vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	// vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	// want.SetNext(vertexPath)

	// engine := cypher.NewEngine(NewFakeStorage())
	// engine.ToPredicateEdge = toPredicateEdge
	// engine.ToPredicateVertex = toPredicateVertex

	// parts, _ := engine.ToQueryPart(&ast.MatchStmt{Pattern: vertexPatn})
	// got := parts[0].Path
	// v, _ := got.Next().(query.VertexNext)
	// if v == nil {
	// 	t.Errorf("VertexNext")
	// }

	// pv, _ := v.(*query.PredicateVertexPath)

	// if pv == nil {
	// 	t.Errorf("PredicateVertexPath")
	// }

	// e, _ := pv.Next().(query.EdgeNext)
	// if e == nil {
	// 	t.Errorf("EdgeNext")
	// }

	// pe, _ := e.(*query.PredicateEdgePath)
	// if pe == nil {
	// 	t.Errorf("PredicateEdgePath")
	// }

	// if pe.Length().Minimum != 2 {
	// 	t.Errorf("Minimum")
	// }

	// if pe.Length().Maximum != 5 {
	// 	t.Errorf("Maximum")
	// }

	// last, _ := pe.Next().(query.VertexNext)
	// if last != nil {
	// 	t.Errorf("VertexNext")
	// }
}
