package query

import (
	"fmt"
	"strings"
)

// Syntax
const (
	MATCH    = "MATCH"
	OPTIONAL = "OPTIONAL"
	WHERE    = "WHERE"
	WITH     = "WITH"
	RETURN   = "RETURN"
	UNION    = "UNION"
)

const (
	CREATE  = "CREATE"
	MERGE   = "MERGE"
	DELETE  = "DELETE"
	REMOVE  = "REMOVE"
	SET     = "SET"
	FOREACH = "FOREACH"
)

func Parse(query string) {
	fmt.Println(query)
	fragments := strings.Split(query, " ")
	length := len(fragments)
	fmt.Println(length)
	for index := 0; index < length; index++ {
		fmt.Println("test")
		fmt.Println(fragments[index])
	}

	//	fmt.Printf("%q\n", fragments)
}

func keyword(keyword string) {
	switch keyword {
	case MATCH:

	}
}

func match() {

}
