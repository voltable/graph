// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package keyvalue

import "github.com/RossMerr/Caudex.Graph/arch"

// NewAny returns a keyvalue.Any
func NewAny(p interface{}) *Any {
	t, v := arch.EncodeType(p)
	return &Any{
		TypeUrl: t,
		Value:   v,
	}
}

// Unmarshal returns the Value as it's type defined by the TypeUrl
func (s *Any) Unmarshal() interface{} {
	return arch.DecodeType(s.TypeUrl, s.Value)
}
