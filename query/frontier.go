package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type frontierPath struct {
	Vertices []*FrontierVertex
	Cost     float32
}

// FrontierVertex containers a vertex and it's Variable
type FrontierVertex struct {
	*vertices.Vertex
	Variable string
}

// Frontier priority queue containing vertices to be explored and the cost for a Uniform Cost Search
type Frontier struct {
	Values   []*frontierPath
	Traverse Traverse
}

// Sort interface
func (f Frontier) Len() int           { return len(f.Values) }
func (f Frontier) Swap(i, j int)      { f.Values[i], f.Values[j] = f.Values[j], f.Values[i] }
func (f Frontier) Less(i, j int) bool { return f.Values[i].Cost < f.Values[j].Cost }

func (f Frontier) peek() []*FrontierVertex { return f.Values[0].Vertices }
func (f Frontier) Peek() *FrontierVertex   { return f.Values[0].Vertices[0] }

func (f *Frontier) Pop() (vertices []*FrontierVertex, cost float32) {
	vertices = f.Values[0].Vertices
	cost = f.Values[0].Cost
	f.Values = f.Values[1:]
	return
}

// OptimalPath returns what should be the optimal path
func (f Frontier) OptimalPath() []*FrontierVertex {
	return f.Values[0].Vertices
}

// Append adds the vertices onto the frontier
func (f *Frontier) Append(vertices []*FrontierVertex, cost float32) {
	f.Values = append(f.Values, &frontierPath{vertices, cost})

}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *vertices.Vertex, variable string) Frontier {
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	f := Frontier{}
	f.Append([]*FrontierVertex{fv}, 0)
	return f
}
