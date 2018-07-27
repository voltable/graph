package keyvalue

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/uuid"
)

// NewKeyValue returns a new KeyValue
func NewKeyValue(i interface{}, bytes ...[]byte) *KeyValue {
	kv := &KeyValue{
		Value: NewAny(i),
	}
	for _, b := range bytes {
		kv.Key = append(kv.Key, b...)
	}
	return kv
}

// UUID looks for the UUID in the KeyValue
func (s *KeyValue) UUID() uuid.UUID {
	split := bytes.Split(s.Key, US)
	if bytes.Equal(split[1], Vertex) {
		return uuid.SliceToUUID(split[0])
	}

	if bytes.Equal(split[1], Properties) {
		return uuid.SliceToUUID(split[0])
	}

	if bytes.Equal(split[1], Relationship) {
		return uuid.SliceToUUID(split[0])
	}

	if bytes.Equal(split[1], Relationshipproperties) {
		return uuid.SliceToUUID(split[0])
	}

	// Transpose

	if bytes.Equal(split[0], Vertex) {
		return uuid.SliceToUUID(split[2])
	}

	if bytes.Equal(split[0], Properties) {
		return uuid.SliceToUUID(split[2])
	}

	if bytes.Equal(split[0], Relationship) {
		return uuid.SliceToUUID(split[3])
	}

	if bytes.Equal(split[0], Relationshipproperties) {
		return uuid.SliceToUUID(split[3])
	}

	return uuid.UUID{}
}
