package widecolumnstore

import (
	"errors"

	graph "github.com/voltable/graph"
	"github.com/Sirupsen/logrus"
)

var storeRegistry = make(map[graph.StorageType]StoreRegistration)

var (
	ErrStoreNotRegistred = errors.New("This Store is not registered.")
)

type NewStoreFunc func() (Storage, error)

type StoreRegistration struct {
	NewFunc NewStoreFunc
}

func RegisterStorage(name graph.StorageType, register StoreRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := storeRegistry[name]; found {
		logrus.Panicf("Already registered Storage %q.", name)
	}
	storeRegistry[name] = register
}

func NewStorage(name graph.StorageType) (Storage, error) {
	r, registered := storeRegistry[name]
	if !registered {
		return nil, ErrStoreNotRegistred
	}

	return r.NewFunc()
}
