package main

import "bitbucket.org/rossmerr/caudex/graph"

func main() {
	var graph = caudex.Open(&caudex.Options{"test"})	
	caudex.Query(graph, "MATCH (node:Label) RETURN node.property")
	caudex.Close(graph)
}


