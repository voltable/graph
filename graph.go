package graphs

import "errors"

var (
	errVertexNotFound = errors.New("Vertex Not found")
	errCreatVertex    = errors.New("Failed to create Vertex")
)

type (
	// Graph structure consisting of vertices and edges
	Graph struct {
		db StorageEngine
	}

	// Options for the graph
	Options struct {
		Name string
	}
)

func NewGraph(se StorageEngine) *Graph {
	g := &Graph{db: se}
	return g
}

func (g *Graph) Close() {
	g.db.Close()
}

// UpdateVertex retrieves a give vertex then lets you update it
func (g *Graph) UpdateVertex(ID string) (*Vertex, error) {

	var v *Vertex
	var err error
	if v, err = g.db.Find(ID); err == nil {
		return v, nil
	}

	return nil, errVertexNotFound
}

// DeleteVertex removes the vertex from the graph with any edges linking it
func (g *Graph) DeleteVertex(ID string) error {
	if v, err := g.db.Find(ID); err == nil {
		v.removeRelationships()
		return g.db.Delete(v)
	}

	return errVertexNotFound
}

func (g *Graph) CreateVertex() (*Vertex, error) {

	v, err := newVertex()
	if err != nil {
		return nil, err
	}

	if err := g.db.Create(v); err == nil {
		return v, nil
	}

	return nil, errCreatVertex
}

func (g *Graph) Query(root *Vertex) *Query {
	//todo need to setup channel from DFS or BFS
	c := make(chan *Vertex)
	return &Query{
		Iterate: func() Iterator {
			return func() (item *Vertex, ok bool) {
				v, ok := <-c
				return v, ok
			}
		},
	}

	return nil
}
