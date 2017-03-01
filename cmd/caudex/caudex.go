package main

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
)

func main() {
	g := graphs.NewGraph(storageEngines.BuildGraphDefault(&graphs.Options{"test"}))

	//g.Query("MATCH (node:Label) RETURN node.property")
	g.Close()
}
