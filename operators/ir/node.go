package ir

import "github.com/google/uuid"

type Node struct {
	Id uuid.UUID
	Variable string
	Label string
	Properties map[string]interface{}
}

type MapLiteral struct {

}

type KeyValue struct {
	Key string
	Value interface{}
}


type Create struct {

}
