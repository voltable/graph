package traversal

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

// TODO fix this next to get the queryEngine_test's working
func UniformCostSearch(graph query.Graph, operator widecolumnstore.Operator, frontier *query.Frontier) bool {

	if frontier.Len() > 0 {
		queue := frontier.Pop()
		depth := len(queue.Parts)
		part := queue.Parts[depth-1]

		if _, ok := frontier.Explored[part.UUID]; !ok {
			frontier.Explored[part.UUID] = true
			frontier.AppendKeyValue(queue, part.UUID, part.Variable)
			sort.Sort(frontier)
		}

		//	if pe := predicates[depth]; pe != nil {
		//graph.HasPrefix()
		// iterator := graph.Edges(part.UUID)
		// for kv, weight, hasEdges := iterator(); hasEdges; kv, weight, hasEdges = iterator() {
		// 	if _, ok := frontier.Explored[kv]; !ok {
		// 		if variable, p := pe(part.UUID, kv, depth); p == query.Visiting || p == query.Matching {
		// 			frontier.AppendEdgeKeyValue(queue, kv, variable, weight)
		// 		}
		// 	}
		// }
		//	}
	}

	// if frontier.Len() > 0 {
	// 	queue := frontier.Pop()
	// 	depth := len(queue.Parts)
	// 	searchDepth := len(predicates)
	// 	part := queue.Parts[depth-1]

	// 	if _, ok := frontier.Explored[part.UUID]; !ok {
	// 		frontier.Explored[part.UUID] = true
	// 		pv := predicates[depth-1]
	// 		if variable, p := pv(part.UUID, nil, depth-1); p == query.Matched {
	// 			queue.Parts[depth-1].Variable = variable
	// 			frontier.AppendKeyValue(queue, part.UUID, part.Variable)
	// 			sort.Sort(frontier)
	// 			return depth == searchDepth
	// 		}

	// 	}

	// 	if depth >= searchDepth {
	// 		return false
	// 	}

	// 	if pe := predicates[depth]; pe != nil {
	// 		iterator := graph.Edges(part.UUID)
	// 		for kv, weight, hasEdges := iterator(); hasEdges; kv, weight, hasEdges = iterator() {
	// 			if _, ok := frontier.Explored[kv]; !ok {
	// 				if variable, p := pe(part.UUID, kv, depth); p == query.Visiting || p == query.Matching {
	// 					frontier.AppendEdgeKeyValue(queue, kv, variable, weight)
	// 				}
	// 			}
	// 		}
	// 	}

	// 	sort.Sort(frontier)
	// }
	return false
}
