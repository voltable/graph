package caudex

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
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
	db      *bolt.DB
	Options *Options
	opend   bool
	ready   bool

	change       map[string]Vertex
	delete       map[string]Vertex
	labelIndexes map[string]string
}

type GraphOperation struct {
	db *Graph
}

// Open graph
func Open(o *Options) *Graph {
	var st = &Graph{opend: true, Options: o}

	log.Println("Opening " + st.Options.Name)
	// It will be created if it doesn't exist.
	db, err := bolt.Open(st.Options.Name+".db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// create the bucket to hold the Adjacency list.
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketGraph))
		if err != nil {
			log.Fatal(err)
		}

		b, err = tx.CreateBucketIfNotExists([]byte(bucketIndex))
		if err != nil {
			log.Fatal(err)
		}

		// Loads up the types of label indexs, but not the actual indexes themself
		b.ForEach(func(k, v []byte) error {
			s := string(k)
			st.labelIndexes[s] = s
			return nil
		})
		return nil
	})

	st.ready = true
	st.db = db

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
	op := GraphOperation{db: g}
	err := fn(&op)

	g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))

		for _, vertex := range g.change {
			buf, err := json.Marshal(vertex)
			b.Put([]byte(vertex.ID), buf)
		}

		for _, vertex := range g.delete {
			b.Delete([]byte(vertex.ID))
		}

		return err
	})

	return err
}

// CreateVertex creates a vetex and returns the new vertex.
func (g *GraphOperation) CreateVertex() (*VertexOperation, *Vertex) {
	u1 := uuid.NewV4()
	vertex := Vertex{ID: u1.String(), Value: new(interface{})}
	g.db.change[u1.String()] = vertex
	op := VertexOperation{v: &vertex}
	return &op, &vertex
}

// RemoveVertex remvoes the vertex from the graph with any edges linking it
func (g *GraphOperation) RemoveVertex(v *Vertex) {
	if v == nil {
		return
	}

	if len(v.edges) > 0 {
		for _, edge := range v.edges {
			for i, otherEdge := range edge.to.edges {
				if otherEdge.edge == edge.edge {
					c := make([]Edge, len(edge.to.edges)-1)
					edge.to.edges = append(append(c, edge.to.edges[:i]...), edge.to.edges[i+1:]...)
					break
				}
			}
		}
	}
	g.db.delete[v.ID] = *v
}

func (g *Graph) updateLabelIndex(v *Vertex) {
	found := false

	for _, index := range g.labelIndexes {
		if index == v.label {
			found = true
			break
		}
	}

	if !found {
		c := make([]string, len(g.labelIndexes)+1)
		g.labelIndexes[v.label] = v.label
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
