// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package keyvalue

func (s *Any) Decode() interface{} {
	return DecodeBytes(s.TypeUrl, s.Value)
}
