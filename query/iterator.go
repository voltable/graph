package query

import "github.com/RossMerr/Caudex.Graph/uuid"

// IteratorUUIDWeight is an alias for function to iterate over uuid's with weight
type IteratorUUIDWeight func() (*uuid.UUID, float64, bool)

type IteratorUUID func() (*uuid.UUID, bool)
