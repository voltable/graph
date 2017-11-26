package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type FrontierQueue struct {
	Parts []interface{}
	Cost  float32
}

// FrontierVertex containers a vertex and it's Variable used by a query
type FrontierVertex struct {
	*vertices.Vertex
	Variable string
}

// FrontierEdge containers a edge and it's Variable used by a query
type FrontierEdge struct {
	*vertices.Edge
	Variable string
}

// Frontier priority queue containing vertices to be explored and the cost for a Uniform Cost Search
type Frontier struct {
	Values   []*FrontierQueue
	Explored map[string]bool
}

// Sort interface
func (f Frontier) Len() int           { return len(f.Values) }
func (f Frontier) Swap(i, j int)      { f.Values[i], f.Values[j] = f.Values[j], f.Values[i] }
func (f Frontier) Less(i, j int) bool { return f.Values[i].Cost < f.Values[j].Cost }

// Pop removes the FrontierQueue array
func (f *Frontier) Pop() (queue *FrontierQueue) {
	queue = f.Values[0]
	f.Values = f.Values[1:]
	return
}

// OptimalPath returns what should be the optimal path
//
// Must have run a sort on the Frontier before calling the OptimalPath
func (f Frontier) OptimalPath() []interface{} {
	return f.Values[0].Parts
}

// Append adds the vertices onto the frontier
func (f *Frontier) append(vertices []interface{}, cost float32) {
	fp := &FrontierQueue{vertices, cost}
	f.Values = append(f.Values, fp)
}

func (f *Frontier) AppendQueue(queue *FrontierQueue) {
	f.append(queue.Parts, queue.Cost)
}

func (f *Frontier) AppendVertex(queue *FrontierQueue, v *vertices.Vertex, variable string) {
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	f.append(append(queue.Parts, fv), queue.Cost)
}

func (f *Frontier) AppendEdgeAndVertex(queue *FrontierQueue, e *vertices.Edge, v *vertices.Vertex, variable string, weight float32) {
	fe := &FrontierEdge{Edge: e, Variable: variable}
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	parts := append(queue.Parts, fe)
	parts = append(parts, fv)
	f.append(parts, queue.Cost+weight)
}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *vertices.Vertex, variable string) Frontier {
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	f := Frontier{Explored: make(map[string]bool)}
	f.append([]interface{}{fv}, 0)
	return f
}
