package storageEngines

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines/boltdb"
	"github.com/RossMerr/Caudex.Graph/storageEngines/memorydb"
)

// StorageEngineType the backend persistence storage engine
type StorageEngineType int

const (
	// Bolt use Bolt as the storage engine (Default)
	Bolt StorageEngineType = iota
	// Memory use in memory for the storage engine
	Memory StorageEngineType = iota
)

func BuildGraph(e StorageEngineType) graphs.Graph {
	if e == Memory {
		return memorydb.BuildGraph()
	}
	return boltdb.BuildGraph()
}
