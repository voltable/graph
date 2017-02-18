package graphs

//StorageEngine structure for saving graph data
type StorageEngine interface {
	Create([]Vertex) error
	Delete([]Vertex) error
	Find(string) (*Vertex, error)
	Update([]Vertex) error
}
