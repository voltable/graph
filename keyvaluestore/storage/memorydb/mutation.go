package memorydb

import (
	"time"

	"github.com/RossMerr/Caudex.Graph/keyvaluestore"
)

type crud int

const (
	put crud = iota
)

type Mutation struct {
	operastion map[time.Time]*action
}

type action struct {
	kv     *keyvaluestore.KeyValue
	action crud
}

func (s *Mutation) Put(kv *keyvaluestore.KeyValue) {
	s.operastion[time.Now()] = &action{
		kv:     kv,
		action: put,
	}
}
