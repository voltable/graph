package cypher

import (
	"errors"

	"github.com/voltable/graph/query/cypher/ast"
	"github.com/voltable/graph/widecolumnstore"
)

var (
	ErrNoPattern      = errors.New("No pattern provided")
	ErrNoLastOperator = errors.New("No last operator provided")
	ErrUnknownPattern = errors.New("Unknown pattern")
)

type QueryBuilder interface {
	Build(stmt ast.Clauses) (widecolumnstore.Iterator, error)
}

type CypherQueryBuilder struct {
	storage widecolumnstore.Storage
}

func NewQueryBuilder(storage widecolumnstore.Storage) *CypherQueryBuilder {
	return &CypherQueryBuilder{
		storage: storage,
	}
}

// TODO need to build up the operators to run the query
func (s *CypherQueryBuilder) Build(stmt ast.Clauses) (widecolumnstore.Iterator, error) {

	return nil, nil
}

// func (s *CypherQueryBuilder) Predicate(patterns []ast.Patn) (widecolumnstore.Operator, error) {
// 	if patterns == nil {
// 		return nil, ErrNoPattern
// 	}

// 	var last widecolumnstore.Operator //= operators.NewScan(s.storage)
// 	var err error
// 	for _, patn := range patterns {
// 		last, err = s.toPredicatePath(patn, last)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return last, nil
// }

// // ToPredicatePath creates a PredicatePath out of the Patn̦
// func (s *CypherQueryBuilder) toPredicatePath(patn ast.Patn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
// 	if vertex, ok := patn.(*ast.VertexPatn); ok {
// 		return s.ToPredicateVertexPath(vertex, last)
// 	}

// 	if edge, ok := patn.(*ast.EdgePatn); ok {
// 		return s.ToPredicateEdgePath(edge, last)
// 	}

// 	return nil, ErrUnknownPattern
// }

// // ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
// func (s *CypherQueryBuilder) ToPredicateVertexPath(patn *ast.VertexPatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
// 	//label := strings.ToLower(patn.Label)
// 	if patn == nil {
// 		return nil, ErrNoPattern
// 	}

// 	if last == nil {
// 		return nil, ErrNoLastOperator
// 	}

// 	// var err error

// 	// predicate := func(interface{}) bool {
// 	// 	return true
// 	// }

// 	// for k := range patn.Properties {

// 	// 	//widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte(k), nil, key.ID}).Marshal()
// 	// 	last, err = s.filter(s.storage, []byte{}, predicate)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// }
// 	return last, nil
// }

// // ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
// func (s *CypherQueryBuilder) ToPredicateEdgePath(patn *ast.EdgePatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
// 	//label := strings.ToLower(patn.Body.Type)
// 	if patn == nil {
// 		return nil, ErrNoPattern
// 	}

// 	if last == nil {
// 		return nil, ErrNoLastOperator
// 	}

// 	// predicate := func(interface{}) bool {
// 	// 	return true
// 	// }

// 	// var err error
// 	// if patn.Body != nil {
// 	// 	for k := range patn.Body.Properties {
// 	// 		operator := func(key widecolumnstore.Key) []byte {
// 	// 			return widecolumnstore.NewKey(key.ID, &widecolumnstore.Column{query.Relationshipproperties, []byte(k), nil}).Marshal()
// 	// 		}
// 	// 		last, err = s.filter(s.storage, operator, predicate)
// 	// 		if err != nil {
// 	// 			return nil, err
// 	// 		}
// 	// 	}
// 	// }

// 	return last, nil
// 	// relationshipType := strings.ToLower(patn.Body.Type)
// 	// pvp := query.PredicateEdgePath{PredicateEdge: func(v *graph.Edge, depth uint) (string, query.Traverse) {

// 	// 	if depth < patn.Body.LengthMinimum {
// 	// 		return patn.Body.Variable, query.Visiting
// 	// 	}

// 	// 	if depth > patn.Body.LengthMaximum {
// 	// 		return patn.Body.Variable, query.Failed
// 	// 	}

// 	// 	if relationshipType != emptyString {
// 	// 		if relationshipType != v.RelationshipType() {
// 	// 			return patn.Body.Variable, query.Failed
// 	// 		}
// 	// 	}

// 	// 	for key, value := range patn.Body.Properties {
// 	// 		if v.Property(key) != value {
// 	// 			return patn.Body.Variable, query.Failed
// 	// 		}
// 	// 	}

// 	// 	return patn.Body.Variable, query.Matching
// 	// }, Variable: patn.Variable}

// 	// return &pvp
// }
