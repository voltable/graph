package query

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var queryEngineRegistry = make(map[string]QueryEngineRegistration)

var (
	errQueryNotRegistred = errors.New("This query is not registered.")
)

// QueryEngine is the interface that a queryEngine must implement
type QueryEngine interface {
	// Parser in a string which is your query you want to run, get back a vertexPath that is abstracted from any query language or AST
	Parser(string) (*QueryPath, error)
}

type NewQueryFunc func() (QueryEngine, error)

type QueryEngineRegistration struct {
	NewFunc NewQueryFunc
}

func RegisterQueryEngine(name string, register QueryEngineRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := queryEngineRegistry[name]; found {
		logrus.Panicf("Already registered Query %q.", name)
	}
	queryEngineRegistry[name] = register
}

func NewQueryEngine(name string) (QueryEngine, error) {
	r, registered := queryEngineRegistry[name]
	if !registered {
		return nil, errQueryNotRegistred
	}

	return r.NewFunc()
}
