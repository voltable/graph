package main

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
)

func main() {
	if engine, err := storageEngines.NewStorageEngine("", &graphs.Options{"test"}); err == nil {
		g := graphs.NewGraph(engine)

		//g.Query("MATCH (node:Label) RETURN node.property")
		g.Close()
	}
}
