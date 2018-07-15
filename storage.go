package graph

type Storage interface {
	ForEach() Iterator
	Fetch(string) (*Vertex, error)
	HasPrefix([]byte) Iterator
}
