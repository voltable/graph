package cypher

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/Sirupsen/logrus"
)

var queryBuilderRegistry = make(map[string]QueryBuilderRegistry)

var (
	errQueryBuilderNotRegistred = errors.New("This QueryBuilder is not registered")
)

// QueryBuilder is the interface that a QueryBuilder must implement
type QueryBuilder interface {
	Predicate([]ast.Patn) []query.Predicate
}

type NewQueryBuilderFunc func(query.Storage) QueryBuilder

type QueryBuilderRegistry struct {
	NewFunc NewQueryBuilderFunc
}

func RegisterQueryBuilder(name string, register QueryBuilderRegistry) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := queryBuilderRegistry[name]; found {
		logrus.Panicf("Already registered QueryBuilder %q.", name)
	}
	queryBuilderRegistry[name] = register
}

func NewQueryBuilder(name string, storage query.Storage) (QueryBuilder, error) {
	r, registered := queryBuilderRegistry[name]
	if !registered {
		return nil, errQueryBuilderNotRegistred
	}

	return r.NewFunc(storage), nil
}

func FristQueryBuilder(storage query.Storage) (QueryBuilder, error) {
	for name := range queryBuilderRegistry {
		r, registered := queryBuilderRegistry[name]
		if !registered {
			return nil, errQueryBuilderNotRegistred
		}

		return r.NewFunc(storage), nil
	}
	return nil, errQueryBuilderNotRegistred
}
