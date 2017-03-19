package cypher

import (
	"fmt"
	"io"
	"strconv"
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

func (p *Parser) Properties() (map[string]interface{}, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LCURLY {

		var properties, _ = p.KeyValue()

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RCURLY {
			panic(fmt.Sprintf("found %q, expected %q", lit, RCURLY))
		}
		return properties, true
	}
	p.unscan()
	return nil, false
}

// KeyValue Loop over all our comma-delimited fields.
func (p *Parser) KeyValue() (map[string]interface{}, bool) {
	var properties = make(map[string]interface{})
	for {
		tok, lit := p.scanIgnoreWhitespace()
		var prop = lit

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != COLON {
			panic(fmt.Sprintf("found %q, expected %q", lit, COLON))
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok == QUOTATION {
			// We found a double quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != QUOTATION {
				panic(fmt.Sprintf("found %q, expected %q", lit, QUOTATION))
			}
		} else if tok != IDENT && tok == SINGLEQUOTATION {
			// We found a single quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != SINGLEQUOTATION {
				panic(fmt.Sprintf("found %q, expected %q", lit, SINGLEQUOTATION))
			}
		} else {
			if b, err := strconv.ParseBool(lit); err == nil {
				properties[prop] = b
			} else if i, err := strconv.Atoi(lit); err == nil {
				properties[prop] = i
			} else if f, err := strconv.ParseFloat(lit, 64); err == nil {
				properties[prop] = f
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

	return properties, true
}

func (p *Parser) Node() (*VertexStatement, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LPAREN {
		stmt := &VertexStatement{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == RPAREN {
			return stmt, true
		} else if tok == IDENT {
			stmt.Variable = lit
		} else {
			p.unscan()
		}

		if label, ok := p.Label(); ok {
			stmt.Label = label
		}

		if properties, ok := p.Properties(); ok {
			stmt.Properties = properties
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RPAREN {
			panic(fmt.Sprintf("found %q, expected %q", lit, RPAREN))
		}

		return stmt, true
	}

	p.unscan()
	return nil, false
}

func (p *Parser) Length() (uint, uint, bool) {
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
							panic(fmt.Sprintf("minimum length %d can't exceed maximum length %d for a relationships", min, max))
						}
					} else {
						p.unscan()
					}
				} else {
					panic(fmt.Sprintf("found %q, expected %q", lit, DOT))
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
						panic(fmt.Sprintf("found %q, expected uint", lit))
					}
				} else {
					panic(fmt.Sprintf("found %q, expected uint", lit))
				}
			} else {
				panic(fmt.Sprintf("found %q, expected %q", lit, DOT))
			}
		} else {
			p.unscan()
		}
		return min, max, true
	}
	p.unscan()
	return 0, 0, false
}

func (p *Parser) RelationshipBody() (*EdgeBodyStatement, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LSQUARE {
		stmt := &EdgeBodyStatement{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == IDENT {
			stmt.Variable = lit
		} else {
			p.unscan()
		}

		if label, ok := p.Label(); ok {
			stmt.Label = label
		}

		if min, max, ok := p.Length(); ok {
			stmt.LengthMinimum = min
			stmt.LengthMaximum = max
		} else {
			stmt.LengthMinimum = 1
			stmt.LengthMaximum = 1
		}

		if properties, ok := p.Properties(); ok {
			stmt.Properties = properties
		}

		tok, lit := p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RSQUARE {
			panic(fmt.Sprintf("found %q, expected %q", lit, RSQUARE))
		}
		return stmt, true
	}

	p.unscan()
	return nil, false
}

func (p *Parser) Relationship() (*EdgeStatement, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	// Look for the start of a relationship < or -
	if tok != IDENT && (tok == LT || tok == SUB) {
		stmt := &EdgeStatement{Relationship: Undirected}

		if tok == LT {
			stmt.Relationship = Outbound

			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship -
			if tok != IDENT && tok != SUB {
				panic(fmt.Sprintf("found %q, expected %q", lit, SUB))
			}
		}

		if body, ok := p.RelationshipBody(); ok {
			stmt.Body = body
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != SUB {
			panic(fmt.Sprintf("found %q, expected %q", lit, SUB))
		}

		// Check for inbound relationship
		if tok == SUB {
			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship - or >
			if tok != IDENT && tok == GT {
				stmt.Relationship = Inbound
			} else {
				p.unscan()
			}
		}

		return stmt, true
	}

	p.unscan()
	return nil, false
}

func (p *Parser) Match() (*VertexStatement, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != MATCH {
		return nil, fmt.Errorf("found %q, expected MATCH", lit)
	}

	var first *VertexStatement
	var lastVertex *VertexStatement
	var lastEdge *EdgeStatement

	// Next we should loop over all the pattern.
	for {

		if node, ok := p.Node(); ok {
			lastVertex = node
			if first == nil {
				first = lastVertex
			}
			if lastEdge != nil {
				lastEdge.Vertex = node
			}
			//return node, nil
		}

		if relationship, ok := p.Relationship(); ok {
			lastEdge = relationship
			lastVertex.Edge = relationship
		} else {
			return first, nil
		}
	}

	return nil, nil
}

// Parse parses a cypher MATCH statement.
func (p *Parser) Parse() (*VertexStatement, error) {
	return p.Match()
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
