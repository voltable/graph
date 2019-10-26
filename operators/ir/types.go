package ir

type Variable string

type Type string

type Label string

type Key string

type Arithmetic string

const (
	Addition = Arithmetic("+")
	Subtraction = Arithmetic("-")
	Multiplication  = Arithmetic("*")
	Division  = Arithmetic("/")
)

type String string

func (s String) String() string {
	return "'"+string(s)+"'"
}

