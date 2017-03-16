package cypher

type MatchVertexStatement struct {
	Variable   string
	Properties map[string]interface{}
	Label      string
}

type MatchEdgeStatement struct {
	Variable   string
	Properties map[string]interface{}
	Label      string
}
