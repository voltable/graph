package widecolumnstore

import (
	"github.com/google/uuid"
)

// ParseKeyToUUID looks for the UUID in the Key
func ParseKeyToUUID(key Key) (uuid.UUID, error) {
	return uuid.FromBytes(key.RowKey)
}

// UUID looks for the UUID in the KeyValue
func UUID(s *KeyValue) (uuid.UUID, error) {
	return uuid.FromBytes(s.Key.RowKey)
}
