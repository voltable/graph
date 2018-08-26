package widecolumnstore

import (
	"bytes"
	"time"
)

// NewKey returns a new Key
func NewKey(id []byte, column *Column) *Key {
	return &Key{
		ID:     id,
		Column: column,
	}
}

// Key of the map, broken up logically into a few different parts
type Key struct {
	// ID Unique identifier for the row.
	ID []byte
	// Column Logical grouping of the key
	Column *Column
	// Timestamp for serializability
	Timestamp time.Time
}

// Column Logical grouping of the key
type Column struct {
	// Family This field can be used to partition data within a node.
	Family []byte
	// Extended This is for extended family for deeper partition data.
	Extended []byte
	// Qualifier More specific attribute of the key.
	Qualifier []byte
}

// Marshal a key into bytes
func (s *Key) Marshal() (key []byte) {
	key = append(key, s.ID...)
	key = append(key, US...)
	key = append(key, s.Column.Family...)
	key = append(key, US...)
	key = append(key, s.Column.Extended...)
	key = append(key, US...)
	key = append(key, s.Column.Qualifier...)
	return
}

// Unmarshal bytes into a key
func (s *Key) Unmarshal(key []byte) {

	split := bytes.Split(key, US)

	if len(split) > 0 {
		s.ID = split[0]

		if s.Column == nil {
			s.Column = &Column{}
		}
		if len(split) >= 3 {
			s.Column.Family = split[1]
			s.Column.Extended = split[2]
			s.Column.Qualifier = split[3]
		}
	}
	return
}
