package traversal

import (
	"errors"
	"fmt"
	"sort"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

var errGoalNoFound = errors.New("Goal not found")

type path struct {
	Vertices []uuid.UUID
	Cost     float64
}

type frontier []*path

func (f frontier) Len() int               { return len(f) }
func (f frontier) Swap(i, j int)          { f[i], f[j] = f[j], f[i] }
func (f frontier) Less(i, j int) bool     { return f[i].Cost < f[j].Cost }
func (f frontier) pop() (*path, frontier) { return f[0], f[1:] }

// TODO fix this next to get the queryEngine_test's working
func UniformCostSearch2(g *query.Graph, start *graph.Vertex, goal func(uuid.UUID) bool) ([]uuid.UUID, error) {
	frontier := frontier{&path{[]uuid.UUID{*start.ID()}, 0}}
	explored := make(map[uuid.UUID]bool)
	for {
		if len(frontier) == 0 {
			return nil, errGoalNoFound
		}

		var p *path
		sort.Sort(frontier)
		p, frontier = frontier.pop()
		node := p.Vertices[len(p.Vertices)-1]
		explored[node] = true

		if goal(node) {
			return p.Vertices, nil
		}

		iterator := g.Edges(node)
		for kv, ok := iterator(); ok; kv, ok = iterator() {
			id, weight := query.UnmarshalKeyValueTransposeTRelationship(kv)
			if _, ok := explored[id]; !ok {
				frontier = append(frontier, &path{append(p.Vertices, id), p.Cost + weight})
			} else {
				fmt.Printf("skip: %+v\n", id)
			}
		}

		// for id, weight, ok := iterator(); ok; id, weight, ok = iterator() {
		// 	if _, ok := explored[id]; !ok {
		// 		frontier = append(frontier, &path{append(p.Vertices, id), p.Cost + weight})
		// 	} else {
		// 		fmt.Printf("skip: %+v\n", id)
		// 	}
		// }
	}
}

func UniformCostSearch(graph query.Graph, operator widecolumnstore.Operator, frontier *query.Frontier) bool {
	return false
}

// 	if frontier.Len() > 0 {
// 		queue := frontier.Pop()
// 		depth := len(queue.Parts)
// 		part := queue.Parts[depth-1]

// 		if _, ok := frontier.Explored[part.UUID]; !ok {
// 			frontier.Explored[part.UUID] = true
// 			frontier.AppendKeyValue(queue, part.UUID, part.Variable)
// 			sort.Sort(frontier)
// 		}

// 		//	if pe := predicates[depth]; pe != nil {
// 		//graph.HasPrefix()
// 		// iterator := graph.Edges(part.UUID)
// 		// for kv, weight, hasEdges := iterator(); hasEdges; kv, weight, hasEdges = iterator() {
// 		// 	if _, ok := frontier.Explored[kv]; !ok {
// 		// 		if variable, p := pe(part.UUID, kv, depth); p == query.Visiting || p == query.Matching {
// 		// 			frontier.AppendEdgeKeyValue(queue, kv, variable, weight)
// 		// 		}
// 		// 	}
// 		// }
// 		//	}
// 	}

// 	// if frontier.Len() > 0 {
// 	// 	queue := frontier.Pop()
// 	// 	depth := len(queue.Parts)
// 	// 	searchDepth := len(predicates)
// 	// 	part := queue.Parts[depth-1]

// 	// 	if _, ok := frontier.Explored[part.UUID]; !ok {
// 	// 		frontier.Explored[part.UUID] = true
// 	// 		pv := predicates[depth-1]
// 	// 		if variable, p := pv(part.UUID, nil, depth-1); p == query.Matched {
// 	// 			queue.Parts[depth-1].Variable = variable
// 	// 			frontier.AppendKeyValue(queue, part.UUID, part.Variable)
// 	// 			sort.Sort(frontier)
// 	// 			return depth == searchDepth
// 	// 		}

// 	// 	}

// 	// 	if depth >= searchDepth {
// 	// 		return false
// 	// 	}

// 	// 	if pe := predicates[depth]; pe != nil {
// 	// 		iterator := graph.Edges(part.UUID)
// 	// 		for kv, weight, hasEdges := iterator(); hasEdges; kv, weight, hasEdges = iterator() {
// 	// 			if _, ok := frontier.Explored[kv]; !ok {
// 	// 				if variable, p := pe(part.UUID, kv, depth); p == query.Visiting || p == query.Matching {
// 	// 					frontier.AppendEdgeKeyValue(queue, kv, variable, weight)
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}

// 	// 	sort.Sort(frontier)
// 	// }
// 	return false
// }
