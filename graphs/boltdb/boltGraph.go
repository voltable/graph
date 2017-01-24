package boltdb

import (
	"os"
	"time"

	"bitbucket.org/rossmerr/caudex/graphs"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb/internal"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb/operations"
	state "bitbucket.org/rossmerr/caudex/graphs/internal"
	"bitbucket.org/rossmerr/caudex/query"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

// Graph the underlying graph
type Graph struct {
	db      *bolt.DB
	Options *graphs.Options
	opend   bool
	ready   bool

	vertices     chan graphs.VertexEnvelop
	resultVertex chan *graphs.Vertex
	findVetex    chan string
}

// Open graph
func Open(o *graphs.Options) graphs.Graph {
	st := Graph{opend: true, Options: o}
	var err error
	var db *bolt.DB
	var b *bolt.Bucket

	logrus.Info("Opening " + st.Options.Name)
	// It will be created if it doesn't exist.
	if db, err = bolt.Open(st.Options.Name+".db", 0600, &bolt.Options{Timeout: 1 * time.Second}); err != nil {
		logrus.Panic(err)
	}

	// create the bucket to hold the Adjacency list.
	db.Update(func(tx *bolt.Tx) error {
		if b, err = tx.CreateBucketIfNotExists([]byte(internal.BucketGraph)); err != nil {
			logrus.Panic(err)
		}

		if b, err = tx.CreateBucketIfNotExists([]byte(internal.BucketIndex)); err != nil {
			logrus.Panic(err)
		}

		return nil
	})

	st.ready = true
	st.db = db
	c := make(chan os.Signal, 1)
	st.backgroundTask(c)
	return &st
}

// Close graph
func (g *Graph) Close() {
	close(g.vertices)
	close(g.resultVertex)
	close(g.findVetex)
	defer g.db.Close()
}

// Query over the graph using the cypher query language returns JSON
func (g *Graph) Query(cypher string) string {
	query.Parse(cypher)
	return "test"
}

// Update
func (g *Graph) Update(fn func(*graphs.GraphOperation) error) error {
	op := graphs.BuildGraphOperation(g.vertices)
	return fn(op)
}

func (g *Graph) backgroundTask(c chan os.Signal) {

	go func() {
	Loop:
		for {
			select {
			case <-c:
				logrus.Debug("Received an interrupt, stopping services...")
				break Loop
			case <-g.vertices:
				for ve := range g.vertices {
					switch ve.State {
					case state.Add:
						if err := operations.Add(g.db, ve.Vertices); err != nil {
							logrus.Debug(err)
						}
					case state.Change:
						if err := operations.Change(g.db, ve.Vertices); err != nil {
							logrus.Debug(err)
						}
					case state.Delete:
						if err := operations.Delete(g.db, ve.Vertices); err != nil {
							logrus.Debug(err)
						}
					}
				}
			}
		}
		close(c)
	}()
}
