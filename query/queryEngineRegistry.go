package query

import (
	"errors"

	graph "github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
	"github.com/sirupsen/logrus"
)

var queryRegistry = make(map[graph.QueryType]QueryEngineRegistry)

var (
	ErrGraphNotRegistred = errors.New("This QueryEngine is not registered.")
)

type NewGraphFunc func(widecolumnstore.Storage) (QueryEngine, error)

type QueryEngineRegistry struct {
	NewFunc NewGraphFunc
}

func RegisterQueryEngine(name graph.QueryType, register QueryEngineRegistry) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := queryRegistry[name]; found {
		logrus.Panicf("Already registered QueryEngine %q.", name)
	}
	queryRegistry[name] = register
}

func NewQueryEngine(name graph.QueryType, i widecolumnstore.Storage) (QueryEngine, error) {
	r, registered := queryRegistry[name]
	if !registered {
		return nil, ErrGraphNotRegistred
	}

	return r.NewFunc(i)
}
