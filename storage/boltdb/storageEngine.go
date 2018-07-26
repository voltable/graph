package boltdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RossMerr/Caudex.Graph/keyvalue"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/Sirupsen/logrus"
	bolt "github.com/coreos/bbolt"
	"github.com/gogo/protobuf/proto"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: newStorageEngine,
	})
}

var (
	// ErrVertexNotFound Vertex Not found
	ErrVertexNotFound = errors.New("Vertex Not found")

	// ErrCreatVertex Failed to create Vertex
	ErrCreatVertex = errors.New("Failed to create Vertex")
)

const (
	// GraphType the underlying storage, bolt
	GraphType = "bolt"
)

var (
	// TKey primary table for graph
	TKey = []byte("TKey")

	// TKeyT is the transpose of TKey
	TKeyT = []byte("TKeyT")

	// Ttxt holds original values
	Ttxt = []byte("Ttxt")
)

type (
	//StorageEngine the underlying graph storage engine in this case boltdb
	StorageEngine struct {
		db      *bolt.DB
		Options *graph.Options
	}

	bucket string
)

var _ graph.Graph = (*StorageEngine)(nil)

func createBolt(o *graph.Options) *bolt.DB {
	var err error
	var db *bolt.DB

	logrus.Info("Opening " + o.Name)
	// It will be created if it doesn't exist.
	if db, err = bolt.Open(o.Name+".db", 0600, &bolt.Options{Timeout: 1 * time.Second}); err != nil {
		logrus.Panic(err)
	}

	// create the bucket to hold the Adjacency list.
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(TKey)
		if err != nil {
			logrus.Panic(err)
		}

		_, err = tx.CreateBucketIfNotExists(TKeyT)
		if err != nil {
			logrus.Panic(err)
		}

		_, err = tx.CreateBucketIfNotExists(Ttxt)
		if err != nil {
			logrus.Panic(err)
		}

		return nil
	})

	return db
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*graph.Vertex) error {
	return se.db.Update(func(tx *bolt.Tx) error {
		bucketTKey := tx.Bucket(TKey)
		bucketTKeyT := tx.Bucket(TKeyT)
		var errstrings []string

		for _, v := range c {
			triples := v.MarshalKeyValue()
			transposes := v.MarshalKeyValueTranspose()
			for i := 0; i < len(triples); i++ {
				triple := triples[i]

				buf, _ := proto.Marshal(triple.Value)
				if err := bucketTKey.Put(triple.Key, buf); err != nil {
					errstrings = append(errstrings, err.Error())
				}

				transpose := transposes[i]
				buf, _ = proto.Marshal(transpose.Value)
				if err := bucketTKeyT.Put(transpose.Key, buf); err != nil {
					errstrings = append(errstrings, err.Error())
				}
			}
		}
		if len(errstrings) > 0 {
			return fmt.Errorf(strings.Join(errstrings, "\n"))
		}
		return nil
	})
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*graph.Vertex) error {
	return se.db.Update(func(tx *bolt.Tx) error {
		bucketTKey := tx.Bucket(TKey)
		bucketTKeyT := tx.Bucket(TKeyT)
		var errstrings []string

		for _, v := range c {
			triples := v.MarshalKeyValue()
			transposes := v.MarshalKeyValueTranspose()
			for i := 0; i < len(triples); i++ {
				triple := triples[i]
				if err := bucketTKey.Delete(triple.Key); err != nil {
					errstrings = append(errstrings, err.Error())
				}

				transpose := transposes[i]
				if err := bucketTKeyT.Delete(transpose.Key); err != nil {
					errstrings = append(errstrings, err.Error())
				}
			}
		}
		if len(errstrings) > 0 {
			return fmt.Errorf(strings.Join(errstrings, "\n"))
		}
		return nil
	})
}

// Find a vertex from the persistence
func (se *StorageEngine) Find(ID string) (*graph.Vertex, error) {
	var v *graph.Vertex

	return v, se.db.View(func(tx *bolt.Tx) error {
		bucketTKey := tx.Bucket(TKey)
		c := bucketTKey.Cursor()

		prefix := []byte(ID)
		kv := []*keyvalue.KeyValue{}
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			var any *keyvalue.Any
			err := proto.Unmarshal(v, any)
			if err != nil {
				return err
			}
			kv = append(kv, &keyvalue.KeyValue{Key: k, Value: any})
		}

		v.UnmarshalKeyValue(kv)
		return nil
	})
}

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*graph.Vertex) error {
	var err error
	var buf []byte
	return se.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(TKey)
		for _, vertex := range c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID()), buf)
			} else {
				return err
			}
		}
		return nil
	})
}

// NewStorageEngine creates a bolt graph
func newStorageEngine(o *graph.Options) (graph.Graph, error) {
	se := &StorageEngine{Options: o, db: createBolt(o)}
	c := make(chan os.Signal, 1)
	se.backgroundTask(c)
	return se, nil
}

// Close graph
func (se *StorageEngine) Close() {
	se.db.Close()
}

// Query used to query the graph
func (se *StorageEngine) Query(q string) (*graph.Query, error) {
	// iterate := func() query.Iterator {
	// 	ch := make(chan vertices.Vertex)
	// 	go se.db.View(func(tx *bolt.Tx) error {
	// 		b := tx.Bucket([]byte(bucketGraph))
	// 		b.ForEach(func(k, v []byte) error {
	// 			vertex := vertices.Vertex{}
	// 			if err := json.Unmarshal(v, vertex); err == nil {
	//				//TODO need to return a frontier query.NewFrontier(&v), true
	// 				ch <- vertex
	// 			}
	// 			return nil
	// 		})
	// 		close(ch)
	// 		return nil
	// 	})

	// 	return func() (item interface{}, ok bool) {
	// 		v, ok := <-ch
	// 		frontier := query.Frontier{&query.Path{[]*vertices.Vertex{&v}, 0}}
	// 		return frontier, ok
	// 	}
	// }

	//return query.NewVertexPath(iterate, se.Find), nil
	return nil, nil
}

func (se *StorageEngine) backgroundTask(c chan os.Signal) {

	go func() {
	Loop:
		for {
			select {
			case <-c:
				logrus.Debug("Received an interrupt, stopping services...")
				break Loop
			}
		}
		close(c)
	}()
}
