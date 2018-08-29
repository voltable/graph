package operators

// Project is a set operator that projects a set of tuples onto the specified attributes
type Project struct {
}

func (s *Project) Next() (interface{}, bool) {
	return nil, false
}

// func (s *Project) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
// 	iterator := s.iterator.Next(i)
// 	return func() (interface{}, bool) {
// 		for value, ok := iterator(); ok; value, ok = iterator() {
// 			if s.predicate(value) {
// 				return value, true
// 			}
// 		}
// 		return nil, false
// 	}
// }

func (s *Project) Op() {}
