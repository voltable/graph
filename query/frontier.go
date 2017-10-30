package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type frontierPath struct {
	Vertices []*FrontierVertex
	Cost     float32
	Traverse Traverse
}

// FrontierVertex containers a vertex and it's Variable used by a query
type FrontierVertex struct {
	*vertices.Vertex
	Variable string
}

// Frontier priority queue containing vertices to be explored and the cost for a Uniform Cost Search
type Frontier struct {
	Values []*frontierPath
}

// Sort interface
func (f Frontier) Len() int           { return len(f.Values) }
func (f Frontier) Swap(i, j int)      { f.Values[i], f.Values[j] = f.Values[j], f.Values[i] }
func (f Frontier) Less(i, j int) bool { return f.Values[i].Cost < f.Values[j].Cost }

// Pop removes the FrontierVertex array and cost from the Frontier
func (f *Frontier) Pop() (vertices []*FrontierVertex, cost float32) {
	vertices = f.Values[0].Vertices
	cost = f.Values[0].Cost
	f.Values = f.Values[1:]
	return
}

// OptimalPath returns what should be the optimal path
//
// Must have run a sort on the Frontier before calling the OptimalPath
func (f Frontier) OptimalPath() ([]*FrontierVertex, Traverse) {
	return f.Values[0].Vertices, f.Values[0].Traverse
}

// Append adds the vertices onto the frontier
func (f *Frontier) Append(vertices []*FrontierVertex, cost float32, t Traverse) {
	f.Values = append(f.Values, &frontierPath{vertices, cost, t})

}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *vertices.Vertex, variable string) Frontier {
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	f := Frontier{}
	f.Append([]*FrontierVertex{fv}, 0, Visiting)
	return f
}
