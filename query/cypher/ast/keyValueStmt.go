package ast

import "fmt"

// KeyValueStmt a key/value
type KeyValueStmt struct {
	Key   string
	Value interface{}
}

func (b KeyValueStmt) String() string {
	return fmt.Sprintf("{%v => %#v}", b.Key, b.Value)
}
