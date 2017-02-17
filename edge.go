package graphs

// An Edge connects two Vertex in a graph.
type Edge struct {
	id         string
	from       *Vertex
	to         *Vertex
	isDirected Digraph
	weight     float32
	label      string
}

func (e *Edge) Weight() float32 {
	return e.weight
}

func (e *Edge) SetWeight(weight float32) {
	e.weight = weight
}

func (e *Edge) Label() string {
	return e.label
}

func (e *Edge) removeTo() *Vertex {
	if i, ok := e.to.edges["route"]; ok {
		delete(e.to.edges, e.id)
		return i.to
	}

	return nil
}

func (e *Edge) removeFrom() {
	delete(e.from.edges, e.id)
}
