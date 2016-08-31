package caudex

import (
	"encoding/json"
	"log"

	"bitbucket.org/rossmerr/caudex/graph/engines"
	"github.com/boltdb/bolt"
)

const (
	bucketGraph = "graph"
	bucketLabel = "label"
	bucketIndex = "index"
)

// Options for the graph
type Options struct {
	Name string
}

// Graph the underlying graph
type Graph struct {
	db      engines.StorageEngine
	Options *Options
	opend   bool
	ready   bool
}

// Open graph
func Open(o *Options) *Graph {
	var st = &Graph{opend: true, Options: o}

	log.Println("Opening " + st.Options.Name)
	p := engines.StorageEngineProvider{}
	e := p.CreateEngine()

	// It will be created if it doesn't exist.
	err := e.Open(st.Options.Name + ".db")
	if err != nil {
		log.Fatal(err)
	}

	e.CreateIndexes()

	st.ready = true
	st.db = e

	return st
}

// Close graph
func (g *Graph) Close() {
	defer g.db.Close()
}

// Query over the graph using the cypher query language returns JSON
func (g *Graph) Query(cypher string) string {
	parse(cypher)
	return "test"
}

// Update
func (g *Graph) Update(fn func(*GraphOperation) error) error {
	o := GraphOperation{db: g}
	err := fn(&o)
	labelIndexCountChange := make(map[string]uint64)

	g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))

		for _, vertex := range o.add {
			buf, err := json.Marshal(vertex)
			b.Put([]byte(vertex.ID), buf)
		}

		for _, vertex := range o.change {
			buf, err := json.Marshal(vertex)
			b.Put([]byte(vertex.ID), buf)
		}

		for _, vertex := range o.delete {
			b.Delete([]byte(vertex.ID))
		}

		return err
	})

	return err
}

func updateAddCount(m map[string]uint64, v *Vertex) {
	for key, value := range m {
		if v.label == key {
			m[key] = value + 1
			break
		}
	}
}

func (g *Graph) updateLabelIndex(v *Vertex) {
	found := false

	for key, index := range g.labelIndexes {
		if key == v.label {
			found = true
			break
		}
	}

	if !found {
		c := make([]string, len(g.labelIndexes)+1)
		g.labelIndexes[v.label] = uint64(1)
		g.db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(bucketLabel + "_" + v.label))
			if err != nil {
				log.Fatal(err)
			}

			b = tx.Bucket([]byte(bucketIndex))
			// we store the number of vertex in the index
			b.Put([]byte(v.label), []byte("1"))

			return nil
		})
	}
}
