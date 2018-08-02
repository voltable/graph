package query

import "github.com/RossMerr/Caudex.Graph/uuid"

type Storage interface {
	ForEach() IteratorUUID
	HasPrefix([]byte) Iterator
	Edges(*uuid.UUID) IteratorUUIDWeight
}
