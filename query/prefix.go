package query

import (
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

// RelationshipPrefix builds up the key for a prefix search to find any relationship
func RelationshipPrefix(id *uuid.UUID) []byte {
	return append(id[:], append(widecolumnstore.US, Relationship...)...)
}
