package main

import (
	"bitbucket.org/rossmerr/caudex/graphs"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb"
)

func main() {
	g := boltdb.Open(&graphs.Options{"test"})
	g.Query("MATCH (node:Label) RETURN node.property")
	g.Close()
}
