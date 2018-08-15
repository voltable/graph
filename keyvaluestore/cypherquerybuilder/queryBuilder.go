package cypherquerybuilder

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/keyvaluestore"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func init() {
	cypher.RegisterQueryBuilder(QueryBuilderType, cypher.QueryBuilderRegistry{
		NewFunc: newKeyValueCyperQueryBuilder,
	})
}

const (
	// QueryBuilderType the underlying query builder cypher/keyvalue
	QueryBuilderType = "cypherkeyvalue"
)

var (
	errNoPattern      = errors.New("No pattern provided")
	errUnknownPattern = errors.New("Unknown pattern")
)

type KeyValueCyperQueryBuilder struct {
	storage query.Storage
}

func newKeyValueCyperQueryBuilder(storage query.Storage) (cypher.QueryBuilder, error) {
	return NewKeyValueCyperQueryBuilder(storage), nil
}

func NewKeyValueCyperQueryBuilder(storage query.Storage) *KeyValueCyperQueryBuilder {
	return &KeyValueCyperQueryBuilder{
		storage: storage,
	}
}

func (s *KeyValueCyperQueryBuilder) Predicate(patterns []ast.Patn) ([]query.Predicate, error) {
	result := make([]query.Predicate, 0)
	for _, patn := range patterns {
		p, err := s.toPredicatePath(patn)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

// ToPredicatePath creates a PredicatePath out of the Patn
func (s *KeyValueCyperQueryBuilder) toPredicatePath(patn ast.Patn) (query.Predicate, error) {
	if vertex, ok := patn.(*ast.VertexPatn); ok {
		return s.ToPredicateVertexPath(vertex)
	}

	if edge, ok := patn.(*ast.EdgePatn); ok {
		return s.ToPredicateEdgePath(edge)
	}

	return nil, errUnknownPattern
}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (s *KeyValueCyperQueryBuilder) ToPredicateVertexPath(patn *ast.VertexPatn) (query.Predicate, error) {
	//label := strings.ToLower(patn.Label)
	if patn == nil {
		return nil, errNoPattern
	}
	return func(from, to *uuid.UUID, depth int) (string, query.Traverse) {
		keyValues := make([]*keyvaluestore.KeyValue, 0)

		for k, p := range patn.Properties {
			kv, _ := keyvaluestore.NewKeyValueProperty(from, k, p)

			iterator := s.storage.HasPrefix(kv.Key)
			for i, ok := iterator(); ok; i, ok = iterator() {
				if kv, ok := i.(*keyvaluestore.KeyValue); ok {
					if p != nil && p == keyvaluestore.Unmarshal(kv.Value) {
						keyValues = append(keyValues, kv)
					}
				}
			}

			if len(keyValues) > 0 {
				return patn.Variable, query.Matched
			}
		}

		return patn.Variable, query.Failed
	}, nil
}

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func (s *KeyValueCyperQueryBuilder) ToPredicateEdgePath(patn *ast.EdgePatn) (query.Predicate, error) {
	//label := strings.ToLower(patn.Body.Type)
	if patn == nil {
		return nil, errNoPattern
	}
	return func(from, to *uuid.UUID, depth int) (string, query.Traverse) {

		keyValues := make([]*keyvaluestore.KeyValue, 0)

		if patn.Body != nil {
			for k, p := range patn.Body.Properties {
				kv, _ := keyvaluestore.NewKeyValueRelationshipProperty(from, to, k, p)
				iterator := s.storage.HasPrefix(kv.Key)
				for i, ok := iterator(); ok; i, ok = iterator() {
					if kv, ok := i.(*keyvaluestore.KeyValue); ok {
						if p != nil && p == keyvaluestore.Unmarshal(kv.Value) {
							keyValues = append(keyValues, kv)
						}
					}
				}

				if len(keyValues) > 0 {
					return patn.Variable, query.Visiting
				}
			}
		}

		return patn.Variable, query.Matching

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

	}, nil

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
