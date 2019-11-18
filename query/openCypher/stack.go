package openCypher

// StackExpr a simple stack for the AST
type StackExpr []interface{}

// Push add's a item the the StackExpr
func (s StackExpr) push(v interface{}) StackExpr {
	return append(s, v)
}

// pop removes the last item on the StackExpr and returns it
func (s StackExpr) pop() (StackExpr, interface{}) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1]
	}
	return s, nil
}

// pop removes the first item on the StackExpr and returns it
//func (s StackExpr) shift() (StackExpr, interface{}) {
//	l := len(s)
//	if l > 0 {
//		return s[1:], s[0]
//	}
//	return s, nil
//}

// top returns the last item on the StackExpr without removing it
func (s StackExpr) top()interface{} {
	l := len(s)
	if l > 0 {
		return s[l-1]
	}
	return nil
}