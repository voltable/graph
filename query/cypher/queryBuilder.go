package cypher

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

var (
	errNoPattern      = errors.New("No pattern provided")
	errUnknownPattern = errors.New("Unknown pattern")
)

type QueryBuilder struct {
	storage widecolumnstore.Storage
}

func NewQueryBuilder(storage widecolumnstore.Storage) *QueryBuilder {
	return &QueryBuilder{
		storage: storage,
	}
}

func (s *QueryBuilder) Predicate(patterns []ast.Patn) (widecolumnstore.Operator, error) {
	var last widecolumnstore.Operator = widecolumnstore.NewScan(s.storage)
	var err error
	for _, patn := range patterns {
		last, err = s.toPredicatePath(patn, last)
		if err != nil {
			return nil, err
		}
	}

	return last, nil
}

// ToPredicatePath creates a PredicatePath out of the Patn̦P
func (s *QueryBuilder) toPredicatePath(patn ast.Patn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
	if vertex, ok := patn.(*ast.VertexPatn); ok {
		return s.ToPredicateVertexPath(vertex, last)
	}

	if edge, ok := patn.(*ast.EdgePatn); ok {
		return s.ToPredicateEdgePath(edge, last)
	}

	return nil, errUnknownPattern
}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (s *QueryBuilder) ToPredicateVertexPath(patn *ast.VertexPatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
	//label := strings.ToLower(patn.Label)
	if patn == nil {
		return nil, errNoPattern
	}
	for k := range patn.Properties {
		p := func(kv widecolumnstore.KeyValue) []byte {
			key := widecolumnstore.Key{}
			key.Unmarshal(kv.Key)
			return widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte(k), nil, key.ID}).Marshal()

		}
		last = widecolumnstore.NewFilter(s.storage, last, p)
	}
	return last, nil

	// return func(from, to *uuid.UUID, depth int) (string, query.Traverse) {
	// 	keyValues := make([]*widecolumnstore.KeyValue, 0)

	// 	for k, p := range patn.Properties {
	// 		kv, _ := query.NewKeyValueProperty(from, k, p)

	// 		iterator := s.storage.HasPrefix(kv.Key)
	// 		for i, ok := iterator(); ok; i, ok = iterator() {
	// 			if kv, ok := i.(*widecolumnstore.KeyValue); ok {
	// 				if p != nil && p == widecolumnstore.Unmarshal(kv.Value) {
	// 					keyValues = append(keyValues, kv)
	// 				}
	// 			}
	// 		}

	// 		if len(keyValues) > 0 {
	// 			return patn.Variable, query.Matched
	// 		}
	// 	}

	// 	return patn.Variable, query.Failed
	// }, nil
}

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func (s *QueryBuilder) ToPredicateEdgePath(patn *ast.EdgePatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
	//label := strings.ToLower(patn.Body.Type)
	if patn == nil {
		return nil, errNoPattern
	}

	return last, nil

	// return func(from, to *uuid.UUID, depth int) (string, query.Traverse) {

	// 	keyValues := make([]*widecolumnstore.KeyValue, 0)

	// 	if patn.Body != nil {
	// 		for k, p := range patn.Body.Properties {
	// 			kv, _ := query.NewKeyValueRelationshipProperty(from, to, k, p)
	// 			iterator := s.storage.HasPrefix(kv.Key)
	// 			for i, ok := iterator(); ok; i, ok = iterator() {
	// 				if kv, ok := i.(*widecolumnstore.KeyValue); ok {
	// 					if p != nil && p == widecolumnstore.Unmarshal(kv.Value) {
	// 						keyValues = append(keyValues, kv)
	// 					}
	// 				}
	// 			}

	// 			if len(keyValues) > 0 {
	// 				return patn.Variable, query.Visiting
	// 			}
	// 		}
	// 	}

	// 	return patn.Variable, query.Matching

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

	//}, nil

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
