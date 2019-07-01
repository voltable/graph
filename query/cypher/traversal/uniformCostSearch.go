package traversal

import (
	"errors"
	"fmt"
	"sort"

	graph "github.com/voltable/graph"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/uuid"
	"github.com/voltable/graph/widecolumnstore"
)

var errGoalNoFound = errors.New("Goal not found")

type Filter struct {
	Storage widecolumnstore.HasPrefix
	ID      uuid.UUID
}

func (s *Filter) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	key := prefix(s.ID)
	bytes := key.Marshal()
	return s.Storage.HasPrefix(bytes)
}

func (s *Filter) Op() {}

type path struct {
	Vertices []widecolumnstore.Key
	Cost     float64
}

type frontier []*path

func (f frontier) Len() int               { return len(f) }
func (f frontier) Swap(i, j int)          { f[i], f[j] = f[j], f[i] }
func (f frontier) Less(i, j int) bool     { return f[i].Cost < f[j].Cost }
func (f frontier) pop() (*path, frontier) { return f[0], f[1:] }

// TODO fix this next to get the queryEngine_test's working
// TODO I think the input f operators.Filter should be a operators.Scan
func UniformCostSearch2(f *Filter, start *graph.Vertex, goal func(widecolumnstore.Key) bool) ([]uuid.UUID, error) {
	root := prefix(start.ID())
	frontier := frontier{&path{[]widecolumnstore.Key{root}, 0}}
	explored := make(map[uuid.UUID]bool)

	for {
		if len(frontier) == 0 {
			return nil, errGoalNoFound
		}

		var p *path
		sort.Sort(frontier)
		p, frontier = frontier.pop()
		key := p.Vertices[len(p.Vertices)-1]
		id, _ := query.ParseKeyToUUID(key)
		explored[id] = true

		if goal(key) {
			results := make([]uuid.UUID, 0)
			for _, key := range p.Vertices {
				id, _ := query.ParseKeyToUUID(key)
				results = append(results, id)
			}
			return results, nil
		}

		fmt.Printf("edges: %+v\n", id)

		// TODO need to use operators.Filter
		filter := Filter{Storage: f.Storage, ID: id}

		iterator := filter.Next(nil)

		for kv, ok := iterator(); ok; kv, ok = iterator() {

			weight, ok := widecolumnstore.Unmarshal(kv.Value).(float64)

			if ok {
				key := widecolumnstore.Key{}
				key.Unmarshal(kv.Key)
				tKey := TransposeRelationship(key)
				edge := uuid.SliceToUUID(key.Column.Qualifier)

				if _, ok := explored[edge]; !ok {
					fmt.Printf("add: %+v\n", edge)
					frontier = append(frontier, &path{append(p.Vertices, tKey), p.Cost + weight})
				} else {
					fmt.Printf("skip: %+v\n", edge)
				}
			}
		}
	}
}

func TransposeRelationship(key widecolumnstore.Key) widecolumnstore.Key {
	return widecolumnstore.NewKey(key.Column.Qualifier, &widecolumnstore.Column{query.TRelationship, key.Column.Extended, key.ID})
}

func prefix(id uuid.UUID) widecolumnstore.Key {
	return widecolumnstore.NewKey(id[:], &widecolumnstore.Column{Family: query.Relationship})
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
