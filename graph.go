package graph



//Graph structure for saving graph data
type Graph interface {
	Create(...*Vertex) error
	Delete(...*Vertex) error
	Update(...*Vertex) error
	Close()
	Query(string) (*Query, error)
}
