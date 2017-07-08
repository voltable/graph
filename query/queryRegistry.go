package query

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var queryRegistry = make(map[string]QueryRegistration)

var (
	errQueryNotRegistred = errors.New("This query is not registered.")
)

type NewQueryFunc func() (Query, error)

type QueryRegistration struct {
	NewFunc NewQueryFunc
}

func RegisterQuery(name string, register QueryRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := queryRegistry[name]; found {
		logrus.Panicf("Already registered Query %q.", name)
	}
	queryRegistry[name] = register
}

func NewQuery(name string) (Query, error) {
	r, registered := queryRegistry[name]
	if !registered {
		return nil, errQueryNotRegistred
	}

	return r.NewFunc()
}
