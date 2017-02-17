package boltdb

import (
	"encoding/json"
	"os"
	"time"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

// Graph the underlying graph
type Graph struct {
	db      *bolt.DB
	Options *graphs.Options
}

func CreateBolt(o *graphs.Options) *bolt.DB {
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
		if b, err = tx.CreateBucketIfNotExists([]byte(BucketGraph)); err != nil {
			logrus.Panic(err)
		}

		if b, err = tx.CreateBucketIfNotExists([]byte(BucketIndex)); err != nil {
			logrus.Panic(err)
		}

		return nil
	})

	return db
}

func (g *Graph) Create(c *[]graphs.Vertex) error {
	var err error
	var buf []byte
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketGraph))
		for _, vertex := range *c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID), buf)
			} else {
				return err
			}
		}
		return nil
	})
}

func (g *Graph) Delete(c *[]graphs.Vertex) error {
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketGraph))
		for _, vertex := range *c {
			b.Delete([]byte(vertex.ID))
		}
		return nil
	})
}

func (g *Graph) Find(ID string) (*graphs.Vertex, error) {
	var err error
	var buf []byte
	var v graphs.Vertex
	return &v, g.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketGraph))
		buf = b.Get([]byte(ID))

		if err = json.Unmarshal(buf, v); err == nil {
			return nil
		}

		return err
	})
}

func (g *Graph) Update(c *[]graphs.Vertex) error {
	var err error
	var buf []byte
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketGraph))
		for _, vertex := range *c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID), buf)
			} else {
				return err
			}
		}
		return nil
	})
}

// Open graph
func Open(o *graphs.Options) graphs.Graph {
	g := Graph{Options: o, db: CreateBolt(o)}
	c := make(chan os.Signal, 1)
	g.backgroundTask(c)
	return &g
}

// Close graph
func (g *Graph) Close() {
	g.db.Close()
}

// Query over the graph using the cypher query language returns JSON
func (g *Graph) Query(cypher string) string {
	query.Parse(cypher)
	return "test"
}

// Update
func (g *Graph) Command(fn func(*graphs.GraphOperation) error) error {
	op := &graphs.GraphOperation{DB: g}
	return fn(op)
}

func (g *Graph) backgroundTask(c chan os.Signal) {

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
