package ir

import (
	"bytes"
	"strings"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

// ToPredicatePath creates a PredicatePath out of the Patn
func ToPredicatePath(patn *ast.Patn) *query.PredicatePath {
	label := strings.ToLower(patn.Label)
	pvp := query.PredicatePath{Predicate: func(kv *keyvalue.KeyValue) (string, query.Traverse) {
		split := bytes.Split(kv.Key, US)

		if bytes.Equal(split[1], vertex) {
			value, ok := kv.Value.Unmarshal().(string)
			if ok && label != value {
				return patn.Variable, query.Failed
			}

			return patn.Variable, query.Matched
		}

		if bytes.Equal(split[1], properties) {
			key := split[2]
			property := string(key)
			if value, ok := patn.Properties[property]; ok {
				if value != kv.Value.Unmarshal() {
					return patn.Variable, query.Failed
				}
			}

			return patn.Variable, query.Matched
		}

		return patn.Variable, query.Failed

	}, Variable: patn.Variable}

	return &pvp
}
