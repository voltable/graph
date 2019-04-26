package query

import (
	"github.com/voltable/graph/uuid"
	"github.com/voltable/graph/widecolumnstore"
)

// RelationshipPrefix builds up the key for a prefix search to find any relationship
func RelationshipPrefix(id *uuid.UUID) []byte {
	return append(id[:], append(widecolumnstore.US, Relationship...)...)
}
