package query

import (
	"github.com/RossMerr/Caudex.Graph/uuid"
)

// IteratorFrontier is an alias for function to iterate over Frontier.
type IteratorFrontier func() (item *Frontier, ok bool)

type FrontierQueue struct {
	Parts []FrontierProperties
	Cost  float64
}

// FrontierProperties containers a KeyValue (vertex or edge) and it's Variable used by a query
type FrontierProperties struct {
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

func (f *Frontier) AppendKeyValue(queue *FrontierQueue, id uuid.UUID, variable string) {
	fv := FrontierProperties{Variable: variable, UUID: id}
	f.append(append(queue.Parts, fv), queue.Cost)
}

// NewFrontier create the Frontier using the inistal uuid as the root of the graph
func NewFrontier(id uuid.UUID, variable string) Frontier {
	fv := FrontierProperties{Variable: variable, UUID: id}
	f := Frontier{Explored: make(map[uuid.UUID]bool)}
	f.append([]FrontierProperties{fv}, 0)
	return f
}
