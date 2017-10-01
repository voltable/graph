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
type Frontier []*frontierPath

func (f Frontier) Len() int                { return len(f) }
func (f Frontier) Swap(i, j int)           { f[i], f[j] = f[j], f[i] }
func (f Frontier) Less(i, j int) bool      { return f[i].Cost < f[j].Cost }
func (f Frontier) peek() []*FrontierVertex { return f[0].Vertices }
func (f Frontier) Peek() *FrontierVertex   { return f[0].Vertices[0] }

func (f Frontier) Pop() ([]*FrontierVertex, float32, Frontier) {
	return f[0].Vertices, f[0].Cost, f[1:]
}

// TODO need todo somthing to remove deadends for explored frontierPath's
// Append adds the vertices onto the frontier
func (f Frontier) Append(vertices []*FrontierVertex, cost float32) Frontier {
	f = append(f, &frontierPath{vertices, cost})
	return f
}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *vertices.Vertex) Frontier {
	fv := &FrontierVertex{Vertex: v}
	f := Frontier{}
	f = f.Append([]*FrontierVertex{fv}, 0)
	return f
}

// NewFrontierVariable create the Frontier using the inistal Vertex as the root of the graph
func NewFrontierVariable(v *vertices.Vertex, variable string) Frontier {
	fv := &FrontierVertex{Vertex: v, Variable: variable}
	f := Frontier{}
	f = f.Append([]*FrontierVertex{fv}, 0)
	return f
}
