package ir

type Pattern struct {
	Parts   []*PatternPart
}

func NewPattern() *Pattern {
	return &Pattern{
	}
}

type PatternPart struct {
	Variable   Variable
	Nodes []*Node
	Relationships []*Relationship
}

func NewPatternPart() *PatternPart {
	return &PatternPart{}
}