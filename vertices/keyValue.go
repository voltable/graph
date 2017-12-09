package vertices

import "fmt"

type KeyValue struct {
	Key   string
	Value interface{}
}

func (b KeyValue) String() string {
	return fmt.Sprintf("{%v => %#v}", b.Key, b.Value)
}
