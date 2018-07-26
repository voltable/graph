package ir

import (
	"bytes"
	"strings"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

// ToPredicatePath creates a PredicatePath out of the Patn
func ToPredicatePath(patn ast.Patn) *query.PredicatePath {
	if vertex, ok := patn.(*ast.VertexPatn); ok {
		return ToPredicateVertexPath(vertex)
	}

	if edge, ok := patn.(*ast.EdgePatn); ok {
		return ToPredicateEdgePath(edge)
	}

	return nil
}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func ToPredicateVertexPath(patn *ast.VertexPatn) *query.PredicatePath {
	label := strings.ToLower(patn.Label)
	pvp := query.PredicatePath{Predicate: func(kv *keyvalue.KeyValue) (string, query.Traverse) {
		split := bytes.Split(kv.Key, keyvalue.US)

		if bytes.Equal(split[1], keyvalue.Vertex) {
			value, ok := kv.Value.Unmarshal().(string)
			if ok && label != value {
				return patn.Variable, query.Failed
			}

			return patn.Variable, query.Matched
		}

		if bytes.Equal(split[1], keyvalue.Properties) {
			key := split[2]
			property := string(key)
			if value, ok := patn.Properties[property]; ok {
				if value != kv.Value.Unmarshal() {
					return patn.Variable, query.Failed
				}
			}

			return patn.Variable, query.Matched
		}

		return patn.Variable, query.Failed

	}, Variable: patn.Variable}

	return &pvp
}

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func ToPredicateEdgePath(patn *ast.EdgePatn) *query.PredicatePath {
	label := strings.ToLower(patn.Body.Type)
	pvp := query.PredicatePath{Predicate: func(kv *keyvalue.KeyValue) (string, query.Traverse) {
		split := bytes.Split(kv.Key, keyvalue.US)

		if bytes.Equal(split[1], keyvalue.Vertex) {
			value, ok := kv.Value.Unmarshal().(string)
			if ok && label != value {
				return patn.Variable, query.Failed
			}

			return patn.Variable, query.Matched
		}

		if bytes.Equal(split[1], keyvalue.Properties) {
			key := split[2]
			property := string(key)
			if value, ok := patn.Body.Properties[property]; ok {
				if value != kv.Value.Unmarshal() {
					return patn.Variable, query.Failed
				}
			}

			return patn.Variable, query.Matched
		}

		return patn.Variable, query.Failed

	}, Variable: patn.Variable}

	return &pvp

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
