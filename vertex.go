package graph

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/voltable/graph/uuid"
)

// Vertex .
type Vertex struct {
	id         uuid.UUID
	edges      map[uuid.UUID]*Edge
	label      string
	properties map[string]interface{}
}

var (
	errCreatVertexID = errors.New("Failed to create Vertex ID")
	errEdgeNotFound  = errors.New("Edge Not found")
	errIDNotSet      = errors.New("Use NewVertex to create a new Vertex")
)

// NewVertex creates the default vertex
func NewVertex() (*Vertex, error) {
	return NewVertexWithLabel("")
}

// NewVertexFromID creates a vertex using the id
func NewVertexFromID(ID uuid.UUID) (*Vertex, error) {
	v, err := NewVertexWithLabel("")
	v.id = ID
	return v, err
}

// NewVertexWithLabel create a vertex with the set label
func NewVertexWithLabel(label string) (*Vertex, error) {
	var id uuid.UUID
	var err error

	if id, err = uuid.GenerateRandomUUID(); err != nil {
		return nil, errors.Wrap(errCreatVertexID, "UUID failed")
	}

	v := NewEmptyVertex()
	v.id = id
	v.label = label
	return v, nil
}

// NewEmptyVertex create's a empty vertex with no ID
func NewEmptyVertex() *Vertex {
	v := Vertex{edges: make(map[uuid.UUID]*Edge), properties: make(map[string]interface{})}
	return &v
}

// SetProperty set a property to store against this Vertex
func (v *Vertex) SetProperty(name string, property interface{}) {
	v.properties[name] = property
}

// Property gets a property to store on the Vertex
func (v *Vertex) Property(name string) interface{} {
	if value, ok := v.properties[name]; ok {
		return value
	}

	return nil
}

func (v *Vertex) DeleteProperty(name string) {
	delete(v.properties, name)
}

func (v *Vertex) PropertiesCount() int {
	return len(v.properties)
}

func (v *Vertex) Properties() map[string]interface{} {
	return v.properties
}

// ID returns the generate UUID
func (v *Vertex) ID() uuid.UUID {
	return v.id
}

// SetID set's the vertex id
func (v *Vertex) SetID(id uuid.UUID) {
	v.id = id
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

// Edges a array of all edges against this vertex
func (v *Vertex) Edges() map[uuid.UUID]*Edge {
	return v.edges
}

func (v *Vertex) removeRelationshipOnLabel(label string) Digraph {
	return v.removeRelationshipsF(func(id uuid.UUID, e *Edge) bool {
		return e.relationshipType == label
	})
}

func (v *Vertex) removeRelationships() {
	v.removeRelationshipsF(func(id uuid.UUID, e *Edge) bool {
		return true
	})
}

func (v *Vertex) removeRelationshipsOnVertex(to *Vertex) Digraph {
	return v.removeRelationshipsF(func(id uuid.UUID, e *Edge) bool {
		return id == to.id
	})
}

func (v *Vertex) removeRelationshipsF(f func(id uuid.UUID, e *Edge) bool) Digraph {
	for id, edge := range v.edges {
		if f(id, edge) {
			delete(v.edges, edge.to)
			return edge.isDirected
		}
	}
	return Undirected
}

func (v *Vertex) SetLabel(label string) *Vertex {
	v.label = strings.ToLower(label)
	return v
}

// AddDirectedEdge links two vertex's and returns the edge
func (v *Vertex) AddDirectedEdge(to *Vertex) (*Edge, error) {
	return v.AddDirectedEdgeWeight(to, 0)
}

// AddDirectedEdgeWeight links two vertex's with a weight and returns the edge
func (v *Vertex) AddDirectedEdgeWeight(to *Vertex, weight float64) (*Edge, error) {
	edge := &Edge{from: v.id, to: to.id, isDirected: Directed, properties: make(map[string]interface{}), Weight: weight}
	v.edges[edge.to] = edge
	return edge, nil
}

// AddEdge links two vertex's and returns the edge
func (v *Vertex) AddEdge(to *Edge) {

	v.edges[to.to] = to
}

// RemoveEdgeByLabel remove a edge
func (v *Vertex) RemoveEdgeByLabel(to *Vertex, label string) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := v.removeRelationshipOnLabel(label)

	if isDirected == Undirected {
		to.removeRelationshipOnLabel(label)
	}

	return nil
}

// RemoveEdge remove a edge
func (v *Vertex) RemoveEdge(to *Vertex) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := v.removeRelationshipsOnVertex(to)

	if isDirected == Undirected {
		to.removeRelationshipsOnVertex(v)
	}

	return nil
}

// func (v Vertex) String() string {

// 	var buffer bytes.Buffer
// 	buffer.WriteString("{")
// 	for k, b := range v.properties {
// 		buffer.WriteString(fmt.Sprintf("%v => %#v", k, b))
// 		buffer.WriteString(", ")
// 	}
// 	w := bytes.NewBuffer(buffer.Bytes()[:buffer.Len()-2])
// 	w.WriteString("}")
// 	return fmt.Sprintf(w.String())
// }
