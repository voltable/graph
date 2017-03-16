package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type VertexToRelationshipPath struct {
	Iterate  func() Iterator
	Explored map[string]bool
	Fetch    func(string) (*vertices.Vertex, error)
}

type RelationshipType int

const (
	Inbound    RelationshipType = iota
	Outbound   RelationshipType = iota
	Undirected RelationshipType = iota
)

func (r *VertexToRelationshipPath) Inbound() *EdgePath {
	return nil
}

func (r *VertexToRelationshipPath) Outbound() *EdgePath {
	return nil
}

func (r *VertexToRelationshipPath) Undirected() *EdgePath {
	return nil
}
