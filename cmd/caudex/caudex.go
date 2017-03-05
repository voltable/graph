package main

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func main() {
	var g graph.Graph
	var err error
	if g, err = graph.NewGraph("bolt", &graph.Options{"test"}); err != nil {
		panic(err)
	}

	//v := vertices.N
	//g.Create()

	var results []*vertices.Vertex
	g.Query().Where(func(v *vertices.Vertex) bool {
		return v.Label() == "foo"
	}).ToSlice(results)

	g.Close()
}
