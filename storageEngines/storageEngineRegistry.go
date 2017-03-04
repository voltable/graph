package storageEngines

import (
	"fmt"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/Sirupsen/logrus"
)

var storeEngineRegistry = make(map[string]StorageEngineRegistration)

var (
	ErrQuadStoreNotRegistred = fmt.Errorf("This StorageEngine is not registered.")
)

type NewStorageEngineFunc func(options *graphs.Options) (graphs.StorageEngine, error)

type StorageEngineRegistration struct {
	NewFunc NewStorageEngineFunc
}

func RegisterStorageEngine(name string, register StorageEngineRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := storeEngineRegistry[name]; found {
		logrus.Panicf("Already registered StorageEngine %q.", name)
	}
	storeEngineRegistry[name] = register
}

func NewStorageEngine(name string, options *graphs.Options) (graphs.StorageEngine, error) {
	r, registered := storeEngineRegistry[name]
	if !registered {
		return nil, ErrQuadStoreNotRegistred
	}

	return r.NewFunc(options)
}
