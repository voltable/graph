package keyvalue

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func (s *KeyValue) Weight() float64 {
	split := bytes.Split(s.Key, RS)
	if bytes.Equal(split[1], Relationship) {
		to := Unmarshal(s.Value)
		return to.(float64)
	}

	if bytes.Equal(split[0], Relationship) {
		return arch.DecodeFloat64Bytes(split[3]).(float64)
	}

	return 0
}

func (s *KeyValue) To() *uuid.UUID {
	split := bytes.Split(s.Key, RS)
	if bytes.Equal(split[1], Relationship) {
		return uuid.SliceToUUID(split[3])
	}

	if bytes.Equal(split[0], Relationship) {
		to := Unmarshal(s.Value)
		return to.(*uuid.UUID)
	}

	return nil
}

// UUID looks for the UUID in the KeyValue
func (s *KeyValue) UUID() *uuid.UUID {
	key := &Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Label) {
		return uuid.SliceToUUID(key.ID)
	}

	if bytes.Equal(key.Column.Family, Properties) {
		return uuid.SliceToUUID(key.ID)
	}

	subSplit := bytes.Split(key.Column.Family, US)

	if len(subSplit) > 0 {
		if bytes.Equal(subSplit[0], Relationship) {
			return uuid.SliceToUUID(key.ID)
		}

		if bytes.Equal(subSplit[0], Relationshipproperties) {
			return uuid.SliceToUUID(key.ID)
		}
	}

	// Transpose

	if bytes.Equal(key.ID, TLabel) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, TProperties) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	subSplit = bytes.Split(key.ID, US)

	if len(subSplit) > 0 {
		if bytes.Equal(subSplit[0], TRelationship) {
			return uuid.SliceToUUID(key.Column.Qualifier)
		}

		if bytes.Equal(subSplit[0], TRelationshipproperties) {
			return uuid.SliceToUUID(key.Column.Qualifier)
		}
	}

	return nil
}

func (s *KeyValue) Interpret(value string) interface{} {
	split := bytes.Split(s.Key, RS)
	if len(split) > 1 {
		if bytes.Equal(split[1], Properties) {
			if value == string(split[2]) {
				return Unmarshal(s.Value)
			}
		}
	}

	return nil
}
