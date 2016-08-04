package caudex

import "github.com/satori/go.uuid"

// Tx represents a transaction on the Graph.
type Tx struct {
	db      *Graph
	vertexs map[string]Vertex
	deletes map[string]Vertex
}

