package widecolumnstore

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var storeRegistry = make(map[string]StoreRegistration)

var (
	ErrStoreNotRegistred = errors.New("This Store is not registered.")
)

type NewStoreFunc func() (Storage, error)

type StoreRegistration struct {
	NewFunc NewStoreFunc
}

func RegisterStorage(name string, register StoreRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := storeRegistry[name]; found {
		logrus.Panicf("Already registered Storage %q.", name)
	}
	storeRegistry[name] = register
}

func NewStorage(name string) (Storage, error) {
	r, registered := storeRegistry[name]
	if !registered {
		return nil, ErrStoreNotRegistred
	}

	return r.NewFunc()
}
