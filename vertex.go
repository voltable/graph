package graph

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	uuid "github.com/hashicorp/go-uuid"
)

// Vertex .
type Vertex struct {
	id         string
	edges      map[string]*Edge
	label      string
	properties map[string]interface{}
}

var (
	errCreatVertexID = errors.New("Failed to create Vertex ID")
	errEdgeNotFound  = errors.New("Edge Not found")
	errIdNotSet      = errors.New("Use NewVertex to create a new Vertex")
)

var _ Properties = (*Vertex)(nil)

// NewVertex creates the default vertex
func NewVertex() (*Vertex, error) {
	return NewVertexWithLabel("")
}

// NewVertexWithLabel create a vertex with the set label
func NewVertexWithLabel(label string) (*Vertex, error) {
	var id string
	var err error

	if id, err = uuid.GenerateUUID(); err != nil {
		return nil, errCreatVertexID
	}

	v := Vertex{id: id, edges: make(map[string]*Edge), properties: make(map[string]interface{}), label: label}
	return &v, nil
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
func (v *Vertex) ID() string {
	return v.id
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

// Edges a array of all edges against this vertex
func (v *Vertex) Edges() Edges {
	edges := make(Edges, 0, len(v.edges))
	for _, value := range v.edges {
		edges = append(edges, value)
	}

	return edges
}

func (v *Vertex) removeRelationshipOnLabel(label string) Digraph {
	return v.removeRelationshipsF(func(id string, e Edge) bool {
		return e.relationshipType == label
	})
}

func (v *Vertex) removeRelationships() {
	v.removeRelationshipsF(func(id string, e Edge) bool {
		return true
	})
}

func (v *Vertex) removeRelationshipsOnVertex(to *Vertex) Digraph {
	return v.removeRelationshipsF(func(id string, e Edge) bool {
		return id == to.id
	})
}

func (v *Vertex) removeRelationshipsF(f func(id string, e Edge) bool) Digraph {
	for id, edge := range v.edges {
		if f(id, *edge) {
			delete(v.edges, edge.id)
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
	if to.id == EmptyString {
		return nil, errIdNotSet
	}

	edge := &Edge{id: to.id, isDirected: Directed}
	v.edges[edge.id] = edge
	return edge, nil
}

// AddEdge links two vertex's and returns the edge
func (v *Vertex) AddEdge(to *Vertex) (*Edge, *Edge, error) {
	return v.AddEdgeWeight(to, float64(0))
}

// AddEdgeWeight links two vertex's with a weight and returns the edge
func (v *Vertex) AddEdgeWeight(to *Vertex, weight float64) (*Edge, *Edge, error) {
	if to.id == EmptyString {
		return nil, nil, errIdNotSet
	}

	edge := &Edge{id: to.id, isDirected: Undirected, Weight: weight}
	v.edges[edge.id] = edge

	if v.id == EmptyString {
		return nil, nil, errIdNotSet
	}

	edge2 := &Edge{id: v.id, isDirected: Undirected, Weight: weight}
	to.edges[edge2.id] = edge2
	return edge, edge2, nil
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

func (b Vertex) String() string {

	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, v := range b.properties {
		buffer.WriteString(fmt.Sprintf("%v => %#v", k, v))
		buffer.WriteString(", ")
	}
	w := bytes.NewBuffer(buffer.Bytes()[:buffer.Len()-2])
	w.WriteString("}")
	return fmt.Sprintf(w.String())
}
