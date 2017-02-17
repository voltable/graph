package main

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/persistences/boltdb"
)

func main() {
	g := boltdb.Open(&graphs.Options{"test"})
	g.Query("MATCH (node:Label) RETURN node.property")
	g.Close()
}
