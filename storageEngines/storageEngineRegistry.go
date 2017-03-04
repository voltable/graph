package storageEngines

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

var storeEngineRegistry = make(map[string]StorageEngineRegistration)

var (
	ErrStorageEngineNotRegistred = errors.New("This StorageEngine is not registered.")
)

type NewStorageEngineFunc func(options *Options) (StorageEngine, error)

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

func NewStorageEngine(name string, options *Options) (StorageEngine, error) {
	r, registered := storeEngineRegistry[name]
	if !registered {
		return nil, ErrStorageEngineNotRegistred
	}

	return r.NewFunc(options)
}
