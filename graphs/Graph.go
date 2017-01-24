package graphs

type Graph interface {
	Close()
	Query(cypher string) string
	Update(fn func(*GraphOperation) error) error
}

// Options for the graph
type Options struct {
	Name string
}
