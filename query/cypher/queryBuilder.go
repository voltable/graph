package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"

	"github.com/RossMerr/Caudex.Graph/query"
)

// ToVertexPath converts a cypher.Stmt to a VertexPath to keep it all abstracted
func ToVertexPath(stmt ast.Stmt) (*query.VertexPath, error) {

	var s []query.PredicateVertex

	if b, ok := stmt.(*ast.CreateStmt); ok {
		if v, ok := b.Pattern.(*ast.VertexPatn); ok {
			// todo need to iterate over this
			s = append(s, buildPredicateVertex(v))

		}
	}

	return nil, nil
}

func buildPredicateVertex(patn *ast.VertexPatn) query.PredicateVertex {
	return func(v *vertices.Vertex) bool {
		if patn.Label != v.Label() {
			return false
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return false
			}
		}

		return true
	}
}
