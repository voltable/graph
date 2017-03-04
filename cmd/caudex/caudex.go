package main

import (
	"github.com/RossMerr/Caudex.Graph/storageEngines"
)

func main() {
	if engine, err := storageEngines.NewStorageEngine("bolt", &storageEngines.Options{"test"}); err == nil {
		engine.Close()
	}
}
