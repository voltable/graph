package memorydb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/gogo/protobuf/proto"
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
	tKey    map[string][]byte
	tKeyT   map[string][]byte
	Options *graph.Options
	engine  query.Engine
}

var _ graph.Graph = (*StorageEngine)(nil)

var _ graph.Storage = (*StorageEngine)(nil)

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates anew in memory storage engine
func NewStorageEngine(o *graph.Options) (graph.Graph, error) {
	se := StorageEngine{
		Options: o,
		tKey:    make(map[string][]byte),
		tKeyT:   make(map[string][]byte),
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
		triples := v.MarshalKeyValue()
		transposes := v.MarshalKeyValueTranspose()
		var errstrings []string

		for i := 0; i < len(triples); i++ {
			triple := triples[i]
			buf, _ := proto.Marshal(triple.Value)
			se.tKey[string(triple.Key)] = buf

			transpose := transposes[i]
			buf, _ = proto.Marshal(transpose.Value)
			se.tKeyT[string(transpose.Key)] = buf
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
		triples := v.MarshalKeyValue()
		transposes := v.MarshalKeyValueTranspose()
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
func (se *StorageEngine) Find(ID string) (*graph.Vertex, error) {
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

func (se *StorageEngine) Fetch(id string) (*graph.Vertex, error) {
	return se.Find(id)
}

func (se *StorageEngine) ForEach() graph.Iterator {
	// position := 0
	// length := len(se.keys)
	// return func() (item interface{}, ok bool) {
	// 	if position < length {
	// 		key := se.keys[position]
	// 		v := se.vertices[key]
	// 		position = position + 1
	// 		return &v, true
	// 	}
	// 	return nil, false
	// }

	return func() (item interface{}, ok bool) {
		return nil, false
	}
}
