package graphs

import "bitbucket.org/rossmerr/caudex/graphs/internal"

type VertexEnvelop struct {
	Vertices *[]Vertex
	State    internal.State
}
