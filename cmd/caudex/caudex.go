package main

import "github.com/RossMerr/Caudex.Graph/storageEngines/boltdb"

func main() {
	g := boltdb.BuildGraph()
	g.Query("MATCH (node:Label) RETURN node.property")
	g.Close()
}
