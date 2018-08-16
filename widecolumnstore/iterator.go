package widecolumnstore

import "github.com/RossMerr/Caudex.Graph/uuid"

// Iterator is an alias for function to iterate over data.
type Iterator func() (interface{}, bool)

// IteratorUUID is an alias for function to iterate over uuid's.
type IteratorUUID func() (*uuid.UUID, bool)

// IteratorUUIDWeight is an alias for function to iterate over uuid's with weight
type IteratorUUIDWeight func() (*uuid.UUID, float64, bool)
