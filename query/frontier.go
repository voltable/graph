package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type frontierPath struct {
	Vertices []*vertices.Vertex
	Cost     float32
}

// Frontier priority queue containing vertices to be explored
type Frontier []*frontierPath

func (f Frontier) Len() int                 { return len(f) }
func (f Frontier) Swap(i, j int)            { f[i], f[j] = f[j], f[i] }
func (f Frontier) Less(i, j int) bool       { return f[i].Cost < f[j].Cost }
func (f Frontier) peek() []*vertices.Vertex { return f[0].Vertices }

func (f Frontier) pop() ([]*vertices.Vertex, float32, Frontier) {
	return f[0].Vertices, f[0].Cost, f[1:]
}

// Append adds the vertices onto the frontier
func (f Frontier) Append(vertices []*vertices.Vertex, cost float32) Frontier {
	f = append(f, &frontierPath{vertices, cost})
	return f
}

func NewFrontier(v *vertices.Vertex) Frontier {
	f := Frontier{}
	f = f.Append([]*vertices.Vertex{v}, 0)
	return f
}
