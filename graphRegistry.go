package graph

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var graphRegistry = make(map[string]GraphRegistration)

var (
	ErrGraphNotRegistred = errors.New("This Graph is not registered.")
)

type NewGraphFunc func(options *Options) (Graph, error)

type GraphRegistration struct {
	NewFunc NewGraphFunc
}

func RegisterGraph(name string, register GraphRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := graphRegistry[name]; found {
		logrus.Panicf("Already registered Graph %q.", name)
	}
	graphRegistry[name] = register
}

func NewGraph(name string, options *Options) (Graph, error) {
	r, registered := graphRegistry[name]
	if !registered {
		return nil, ErrGraphNotRegistred
	}

	return r.NewFunc(options)
}
