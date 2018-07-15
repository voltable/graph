package ast

type Patn struct {
	Variable      string
	Properties    map[string]interface{}
	Label         string
	Relationship  Digraph
	LengthMinimum uint
	LengthMaximum uint
	Body          *Patn
	Next          *Patn
}

func (*Patn) patnNode() {}
