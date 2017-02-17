package memorydb

import "github.com/rossmerr/Caudex.Graph"

type Graph struct {
}

func (g *Graph) Close() {

}
func (g *Graph) Query(cypher string) string {
	return ""
}

func (g *Graph) Command(fn func(*graphs.GraphOperation) error) error {
	return nil
}
