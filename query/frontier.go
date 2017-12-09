package query

import graph "github.com/RossMerr/Caudex.Graph"

type FrontierQueue struct {
	Parts []FrontierProperties
	Cost  float32
}

// FrontierProperties containers a vertex or edge it's Variable used by a query
type FrontierProperties struct {
	Object   graph.Properties
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
func (f Frontier) OptimalPath() []FrontierProperties {
	return f.Values[0].Parts
}

// Clear removes the non optimal paths
func (f *Frontier) Clear() {
	f.Values = f.Values[:1]
	return
}

// Append adds the vertices onto the frontier
func (f *Frontier) append(vertices []FrontierProperties, cost float32) {
	fp := &FrontierQueue{vertices, cost}
	f.Values = append(f.Values, fp)
}

func (f *Frontier) AppendQueue(queue *FrontierQueue) {
	f.append(queue.Parts, queue.Cost)
}

func (f *Frontier) AppendVertex(queue *FrontierQueue, v *graph.Vertex, variable string) {
	fv := FrontierProperties{Object: v, Variable: variable}
	f.append(append(queue.Parts, fv), queue.Cost)
}

func (f *Frontier) AppendEdgeAndVertex(queue *FrontierQueue, e *graph.Edge, v *graph.Vertex, variable string, weight float32) {
	fe := FrontierProperties{Object: e, Variable: variable}
	fv := FrontierProperties{Object: v, Variable: variable}
	parts := append(queue.Parts, fe)
	parts = append(parts, fv)
	f.append(parts, queue.Cost+weight)
}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *graph.Vertex, variable string) Frontier {
	fv := FrontierProperties{Object: v, Variable: variable}
	f := Frontier{Explored: make(map[string]bool)}
	f.append([]FrontierProperties{fv}, 0)
	return f
}
