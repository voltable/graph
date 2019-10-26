package ir

type Properties struct {
	Map *MapLiteral
}

func NewProperties() *Properties {
	return &Properties{
		Map: NewMapLiteral(),
	}
}