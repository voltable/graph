package ir

import (
	"bytes"

	"github.com/voltable/graph/expressions"
)

type ListLiteral struct {
	Items []expressions.Expression
}

func (s ListLiteral) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")
	max := len(s.Items) - 1
	for i, item := range s.Items {
		buffer.WriteString(item.String())
		if i != max {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}
