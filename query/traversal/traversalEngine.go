package traversal

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Query() *query.VertexPath {
	//t, _ := NewTraversal("dfs")
	//t.Query(root)
	//todo need to setup channel from DFS or BFS

	c := make(chan *vertices.Vertex)
	return &query.Query{
		Iterate: func() query.Iterator {
			return func() (item *vertices.Vertex, ok bool) {
				v, ok := <-c
				return v, ok
			}
		},
	}

	return nil
}
