package main

import (
	"fmt"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func main() {
	var g graph.Graph
	var err error
	if g, err = graph.NewGraph("bolt", &graph.Options{"test"}); err != nil {
		panic(err)
	}

	slice := g.Query().Match(func(v *vertices.Vertex) bool {
		return v.Label() == "foo"
	}).Match(func(e *vertices.Edge) bool {
		return true
	}).ToSliceAll()

	for _, v := range slice {
		fmt.Println(v)
	}

	g.Close()
}
