package boltdb

import (
	"os"

	"github.com/RossMerr/Caudex.Graph/graphs"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/Sirupsen/logrus"
)

// Graph the underlying graph
type BoltGraph struct {
	db      *BoltWrapper
	Options *graphs.Options
}

// Open graph
func Open(o *graphs.Options) graphs.Graph {
	st := BoltGraph{Options: o, db: CreateBoltWrapper(o)}
	c := make(chan os.Signal, 1)
	st.backgroundTask(c)
	return &st
}

// Close graph
func (g *BoltGraph) Close() {
	g.db.Close()
}

// Query over the graph using the cypher query language returns JSON
func (g *BoltGraph) Query(cypher string) string {
	query.Parse(cypher)
	return "test"
}

// Update
func (g *BoltGraph) Update(fn func(*graphs.GraphOperation) error) error {
	op := &graphs.GraphOperation{DB: g.db}
	return fn(op)
}

func (g *BoltGraph) backgroundTask(c chan os.Signal) {

	go func() {
	Loop:
		for {
			select {
			case <-c:
				logrus.Debug("Received an interrupt, stopping services...")
				break Loop
			}
		}
		close(c)
	}()
}
