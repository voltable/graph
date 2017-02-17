package boltdb

type Bucket string

const (
	BucketGraph Bucket = "graph"
	BucketLabel Bucket = "label"
	BucketIndex Bucket = "index"
)
