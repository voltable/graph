package graphs

type (
	// Graph structure consisting of vertices and edges
	Graph interface {
		Close()
		Query(fn func(*QueryOperation) error) string
		Command(fn func(*GraphOperation) error) error
	}

	// Options for the graph
	Options struct {
		Name string
	}
)
