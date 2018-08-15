package keyvaluestore

import (
	"github.com/RossMerr/Caudex.Graph/uuid"
)

type Storage interface {
	Each() Iterator
	ForEach() IteratorUUID
	HasPrefix([]byte) Iterator
	Edges(*uuid.UUID) IteratorUUIDWeight
}
