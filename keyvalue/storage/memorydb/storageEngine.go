package memorydb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/uuid"

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
	tKey       map[string]*keyvalue.Any
	tKeyTIndex map[int]string
	tKeyT      map[string]*keyvalue.Any
	Options    *graph.Options
	engine     query.Engine
}

var _ graph.Graph = (*StorageEngine)(nil)

var _ query.Storage = (*StorageEngine)(nil)

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates anew in memory storage engine
func NewStorageEngine(o *graph.Options) (graph.Graph, error) {
	se := StorageEngine{
		Options:    o,
		tKeyIndex:  make(map[int]string),
		tKey:       make(map[string]*keyvalue.Any),
		tKeyT:      make(map[string]*keyvalue.Any),
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
		triples := keyvalue.MarshalKeyValue(v)
		transposes := keyvalue.MarshalKeyValueTranspose(v)
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
		triples := keyvalue.MarshalKeyValue(v)
		transposes := keyvalue.MarshalKeyValueTranspose(v)
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
func (se *StorageEngine) Find(ID string) (*keyvalue.KeyValue, error) {
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

func (se *StorageEngine) ForEach() query.IteratorUUID {
	position := 0
	length := len(se.tKey)
	return func() (*uuid.UUID, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			id, _ := uuid.ParseUUID(key)
			position = position + 1
			return id, true
		}
		return nil, false
	}
}

func (se *StorageEngine) HasPrefix(prefix []byte) query.Iterator {
	position := 0
	length := len(se.tKey)
	p := string(prefix)
	return func() (interface{}, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			position = position + 1

			if strings.HasPrefix(key, p) {
				v := se.tKey[key]
				kv := &keyvalue.KeyValue{Key: []byte(key), Value: v}
				return kv, true
			}
		}

		return nil, false
	}
}

func (se *StorageEngine) HasPrefixRange(prefixes [][]byte) query.Iterator {
	position := 0
	length := len(se.tKey)
	pp := make([]string, len(prefixes))
	for i := 0; i < len(prefixes); i++ {
		pp[i] = string(prefixes[i])
	}
	return func() (interface{}, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			position = position + 1

			for _, p := range pp {
				if strings.HasPrefix(key, p) {
					v := se.tKey[key]
					kv := &keyvalue.KeyValue{Key: []byte(key), Value: v}
					return kv, true
				}
			}

		}

		return nil, false
	}
}

func (se *StorageEngine) Edges(id *uuid.UUID) query.IteratorUUID {
	position := 0
	length := len(se.tKey)
	p := string(keyvalue.RelationshipPrefix(id))
	return func() (*uuid.UUID, bool) {
		for position < length {
			key := se.tKeyIndex[position]
			position = position + 1

			if strings.HasPrefix(key, p) {
				id, _ := uuid.ParseUUID(key)
				return id, true

			}

		}

		return nil, false
	}
}
