package main

import (
	"github.com/RossMerr/Caudex.Graph"
)

func main() {
	var g graph.Graph
	var err error
	if g, err = graph.NewGraph("bolt", &graph.Options{"test"}); err != nil {
		g.Close()
	}

	//g.Query().Where()

}
