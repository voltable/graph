package boltdb

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: newStorageEngine,
	})
}

var (
	ErrVertexNotFound = errors.New("Vertex Not found")
	ErrCreatVertex    = errors.New("Failed to create Vertex")
)

const (
	GraphType          = "bolt"
	bucketGraph bucket = "graph"
	bucketLabel bucket = "label"
	bucketIndex bucket = "index"
)

type (
	//StorageEngine the underlying graph storage engine in this case boltdb
	StorageEngine struct {
		db      *bolt.DB
		Options *graph.Options
	}

	bucket string
)

func createBolt(o *graph.Options) *bolt.DB {
	var err error
	var db *bolt.DB
	var b *bolt.Bucket

	logrus.Info("Opening " + o.Name)
	// It will be created if it doesn't exist.
	if db, err = bolt.Open(o.Name+".db", 0600, &bolt.Options{Timeout: 1 * time.Second}); err != nil {
		logrus.Panic(err)
	}

	// create the bucket to hold the Adjacency list.
	db.Update(func(tx *bolt.Tx) error {
		if b, err = tx.CreateBucketIfNotExists([]byte(bucketGraph)); err != nil {
			logrus.Panic(err)
		}

		if b, err = tx.CreateBucketIfNotExists([]byte(bucketIndex)); err != nil {
			logrus.Panic(err)
		}

		return nil
	})

	return db
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*vertices.Vertex) error {
	var err error
	var buf []byte
	return se.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
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

// Delete the array of vertices from the persistence
func (g *StorageEngine) Delete(c ...*vertices.Vertex) error {
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
		for _, vertex := range c {
			b.Delete([]byte(vertex.ID()))
		}
		return nil
	})
}

// Find a vertex from the persistence
func (se *StorageEngine) Find(ID string) (*vertices.Vertex, error) {
	var err error
	var buf []byte
	var v vertices.Vertex
	return &v, se.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
		buf = b.Get([]byte(ID))

		if err = json.Unmarshal(buf, v); err == nil {
			return nil
		}

		return err
	})
}

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*vertices.Vertex) error {
	var err error
	var buf []byte
	return se.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
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
func (g *StorageEngine) Close() {
	g.db.Close()
}

func (se *StorageEngine) Query() *query.Query {
	//todo need to setup channel from DFS or BFS
	c := make(chan *vertices.Vertex)
	return &query.Query{
		Iterate: func() query.Iterator {
			return func() (item *vertices.Vertex, ok bool) {
				v, ok := <-c
				return v, ok
			}
		},
	}

	return nil
}

func (g *StorageEngine) backgroundTask(c chan os.Signal) {

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
