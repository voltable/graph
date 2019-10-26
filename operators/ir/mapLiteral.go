package ir

import "bytes"

type MapLiteral struct {
	Items map[Key]*Expression
}

func NewMapLiteral() *MapLiteral {
	return &MapLiteral{
		Items: make(map[Key]*Expression, 0),
	}
}

func (s MapLiteral) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	max := len(s.Items) - 1
	i := 0
	for key, value := range s.Items {
		buffer.WriteString(string(key))
		buffer.WriteString(": ")
		buffer.WriteString(value.String())
		if i != max {
			buffer.WriteString(", ")
		}
		i++
	}

	buffer.WriteString("}")
	return buffer.String()
}