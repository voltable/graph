package caudex

// Vertex .
type Vertex struct {
	id    string
	edges []Edge
	Value *interface{}
	Label string
}

// AddDirectedEdge links two vertex's and returns the edge
func AddDirectedEdge(from *Vertex, to *Vertex) Edge {
	e := edge{}
	edge := Edge{from: from, to: to, edge: &e}
	from.edges = append(from.edges, edge)
	e.Edges = append(e.Edges, &edge)
	return edge
}

// AddEdge links two vertex's and returns the edge
func AddEdge(from *Vertex, to *Vertex) Edge {
	e := edge{}
	edge := Edge{from: from, to: to, edge: &e}
	from.edges = append(from.edges, edge)
	e.Parents = append(e.Parents, &edge)

	edge2 := Edge{from: to, to: from, edge: &e}
	to.edges = append(to.edges, edge2)
	e.Parents = append(e.Parents, &edge2)
	return edge
}

// RemoveEdge remove a edge
func RemoveEdge(e *Edge) {
	length := len(e.edge.Parents)

	for index := 0; index < length; index++ {
		var edge = e.edge.Parents[index]

		//a = append(a[:i], a[i+1:]...)
		//edge.from.edges
	}
}
