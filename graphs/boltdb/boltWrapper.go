package boltdb

import (
	"encoding/json"
	"time"

	"github.com/RossMerr/Caudex.Graph/graphs"
	"github.com/RossMerr/Caudex.Graph/graphs/boltdb/internal"
	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

type BoltWrapper struct {
	bolt *bolt.DB
}

func CreateBoltWrapper(o *graphs.Options) *BoltWrapper {
	database := BoltWrapper{}
	var err error
	var db *bolt.DB
	var b *bolt.Bucket

	logrus.Info("Opening " + o.Name)
	// It will be created if it doesn't exist.
	if db, err = bolt.Open(o.Name+".db", 0600, &bolt.Options{Timeout: 1 * time.Second}); err != nil {
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

	database.bolt = db
	return &database
}

func (db *BoltWrapper) Close() {
	db.Close()
}

func (db *BoltWrapper) Create(c *[]graphs.Vertex) error {
	var err error
	var buf []byte
	return db.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		for _, vertex := range *c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID), buf)
			} else {
				return err
			}
		}
		return nil
	})
}

func (db *BoltWrapper) Delete(c *[]graphs.Vertex) error {
	return db.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		for _, vertex := range *c {
			b.Delete([]byte(vertex.ID))
		}
		return nil
	})
}

func (db *BoltWrapper) Find(ID string) (*graphs.Vertex, error) {
	var err error
	var buf []byte
	var v graphs.Vertex
	return &v, db.bolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		buf = b.Get([]byte(ID))

		if err = json.Unmarshal(buf, v); err == nil {
			return nil
		}

		return err
	})
}

func (db *BoltWrapper) Update(c *[]graphs.Vertex) error {
	var err error
	var buf []byte
	return db.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		for _, vertex := range *c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID), buf)
			} else {
				return err
			}
		}
		return nil
	})
}
