package expressions

import "reflect"

func AreReferenceAssignable(dest interface{}, src interface{}) bool {
	// TODO need to validate this is right behaviour
	return reflect.TypeOf(dest).AssignableTo(reflect.TypeOf(src))
}
