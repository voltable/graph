package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/Sirupsen/logrus"
)

var queryRegistry = make(map[string]QueryEngineRegistry)

var (
	ErrGraphNotRegistred = errors.New("This QueryEngine is not registered.")
)

type NewGraphFunc func(widecolumnstore.Storage) (QueryEngine, error)

type QueryEngineRegistry struct {
	NewFunc NewGraphFunc
}

func RegisterQueryEngine(name string, register QueryEngineRegistry) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := queryRegistry[name]; found {
		logrus.Panicf("Already registered QueryEngine %q.", name)
	}
	queryRegistry[name] = register
}

func NewQueryEngine(name string, i widecolumnstore.Storage) (QueryEngine, error) {
	r, registered := queryRegistry[name]
	if !registered {
		return nil, ErrGraphNotRegistred
	}

	return r.NewFunc(i)
}
