package widecolumnstore

import (
	"time"
)

type crud int

const (
	put crud = iota
)

type Mutation struct {
	operastion map[time.Time]*action
}

type action struct {
	kv     *KeyValue
	action crud
}

func (s *Mutation) Put(kv *KeyValue) {
	s.operastion[time.Now()] = &action{
		kv:     kv,
		action: put,
	}
}
