package cypherquerybuilder

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func init() {
	cypher.RegisterQueryBuilder(QueryBuilderType, cypher.QueryBuilderRegistry{
		NewFunc: func() cypher.QueryBuilder {
			return &KeyValueCyperQueryBuilder{}
		},
	})
}

const (
	// QueryBuilderType the underlying query builder cypher/keyvalue
	QueryBuilderType = "cypherkeyvalue"
)

type KeyValueCyperQueryBuilder struct {
}

func (s *KeyValueCyperQueryBuilder) Predicate(patterns []ast.Patn) []query.Predicate {
	result := make([]query.Predicate, 0)
	for _, patn := range patterns {
		result = append(result, ToPredicatePath(patn))
	}

	return result
}

// ToPredicatePath creates a PredicatePath out of the Patn
func ToPredicatePath(patn ast.Patn) query.Predicate {
	if vertex, ok := patn.(*ast.VertexPatn); ok {
		return ToPredicateVertexPath(vertex)
	}

	if edge, ok := patn.(*ast.EdgePatn); ok {
		return ToPredicateEdgePath(edge)
	}

	return nil
}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func ToPredicateVertexPath(patn *ast.VertexPatn) query.Predicate {
	//label := strings.ToLower(patn.Label)
	return func(uuid uuid.UUID, depth int) (string, query.Traverse) {
		// split := bytes.Split(kv.Key, US)

		// if bytes.Equal(split[1], Vertex) {
		// 	value, ok := kv.Value.Unmarshal().(string)
		// 	if ok && label != value {
		// 		return patn.Variable, Failed
		// 	}

		// 	return patn.Variable, Matched
		// }

		// if bytes.Equal(split[1], Properties) {
		// 	key := split[2]
		// 	property := string(key)
		// 	if value, ok := patn.Properties[property]; ok {
		// 		if value != kv.Value.Unmarshal() {
		// 			return patn.Variable, Failed
		// 		}
		// 	}

		// 	return patn.Variable, Matched
		// }

		return patn.Variable, query.Failed

	}
}

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func ToPredicateEdgePath(patn *ast.EdgePatn) query.Predicate {
	//label := strings.ToLower(patn.Body.Type)
	return func(uuid uuid.UUID, depth int) (string, query.Traverse) {
		// split := bytes.Split(kv.Key, US)

		// if bytes.Equal(split[1], Vertex) {
		// 	value, ok := kv.Value.Unmarshal().(string)
		// 	if ok && label != value {
		// 		return patn.Variable, Failed
		// 	}

		// 	return patn.Variable, Matched
		// }

		// if bytes.Equal(split[1], Properties) {
		// 	key := split[2]
		// 	property := string(key)
		// 	if value, ok := patn.Body.Properties[property]; ok {
		// 		if value != kv.Value.Unmarshal() {
		// 			return patn.Variable, Failed
		// 		}
		// 	}

		// 	return patn.Variable, Matched
		// }

		return patn.Variable, query.Failed

	}

	// relationshipType := strings.ToLower(patn.Body.Type)
	// pvp := query.PredicateEdgePath{PredicateEdge: func(v *graph.Edge, depth uint) (string, query.Traverse) {

	// 	if depth < patn.Body.LengthMinimum {
	// 		return patn.Body.Variable, query.Visiting
	// 	}

	// 	if depth > patn.Body.LengthMaximum {
	// 		return patn.Body.Variable, query.Failed
	// 	}

	// 	if relationshipType != emptyString {
	// 		if relationshipType != v.RelationshipType() {
	// 			return patn.Body.Variable, query.Failed
	// 		}
	// 	}

	// 	for key, value := range patn.Body.Properties {
	// 		if v.Property(key) != value {
	// 			return patn.Body.Variable, query.Failed
	// 		}
	// 	}

	// 	return patn.Body.Variable, query.Matching
	// }, Variable: patn.Variable}

	// return &pvp
}
