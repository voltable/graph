package query

import (
	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

type FrontierQueue struct {
	Parts []FrontierProperties
	Cost  float64
}

// FrontierProperties containers a KeyValue (vertex or edge) and it's Variable used by a query
type FrontierProperties struct {
	KeyValue *keyvalue.KeyValue
	Variable string
	UUID     uuid.UUID
}

// Frontier priority queue containing vertices to be explored and the cost for a Uniform Cost Search
type Frontier struct {
	Values   []*FrontierQueue
	Explored map[uuid.UUID]bool
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
func (f *Frontier) append(vertices []FrontierProperties, cost float64) {
	fp := &FrontierQueue{vertices, cost}
	f.Values = append(f.Values, fp)
}

func (f *Frontier) AppendKeyValue(queue *FrontierQueue, v *keyvalue.KeyValue, variable string) {
	fv := FrontierProperties{KeyValue: v, Variable: variable, UUID: v.UUID()}
	f.append(append(queue.Parts, fv), queue.Cost)
}

// NewFrontier create the Frontier using the inistal Vertex as the root of the graph
func NewFrontier(v *keyvalue.KeyValue, variable string) Frontier {
	fv := FrontierProperties{KeyValue: v, Variable: variable, UUID: v.UUID()}
	f := Frontier{Explored: make(map[uuid.UUID]bool)}
	f.append([]FrontierProperties{fv}, 0)
	return f
}
