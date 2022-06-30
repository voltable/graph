package memorydb

import (
	"bytes"
	"errors"

	"github.com/voltable/graph/storage"
	"github.com/voltable/graph/storage/table"
)

const StorageType = "memory"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type StorageEngine struct {
	table *table.Table[float32]
}

var _ storage.Storage = (*StorageEngine)(nil)

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates a new in memory storage engine
func NewStorageEngine() (storage.Storage, error) {
	se := StorageEngine{
		table: table.NewTable[float32](1000, 1000),
	}

	return &se, nil
}

func (se *StorageEngine) Each() storage.Iterator {
	position := 0
	length := len(se.tKey)
	return func() (storage.SparseValue, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			position = position + 1
			kv := se.tKey[key]
			return kv, true
		}

		return nil, false
	}
}

func (se *StorageEngine) HasPrefix(subject, predicate string) storage.Iterator {
	vector := se.table.ColumnsAt(subject)
	v := vector.At(predicate)
	position := 0
	length := len(se.tKey)
	return func() (storage.SparseValue, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			position = position + 1
			if bytes.HasPrefix([]byte(key), prefix) {
				kv := se.tKey[key]
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

func (se *StorageEngine) Query(string) storage.Iterator {
	return nil
}

func (se *StorageEngine) Create(triples ...storage.Triple[float32]) error {

	se.table.Create(triples...)
	return nil
}
