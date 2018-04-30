package graph

import uuid "github.com/hashicorp/go-uuid"
import "strings"

type (
	// An Edge connects two Vertex in a graph.
	Edge struct {
		id         string
		isDirected Digraph
		// Weight of a path in a weighted graph
		Weight           float64
		relationshipType string
		properties       map[string]interface{}
	}

	// Edges a array of edges
	Edges []*Edge
)

var _ Properties = (*Edge)(nil)

// NewEdge build a new edge with a id
func NewEdge() (*Edge, error) {
	var id string
	var err error

	if id, err = uuid.GenerateUUID(); err != nil {
		return nil, errCreatVertexID
	}

	v := Edge{id: id, properties: make(map[string]interface{})}
	return &v, nil
}

// ID returns the generate UUID
func (e *Edge) ID() string {
	return e.id
}

// RelationshipType the type of relationship
func (e *Edge) RelationshipType() string {
	return e.relationshipType
}

// SetRelationshipType the type of relationship
func (e *Edge) SetRelationshipType(label string) {
	e.relationshipType = strings.ToLower(label)
}

// SetProperty set a property to store against this Edge
func (e *Edge) SetProperty(name string, property interface{}) {
	e.properties[name] = property
}

// Property gets a property to store on the Edge
func (e *Edge) Property(name string) interface{} {
	return e.properties[name]
}

func (e *Edge) DeleteProperty(name string) {
	delete(e.properties, name)
}

func (e *Edge) PropertiesCount() int {
	return len(e.properties)
}

func (e *Edge) Properties() map[string]interface{} {
	return e.properties
}

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].Weight > a[j].Weight }
