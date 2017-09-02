package query

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var engineRegistry = make(map[string]EngineRegistration)

var (
	errQueryNotRegistred = errors.New("This query is not registered")
)

// Engine is the interface that a engine must implement
type Engine interface {
	// Parser in a string which is your query you want to run, get back a vertexPath that is abstracted from any query language or AST
	Parser(string) (Path, error)

	// Filter is used to run any final part of the AST on the result set
	Filter(*Query) error
}

type NewQueryFunc func() (Engine, error)

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

func NewQueryEngine(name string) (Engine, error) {
	r, registered := engineRegistry[name]
	if !registered {
		return nil, errQueryNotRegistred
	}

	return r.NewFunc()
}
