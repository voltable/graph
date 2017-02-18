package memorydb

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/rossmerr/Caudex.Graph"
)

type Graph struct {
	vertices []graphs.Vertex
}

func (g *Graph) Close() {

}
func (g *Graph) Query(cypher string) string {
	return ""
}

func (g *Graph) Command(fn func(*graphs.GraphOperation) error) error {
	op := &graphs.GraphOperation{DB: g}
	return fn(op)
}

// Create adds a array of vertices to the persistence
func (g *Graph) Create(c []graphs.Vertex) error {
	var err error
	var buf []byte

	g.vertices = append(g.vertices, c...)

	return nil
}

// Delete the array of vertices from the persistence
func (g *Graph) Delete(c []graphs.Vertex) error {
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
		for _, vertex := range *c {
			b.Delete([]byte(vertex.ID))
		}
		return nil
	})
}

// Find a vertex from the persistence
func (g *Graph) Find(ID string) (*graphs.Vertex, error) {
	var err error
	var buf []byte
	var v graphs.Vertex
	return &v, g.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
		buf = b.Get([]byte(ID))

		if err = json.Unmarshal(buf, v); err == nil {
			return nil
		}

		return err
	})
}

// Update the array of vertices from the persistence
func (g *Graph) Update(c []graphs.Vertex) error {
	var err error
	var buf []byte
	return g.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketGraph))
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
