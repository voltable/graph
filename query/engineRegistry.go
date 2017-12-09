package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/Sirupsen/logrus"
)

var engineRegistry = make(map[string]EngineRegistration)

var (
	errQueryNotRegistred = errors.New("This query is not registered")
)

// Engine is the interface that a engine must implement
type Engine interface {
	// Parse in a string which is your query you want to run, get back a vertexPath that is abstracted from any query language or AST
	Parse(q string) (*graph.Query, error)
}

type NewQueryFunc func(i graph.Storage) (Engine, error)

type EngineRegistration struct {
	NewFunc NewQueryFunc
}

func RegisterEngine(name string, register EngineRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := engineRegistry[name]; found {
		logrus.Panicf("Already registered Query %q.", name)
	}
	engineRegistry[name] = register
}

func NewQueryEngine(name string, i graph.Storage) (Engine, error) {
	r, registered := engineRegistry[name]
	if !registered {
		return nil, errQueryNotRegistred
	}

	return r.NewFunc(i)
}
