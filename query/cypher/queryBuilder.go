package cypher

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore/operators"
)

var (
	ErrNoPattern      = errors.New("No pattern provided")
	ErrNoLastOperator = errors.New("No last operator provided")
	ErrUnknownPattern = errors.New("Unknown pattern")
)

type QueryBuilder struct {
	storage widecolumnstore.Storage
	filter  func(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) widecolumnstore.Unary
}

func NewQueryBuilderDefault(storage widecolumnstore.Storage) *QueryBuilder {
	return NewQueryBuilder(storage, operators.NewFilter)
}

func NewQueryBuilder(storage widecolumnstore.Storage,
	filter func(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) widecolumnstore.Unary) *QueryBuilder {
	return &QueryBuilder{
		storage: storage,
		filter:  filter,
	}
}

func (s *QueryBuilder) Predicate(patterns []ast.Patn) (widecolumnstore.Operator, error) {
	var last widecolumnstore.Operator = operators.NewScan(s.storage)
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

	return nil, ErrUnknownPattern
}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (s *QueryBuilder) ToPredicateVertexPath(patn *ast.VertexPatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
	//label := strings.ToLower(patn.Label)
	if patn == nil {
		return nil, ErrNoPattern
	}

	if last == nil {
		return nil, ErrNoLastOperator
	}

	for k := range patn.Properties {
		operator := func(key widecolumnstore.Key) []byte {
			return widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte(k), nil, key.ID}).Marshal()
		}
		last = s.filter(s.storage, last, operator)
	}
	return last, nil
}

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func (s *QueryBuilder) ToPredicateEdgePath(patn *ast.EdgePatn, last widecolumnstore.Operator) (widecolumnstore.Operator, error) {
	//label := strings.ToLower(patn.Body.Type)
	if patn == nil {
		return nil, ErrNoPattern
	}

	if last == nil {
		return nil, ErrNoLastOperator
	}

	if patn.Body != nil {
		for k := range patn.Body.Properties {
			operator := func(key widecolumnstore.Key) []byte {
				return widecolumnstore.NewKey(key.ID, &widecolumnstore.Column{query.Relationshipproperties, []byte(k), nil}).Marshal()
			}
			last = s.filter(s.storage, last, operator)
		}
	}

	return last, nil
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
