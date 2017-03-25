package vertices

type (
	// An Edge connects two Vertex in a graph.
	Edge struct {
		id               string
		isDirected       Digraph
		weight           float32
		relationshipType string
		properties       map[string]interface{}
	}

	Edges []*Edge
)

func (e *Edge) ID() string {
	return e.id
}

// Weight of a path in a weighted graph
func (e *Edge) Weight() float32 {
	return e.weight
}

// SetWeight sets the edge weight
func (e *Edge) SetWeight(weight float32) {
	e.weight = weight
}

// RelationshipType the type of relationship
func (e *Edge) RelationshipType() string {
	return e.relationshipType
}

// SetRelationshipType the type of relationship
func (e *Edge) SetRelationshipType(label string) {
	e.relationshipType = label
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

func (a Edges) Len() int           { return len(a) }
func (a Edges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool { return a[i].weight > a[j].weight }
