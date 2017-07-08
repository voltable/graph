package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"

	"github.com/RossMerr/Caudex.Graph/query"
)

// ToVertexPath converts a cypher.Stmt to a VertexPath to keep it all abstracted
func ToVertexPath(stmt ast.Stmt) (*query.VertexPath, error) {

	var pv []query.PredicateVertex
	var pe []query.PredicateEdge

	if b, ok := stmt.(*ast.CreateStmt); ok {
		pattern := b.Pattern
		for pattern != nil {
			if v, ok := b.Pattern.(*ast.VertexPatn); ok {
				pv = append(pv, v.ToPredicateVertex())
				pattern = v.Edge
			} else if e, ok := b.Pattern.(*ast.EdgePatn); ok {
				pe = append(pe, e.ToPredicateEdge())
				pattern = e.Vertex
			}
		}
	}

	return nil, nil
}
