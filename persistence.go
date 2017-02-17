package graphs

//Persistence structure for saving graph data
type Persistence interface {
	Create(*[]Vertex) error
	Delete(*[]Vertex) error
	Find(string) (*Vertex, error)
	Update(*[]Vertex) error
}
