package graphs

type Persistence interface {
	Create(*[]Vertex) error
	Delete(*[]Vertex) error
	Find(string) (*Vertex, error)
	Update(*[]Vertex) error
}
