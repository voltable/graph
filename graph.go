package graphs

type (
	Graph interface {
		Close()
		Query(cypher string) string
		Command(fn func(*GraphOperation) error) error
	}

	// Options for the graph
	Options struct {
		Name string
	}
)
