package cypher

import (
	"fmt"
	"io"
	"strconv"

	"github.com/RossMerr/Caudex.Graph/query/cypher/statements"
)

const emptyString = ""
const MaxUint uint = ^uint(0)
const MinUint uint = 1

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func (p *Parser) Label() (string, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == COLON {
		tok, lit = p.scanIgnoreWhitespace()
		return lit, true
	}
	p.unscan()
	return emptyString, false
}

func (p *Parser) Properties() (map[string]interface{}, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LCURLY {

		if properties, err := p.KeyValue(); err == nil {
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != RCURLY {
				return nil, fmt.Errorf("found %q, expected %q", lit, RCURLY)
			}
			return properties, nil
		}

	}
	p.unscan()
	return nil, nil
}

// KeyValue Loop over all our comma-delimited fields.
func (p *Parser) KeyValue() (map[string]interface{}, error) {
	var properties = make(map[string]interface{})
	for {
		tok, lit := p.scanIgnoreWhitespace()
		var prop = lit

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != COLON {
			return nil, fmt.Errorf("found %q, expected %q", lit, COLON)
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok == QUOTATION {
			// We found a double quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != QUOTATION {
				return nil, fmt.Errorf("found %q, expected %q", lit, QUOTATION)
			}
		} else if tok != IDENT && tok == SINGLEQUOTATION {
			// We found a single quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != SINGLEQUOTATION {
				return nil, fmt.Errorf("found %q, expected %q", lit, SINGLEQUOTATION)
			}
		} else {
			if i, err := strconv.Atoi(lit); err == nil {
				properties[prop] = i
			} else if f, err := strconv.ParseFloat(lit, 64); err == nil {
				properties[prop] = f
			} else if b, err := strconv.ParseBool(lit); err == nil {
				properties[prop] = b
			} else {
				properties[prop] = lit
			}
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != COMMA {
			p.unscan()
			break
		}

	}

	return properties, nil
}

func (p *Parser) Node() (*statements.VertexStatement, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LPAREN {
		stmt := &statements.VertexStatement{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == RPAREN {
			return stmt, nil
		} else if tok == IDENT {
			stmt.Variable = lit
		} else {
			p.unscan()
		}

		if label, ok := p.Label(); ok {
			stmt.Label = label
		}

		if properties, err := p.Properties(); err == nil && properties != nil {
			stmt.Properties = properties
		} else if err != nil {
			return nil, err
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RPAREN {
			return nil, fmt.Errorf("found %q, expected %q", lit, RPAREN)
		}

		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Length() (uint, uint, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == MUL {
		min := MinUint
		max := MaxUint

		tok, lit = p.scanIgnoreWhitespace()
		// We have a number
		if tok == IDENT {
			if u64, err := strconv.ParseUint(lit, 10, 32); err == nil {
				min = uint(u64)
				max = uint(u64)
			} else {
				p.unscan()
			}

			tok, lit = p.scanIgnoreWhitespace()
			if tok == DOT {
				tok, lit = p.scanIgnoreWhitespace()
				if tok == DOT {
					max = MaxUint
					tok, lit = p.scanIgnoreWhitespace()
					if u64, err := strconv.ParseUint(lit, 10, 32); err == nil {
						max = uint(u64)
						if min > max {
							return 0, 0, fmt.Errorf("minimum length %d can't exceed maximum length %d for a relationships", min, max)
						}
					} else {
						p.unscan()
					}
				} else {
					return 0, 0, fmt.Errorf("found %q, expected %q", lit, DOT)
				}
			} else {
				p.unscan()
			}
			// Else we have a range
		} else if tok == DOT {
			tok, lit = p.scanIgnoreWhitespace()
			if tok == DOT {
				min = MinUint
				tok, lit = p.scanIgnoreWhitespace()
				if tok == IDENT {
					if u64, err := strconv.ParseUint(lit, 10, 32); err == nil {
						max = uint(u64)
					} else {
						return 0, 0, fmt.Errorf("found %q, expected uint", lit)
					}
				} else {
					return 0, 0, fmt.Errorf("found %q, expected uint", lit)
				}
			} else {
				return 0, 0, fmt.Errorf("found %q, expected %q", lit, DOT)
			}
		} else {
			p.unscan()
		}
		return min, max, nil
	}
	p.unscan()

	return 0, 0, nil
}

func (p *Parser) RelationshipBody() (*statements.EdgeBodyStatement, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LSQUARE {
		stmt := &statements.EdgeBodyStatement{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == IDENT {
			stmt.Variable = lit
		} else {
			p.unscan()
		}

		if label, ok := p.Label(); ok {
			stmt.Label = label
		}

		if min, max, err := p.Length(); err == nil && (min != 0 && max != 00) {
			stmt.LengthMinimum = min
			stmt.LengthMaximum = max
		} else if err != nil {
			return nil, err
		} else {
			stmt.LengthMinimum = 1
			stmt.LengthMaximum = 1
		}

		if properties, err := p.Properties(); err == nil && properties != nil {
			stmt.Properties = properties
		} else if err != nil {
			return nil, err
		}

		tok, lit := p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RSQUARE {
			return nil, fmt.Errorf("found %q, expected %q", lit, RSQUARE)
		}
		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Relationship() (*statements.EdgeStatement, error) {
	tok, lit := p.scanIgnoreWhitespace()
	// Look for the start of a relationship < or -
	if tok != IDENT && (tok == LT || tok == SUB) {
		stmt := &statements.EdgeStatement{Relationship: statements.Undirected}

		if tok == LT {
			stmt.Relationship = statements.Outbound

			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship -
			if tok != IDENT && tok != SUB {
				return nil, fmt.Errorf("found %q, expected %q", lit, SUB)
			}
		}

		if body, err := p.RelationshipBody(); err == nil && body != nil {
			stmt.Body = body
		} else if err != nil {
			return nil, err
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != SUB {
			return nil, fmt.Errorf("found %q, expected %q", lit, SUB)
		}

		// Check for inbound relationship
		if tok == SUB {
			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship - or >
			if tok != IDENT && tok == GT {
				stmt.Relationship = statements.Inbound
			} else {
				p.unscan()
			}
		}

		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Comparison() (statements.Comparison, error) {
	tok, lit := p.scanIgnoreWhitespace()

	if tok == EQ {
		return statements.EQ, nil
	} else if tok == LT {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == EQ {
			return statements.LTE, nil
		} else if tok == GT {
			return statements.NEQ, nil
		}
		p.unscan()
		return statements.LT, nil
	} else if tok == GT {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == EQ {
			return statements.GTE, nil
		}
		p.unscan()
		return statements.GT, nil
	}

	return statements.EQ, fmt.Errorf("found %q, expected Comparison", lit)
}

func (p *Parser) Value() (interface{}, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == SINGLEQUOTATION {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == IDENT {
			value := lit
			tok, lit := p.scanIgnoreWhitespace()
			if tok == SINGLEQUOTATION {
				return value, nil
			}

			return emptyString, fmt.Errorf("found %q, expected %q", lit, SINGLEQUOTATION)
		}

		return emptyString, fmt.Errorf("found %q, expected %q", lit, IDENT)
	} else if tok == IDENT {
		fmt.Println(lit)
		if i, err := strconv.Atoi(lit); err == nil {
			return i, nil
		} else if f, err := strconv.ParseFloat(lit, 64); err == nil {
			return f, nil
		} else if b, err := strconv.ParseBool(lit); err == nil {
			return b, nil
		}
		return lit, nil
	}

	p.unscan()
	return emptyString, nil
}

func (p *Parser) Boolean() (statements.BooleanStatement, error) {
	tok, lit := p.scanIgnoreWhitespace()
	fmt.Println(lit)
	if tok == AND {
		state := &statements.AndStatement{}
		if predicate, err := p.Predicate(); err == nil {
			state.Predicate = predicate
		} else {
			return nil, err
		}
		return state, nil
	}

	p.unscan()
	return nil, nil
}
func (p *Parser) Predicate() (*statements.PredicateStatement, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == IDENT {
		state := &statements.PredicateStatement{}
		state.Variable = lit

		tok, lit := p.scanIgnoreWhitespace()
		if tok == DOT {
			tok, lit := p.scanIgnoreWhitespace()
			if tok == IDENT {
				state.Property = lit
				if operator, err := p.Comparison(); err == nil {
					state.Operator = operator
				} else {
					return nil, err
				}

				if value, err := p.Value(); err == nil {
					state.Value = value
				} else {
					return nil, err
				}

			}
		} else {
			return nil, fmt.Errorf("found %q, expected %q", lit, DOT)
		}

		if b, err := p.Boolean(); err == nil && b != nil {
			state.Next = b
		} else if err != nil {
			return nil, err
		}

		return state, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Where() (statements.Statement, error) {
	tok, _ := p.scanIgnoreWhitespace()
	if tok == WHERE {
		state := &statements.WhereStatement{}

		if predicate, err := p.Predicate(); err == nil {
			state.Predicate = predicate
		} else {
			return nil, err
		}

		return state, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Match() (statements.Statement, error) {
	state := &statements.MatchStatement{}

	var lastVertex *statements.VertexStatement
	var lastEdge *statements.EdgeStatement

	// Next we should loop over all the pattern.
	for {

		if node, err := p.Node(); err == nil && node != nil {
			lastVertex = node
			if state.Pattern == nil {
				state.Pattern = lastVertex
			}
			if lastEdge != nil {
				lastEdge.Vertex = node
			}
		} else if err != nil {
			return nil, err
		}

		if relationship, err := p.Relationship(); err == nil && relationship != nil {
			lastEdge = relationship
			lastVertex.Edge = relationship
		} else if err != nil {
			return nil, err
		} else {
			break
		}
	}

	if where, err := p.Where(); err == nil && where != nil {
		state.Next = where
	} else if err != nil {
		return nil, err
	}

	return state, nil
}

func (p *Parser) OptionalMatch() (statements.Statement, error) {
	state := &statements.OptionalMatchStatement{}
	return state, nil
}

func (p *Parser) Clause() (statements.Statement, error) {
	tok, lit := p.scanIgnoreWhitespace()

	if !tok.isClause() {
		return nil, fmt.Errorf("found %q, expected a clause", lit)
	}

	if tok == OPTIONAL {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == MATCH {
			tok = OPTIONAL_MATCH
		} else {
			return nil, fmt.Errorf("found %q, expected MATCH", lit)
		}
	} else if tok == DETACH {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == DELETE {
			tok = DETACH_DELETE
		} else {
			return nil, fmt.Errorf("found %q, expected DELETE", lit)
		}
	}

	switch tok {
	case MATCH:
		return p.Match()
	case OPTIONAL_MATCH:
		return p.OptionalMatch()
	}

	return nil, fmt.Errorf("No matching statement found %q", lit)
}

func (p *Parser) SubClause() (Token, bool) {
	tok, _ := p.scanIgnoreWhitespace()

	if tok.isSubClause() {
		if tok == ON {
			tok, lit := p.scanIgnoreWhitespace()
			if tok == CREATE {
				return ON_CREATE, true
			} else if tok == MATCH {
				return ON_MATCH, true
			} else {
				panic(fmt.Sprintf("found %q, expected CREATE", lit))
			}
		}

		if tok == ORDER {
			tok, lit := p.scanIgnoreWhitespace()
			if tok == BY {
				return ORDER_BY, true
			} else {
				panic(fmt.Sprintf("found %q, expected BY", lit))
			}
		}

		return tok, true
	}

	p.unscan()
	return IDENT, false
}

// Parse parses a cypher Clauses statement.
func (p *Parser) Parse() (statements.Statement, error) {
	return p.Clause()
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
