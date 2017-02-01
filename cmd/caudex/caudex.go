package main

import (
	"github.com/RossMerr/Caudex.Graph/graphs"
	"github.com/RossMerr/Caudex.Graph/graphs/boltdb"
)

func main() {
	g := boltdb.Open(&graphs.Options{"test"})
	g.Query("MATCH (node:Label) RETURN node.property")
	g.Close()
}
