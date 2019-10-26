package ir

import "fmt"

type Expression struct {
	Value interface{}
}

func (s *Expression) Evaluate() []interface{} {
	rows := make([]interface{}, 0)
	rows = append(rows, s.Value)
	return rows
}

func (s Expression) String() string {
	return  fmt.Sprint(s.Value)
}