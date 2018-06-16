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
