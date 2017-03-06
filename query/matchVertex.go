package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// MatchVertex returns all Verteces matching the predicate
func (t VertexPath) MatchVertex(predicate func(v *vertices.Vertex) bool) EdgePath {
	return EdgePath{
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if v, is := item.(*vertices.Vertex); is {
						if predicate(v) {
							for _, e := range v.Edges() {
								return e, true
							}
						}
					}
				}
				return
			}
		},
	}
}

// Iterate: func() Iterator {
// 	index := 0
// 	keys := src.MapKeys()

// 	return func() (item interface{}, ok bool) {
// 		ok = index < len
// 		if ok {
// 			key := keys[index]
// 			item = KeyValue{
// 				Key:   key.Interface(),
// 				Value: src.MapIndex(key).Interface(),
// 			}

// 			index++
// 		}

// 		return
// 	}
// },
