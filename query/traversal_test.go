package query_test

// func Test_Traversal_Travers(t *testing.T) {

// 	fetch := func(string) (*vertices.Vertex, error) {
// 		return nil, nil
// 	}
// 	traversal := query.NewTraversal(fetch)

// 	it := func() (item interface{}, ok bool) {
// 		//state = XOR(state)
// 		//return frontier, state
// 		return nil, true
// 	}

// 	path := &query.Root{}

// 	edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
// 	vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}
// 	var b bool

// 	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
// 		return func(v *vertices.Vertex) bool {
// 			return b
// 		}
// 	}

// 	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
// 		return func(e *vertices.Edge) bool {
// 			return b
// 		}
// 	}
// 	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
// 	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
// 	path.SetNext(vertexPath)

// 	q := query.NewQuery(path, "")
// 	err := traversal.Travers(func() query.Iterator {
// 		return it
// 	}, q)

// 	if err != nil {
// 		t.Errorf("Travers")
// 	}
// }
