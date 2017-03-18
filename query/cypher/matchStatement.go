package cypher

type MatchVertexStatement struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Edge *MatchEdgeStatement
}

type MatchEdgeStatement struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Vertex *MatchVertexStatement
}

type MatchStatement struct {
	Variable   string
	Properties map[string]interface{}
	Label      string
}
