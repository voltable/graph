package memorydb

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/golang/protobuf/ptypes/any"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: NewStorageEngine,
	})
}

const GraphType = "memory"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type StorageEngine struct {
	tKeyIndex  map[int]string
	tKey       map[string]*any.Any
	tKeyTIndex map[int]string
	tKeyT      map[string]*any.Any
	Options    *graph.Options
	engine     query.Engine
}

var _ graph.Graph = (*StorageEngine)(nil)

var _ widecolumnstore.Storage = (*StorageEngine)(nil)

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates anew in memory storage engine
func NewStorageEngine(o *graph.Options) (graph.Graph, error) {
	se := StorageEngine{
		Options:    o,
		tKeyIndex:  make(map[int]string),
		tKey:       make(map[string]*any.Any),
		tKeyT:      make(map[string]*any.Any),
		tKeyTIndex: make(map[int]string),
	}

	queryEngine, err := query.NewQueryEngine(o.QueryEngine, &se)
	if err != nil {
		return nil, err
	}
	se.engine = queryEngine

	return &se, nil
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*graph.Vertex) error {
	for _, v := range c {
		triples, transposes := widecolumnstore.MarshalKeyValue(v)
		var errstrings []string

		for i := 0; i < len(triples); i++ {
			triple := triples[i]
			se.tKeyIndex[len(se.tKey)] = string(triple.Key)
			se.tKey[string(triple.Key)] = triple.Value

			transpose := transposes[i]
			se.tKeyTIndex[len(se.tKeyT)] = string(transpose.Key)
			se.tKeyT[string(transpose.Key)] = transpose.Value
		}

		if len(errstrings) > 0 {
			return fmt.Errorf(strings.Join(errstrings, "\n"))
		}
	}

	return nil
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*graph.Vertex) error {
	for _, v := range c {
		triples, transposes := widecolumnstore.MarshalKeyValue(v)
		var errstrings []string

		for i := 0; i < len(triples); i++ {
			triple := triples[i]
			delete(se.tKey, string(triple.Key))

			transpose := transposes[i]
			delete(se.tKeyT, string(transpose.Key))
		}

		if len(errstrings) > 0 {
			return fmt.Errorf(strings.Join(errstrings, "\n"))
		}
	}

	return nil
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

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*graph.Vertex) error {
	se.Create(c...)
	return nil
}

func (se *StorageEngine) Query(str string) (*graph.Query, error) {
	return se.engine.Parse(str)
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

func (se *StorageEngine) ForEach() widecolumnstore.IteratorUUID {
	position := 0
	length := len(se.tKey)
	return func() (*uuid.UUID, bool) {
		for position < length {
			key := []byte(se.tKeyIndex[position])
			kv := &widecolumnstore.KeyValue{
				Key: key,
			}
			position = position + 1
			return kv.UUID(), true
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

func (se *StorageEngine) Edges(id *uuid.UUID) widecolumnstore.IteratorUUIDWeight {
	position := 0
	length := len(se.tKey)
	p := widecolumnstore.RelationshipPrefix(id)
	return func() (*uuid.UUID, float64, bool) {
		for position < length {
			key := []byte(se.tKeyIndex[position])
			position = position + 1

			if bytes.HasPrefix(key, p) {
				kv := &widecolumnstore.KeyValue{
					Key: key,
				}
				return kv.To(), 0, true
			}
		}
		return nil, 0, false
	}
}
