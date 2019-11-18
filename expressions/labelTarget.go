package expressions

import "reflect"

type  LabelTarget struct {
	_type reflect.Type
	name string
}

func (s *LabelTarget)GetType() reflect.Type {
	return s._type
}

func (s *LabelTarget)GetName() string {
	return s.name
}