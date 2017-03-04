package traversal

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var traversalRegistry = make(map[string]TraversalRegistration)

var ErrTraversalNotRegistred = errors.New("This Traversal is not registered.")

type NewTraversalFunc func() (Traversal, error)

type TraversalRegistration struct {
	NewFunc NewTraversalFunc
}

func RegisterTraversal(name string, register TraversalRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := traversalRegistry[name]; found {
		logrus.Panicf("Already registered Traversal %q.", name)
	}
	traversalRegistry[name] = register
}

func NewTraversal(name string) (Traversal, error) {
	r, registered := traversalRegistry[name]
	if !registered {
		return nil, ErrTraversalNotRegistred
	}

	return r.NewFunc()
}
