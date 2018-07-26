package ast

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
)

var _ TerminalExpr = (*PropertyStmt)(nil)
var _ InterpretExpr = (*PropertyStmt)(nil)

// PropertyStmt represents a node property.
type PropertyStmt struct {
	Variable string
	Value    string
}

func (PropertyStmt) interpretNode() {}
func (PropertyStmt) exprNode()      {}

func (p *PropertyStmt) GetValue() interface{} {
	return p.Value
}

func (p *PropertyStmt) SetValue(x interface{}) {
	if s, ok := x.(string); ok {
		p.Value = s
	}
}

func (p *PropertyStmt) Interpret(variable string, prop *keyvalue.KeyValue) interface{} {
	if p.Variable == variable {
		split := bytes.Split(prop.Key, keyvalue.US)
		if len(split) > 1 {
			if bytes.Equal(split[1], keyvalue.Properties) {
				if p.Value == string(split[2]) {
					return prop.Value.Unmarshal()
				}
			}
		}
	}
	return nil
}
