package query

import "github.com/RossMerr/Caudex.Graph/widecolumnstore"

var (
	ID                     = []byte("i")
	Label                  = []byte("l")
	Properties             = []byte("p")
	Relationship           = []byte("r")
	Relationshipproperties = []byte("k")

	// T = transpose
	TID                     = []byte("ti")
	TLabel                  = []byte("tl")
	TProperties             = []byte("tp")
	TRelationship           = []byte("tr")
	TRelationshipproperties = []byte("tk")

	// This Key can be used to fetch all nodes, it's the prefix for just the transpose ID
	AllNodesKey = widecolumnstore.NewKey(TID, &widecolumnstore.Column{}).Marshal()
)
