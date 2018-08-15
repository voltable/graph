package keyvaluestore

import "github.com/RossMerr/Caudex.Graph/arch"
import any "github.com/golang/protobuf/ptypes/any"

// NewAny returns a keyvalue.Any
func NewAny(p interface{}) *any.Any {
	t, v := arch.EncodeType(p)
	return &any.Any{
		TypeUrl: t,
		Value:   v,
	}
}

// Unmarshal returns the Value as it's type defined by the TypeUrl
func Unmarshal(s *any.Any) interface{} {
	return arch.DecodeType(s.TypeUrl, s.Value)
}
