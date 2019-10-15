package traversal

import (
	"errors"
	"sort"

	"github.com/google/uuid"
	"github.com/voltable/graph"
	"github.com/voltable/graph/encoding/wcs"
	"github.com/voltable/graph/operators"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/widecolumnstore"
)

var errGoalNoFound = errors.New("Goal not found")

type path struct {
	Vertices []widecolumnstore.Key
	Cost     float64
}

type frontier []*path

func (f frontier) Len() int               { return len(f) }
func (f frontier) Swap(i, j int)          { f[i], f[j] = f[j], f[i] }
func (f frontier) Less(i, j int) bool     { return f[i].Cost < f[j].Cost }
func (f frontier) pop() (*path, frontier) { return f[0], f[1:] }

func UniformCostSearch2(storage widecolumnstore.Storage, start *graph.Vertex, goal func(widecolumnstore.Key) bool) ([]uuid.UUID, error) {
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

		bytes := key.Marshal()

		predicate := func(interface{}) bool {
			return true
		}

		iterator := operators.NewFilter(predicate).Next(operators.NewScan(storage, bytes).Next())

		for kv, ok := iterator(); ok; kv, ok = iterator() {

			weight, ok := widecolumnstore.Unmarshal(kv.Value).(float64)

			if ok {
				key := widecolumnstore.Key{}
				key.Unmarshal(kv.Key)

				edge, _ := uuid.FromBytes(key.Column.Qualifier)
				tKey := prefix(edge)

				if _, ok := explored[edge]; !ok {

					frontier = append(frontier, &path{append(p.Vertices, tKey), p.Cost + weight})
				}
			}
		}
	}
}

func prefix(id uuid.UUID) widecolumnstore.Key {
	return widecolumnstore.NewKey(id[:], &widecolumnstore.Column{Family: wcs.Relationship})
}

// TODO need this signature to work for UCS
func UniformCostSearch(graph query.Graph, operator operators.Operator, frontier *query.Frontier) bool {
	return false
}
