package store64_test

import (
	"testing"

	triples "github.com/RossMerr/Caudex.Graph/triplestore/store64"
	proto "github.com/golang/protobuf/proto"
)

func Test_Proto(t *testing.T) {
	v := float64(5)

	triple := &triples.Triple{Row: "1", Column: "2", Value: v}

	in, err := proto.Marshal(triple)
	if err != nil {
		t.Errorf("Failed to encode address book: %+v", err)
	}

	out := &triples.Triple{}
	if err := proto.Unmarshal(in, out); err != nil {
		t.Errorf("Failed to parse address book: %+v", err)
	}

	if out.Value != v {
		t.Errorf("%+v got %+v, want %+v", "float64", out.Value, v)
	}

}
