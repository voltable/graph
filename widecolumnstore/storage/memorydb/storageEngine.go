package memorydb

import (
	"bytes"
	"errors"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/golang/protobuf/ptypes/any"
)

func init() {
	widecolumnstore.RegisterStorage(StorageType, widecolumnstore.StoreRegistration{
		NewFunc: NewStorageEngine,
	})
}

const StorageType = "memory"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type StorageEngine struct {
	tKeyIndex  map[int]string
	tKey       map[string]*any.Any
	tKeyTIndex map[int]string
	tKeyT      map[string]*any.Any
}

var _ widecolumnstore.Storage = (*StorageEngine)(nil)

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates a new in memory storage engine
func NewStorageEngine() (widecolumnstore.Storage, error) {
	se := StorageEngine{
		tKeyIndex:  make(map[int]string),
		tKey:       make(map[string]*any.Any),
		tKeyT:      make(map[string]*any.Any),
		tKeyTIndex: make(map[int]string),
	}

	return &se, nil
}

// Find a vertex from the persistence
func (se *StorageEngine) Find(ID string) (*widecolumnstore.KeyValue, error) {
	// if v, ok := se.vertices[ID]; ok {
	// 	return &v, nil
	// } else {
	// 	return nil, errRecordNotFound
	// }
	return nil, errRecordNotFound
}

func (se *StorageEngine) Each() widecolumnstore.Iterator {
	position := 0
	length := len(se.tKey)
	return func() (interface{}, bool) {
		for position < length {
			key := []byte(se.tKeyIndex[position])
			position = position + 1
			v := se.tKey[string(key)]
			kv := &widecolumnstore.KeyValue{Key: key, Value: v}
			return kv, true
		}

		return nil, false
	}
}

func (se *StorageEngine) HasPrefix(prefix []byte) widecolumnstore.Iterator {
	position := 0
	length := len(se.tKey)
	return func() (interface{}, bool) {
		for position < length {
			key := []byte(se.tKeyIndex[position])
			position = position + 1

			if bytes.HasPrefix(key, prefix) {
				v := se.tKey[string(key)]
				kv := &widecolumnstore.KeyValue{Key: key, Value: v}
				return kv, true
			}
		}

		return nil, false
	}
}

// Count number of keys/value pairs
func (se *StorageEngine) Count() int {
	return len(se.tKey)
}
