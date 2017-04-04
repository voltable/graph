package parser

import (
	"fmt"
	"io"
	"strconv"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/scanner"
	"github.com/RossMerr/Caudex.Graph/query/cypher/token"
)

const emptyString = ""
const MaxUint uint = ^uint(0)
const MinUint uint = 1

// Parser represents a parser.
type Parser struct {
	s   *scanner.Scanner
	buf struct {
		tok token.Token // last read token
		lit string      // last read literal
		n   int         // buffer size (max=1)
	}
}

func (p *Parser) Label() (string, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token.IDENT && tok == token.COLON {
		tok, lit = p.scanIgnoreWhitespace()
		return lit, true
	}
	p.unscan()
	return emptyString, false
}

func (p *Parser) Properties() (map[string]interface{}, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token.IDENT && tok == token.LCURLY {

		if properties, err := p.KeyValue(); err == nil {
			tok, lit = p.scanIgnoreWhitespace()
			if tok != token.IDENT && tok != token.RCURLY {
				return nil, fmt.Errorf("found %q, expected %q", lit, token.RCURLY)
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
		if tok != token.IDENT && tok != token.COLON {
			return nil, fmt.Errorf("found %q, expected %q", lit, token.COLON)
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != token.IDENT && tok == token.QUOTATION {
			// We found a double quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != token.IDENT && tok != token.QUOTATION {
				return nil, fmt.Errorf("found %q, expected %q", lit, token.QUOTATION)
			}
		} else if tok != token.IDENT && tok == token.SINGLEQUOTATION {
			// We found a single quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != token.IDENT && tok != token.SINGLEQUOTATION {
				return nil, fmt.Errorf("found %q, expected %q", lit, token.SINGLEQUOTATION)
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
		if tok != token.COMMA {
			p.unscan()
			break
		}

	}

	return properties, nil
}

func (p *Parser) Node() (*ast.VertexPatn, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token.IDENT && tok == token.LPAREN {
		stmt := &ast.VertexPatn{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == token.RPAREN {
			return stmt, nil
		} else if tok == token.IDENT {
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
		if tok != token.IDENT && tok != token.RPAREN {
			return nil, fmt.Errorf("found %q, expected %q", lit, token.RPAREN)
		}

		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Length() (uint, uint, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token.IDENT && tok == token.MUL {
		min := MinUint
		max := MaxUint

		tok, lit = p.scanIgnoreWhitespace()
		// We have a number
		if tok == token.IDENT {
			if u64, err := strconv.ParseUint(lit, 10, 32); err == nil {
				min = uint(u64)
				max = uint(u64)
			} else {
				p.unscan()
			}

			tok, lit = p.scanIgnoreWhitespace()
			if tok == token.DOT {
				tok, lit = p.scanIgnoreWhitespace()
				if tok == token.DOT {
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
					return 0, 0, fmt.Errorf("found %q, expected %q", lit, token.DOT)
				}
			} else {
				p.unscan()
			}
			// Else we have a range
		} else if tok == token.DOT {
			tok, lit = p.scanIgnoreWhitespace()
			if tok == token.DOT {
				min = MinUint
				tok, lit = p.scanIgnoreWhitespace()
				if tok == token.IDENT {
					if u64, err := strconv.ParseUint(lit, 10, 32); err == nil {
						max = uint(u64)
					} else {
						return 0, 0, fmt.Errorf("found %q, expected uint", lit)
					}
				} else {
					return 0, 0, fmt.Errorf("found %q, expected uint", lit)
				}
			} else {
				return 0, 0, fmt.Errorf("found %q, expected %q", lit, token.DOT)
			}
		} else {
			p.unscan()
		}
		return min, max, nil
	}
	p.unscan()

	return 0, 0, nil
}

func (p *Parser) RelationshipBody() (*ast.EdgeBodyStmt, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != token.IDENT && tok == token.LSQUARE {
		stmt := &ast.EdgeBodyStmt{}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == token.IDENT {
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
		if tok != token.IDENT && tok != token.RSQUARE {
			return nil, fmt.Errorf("found %q, expected %q", lit, token.RSQUARE)
		}
		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Relationship() (*ast.EdgePatn, error) {
	tok, lit := p.scanIgnoreWhitespace()
	// Look for the start of a relationship < or -
	if tok != token.IDENT && (tok == token.LT || tok == token.SUB) {
		stmt := &ast.EdgePatn{Relationship: ast.Undirected}

		if tok == token.LT {
			stmt.Relationship = ast.Outbound

			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship -
			if tok != token.IDENT && tok != token.SUB {
				return nil, fmt.Errorf("found %q, expected %q", lit, token.SUB)
			}
		}

		if body, err := p.RelationshipBody(); err == nil && body != nil {
			stmt.Body = body
		} else if err != nil {
			return nil, err
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != token.IDENT && tok != token.SUB {
			return nil, fmt.Errorf("found %q, expected %q", lit, token.SUB)
		}

		// Check for inbound relationship
		if tok == token.SUB {
			tok, lit = p.scanIgnoreWhitespace()
			// Look for the end of the relationship - or >
			if tok != token.IDENT && tok == token.GT {
				stmt.Relationship = ast.Inbound
			} else {
				p.unscan()
			}
		}

		return stmt, nil
	}

	p.unscan()
	return nil, nil
}

func (p *Parser) Value() (interface{}, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == token.SINGLEQUOTATION {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == token.IDENT {
			value := lit
			tok, lit := p.scanIgnoreWhitespace()
			if tok == token.SINGLEQUOTATION {
				return value, nil
			}

			return emptyString, fmt.Errorf("found %q, expected %q", lit, token.SINGLEQUOTATION)
		}

		return emptyString, fmt.Errorf("found %q, expected %q", lit, token.IDENT)
	} else if tok == token.IDENT {
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

func (p *Parser) Property() (*ast.PropertyStmt, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == token.IDENT {

		state := &ast.PropertyStmt{Variable: lit}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != token.DOT {
			return nil, fmt.Errorf("found %q, expected a DOT", lit)
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != token.IDENT {
			return nil, fmt.Errorf("found %q, expected a IDENT", lit)
		}

		state.Value = lit

		return state, nil
	}
	p.unscan()
	return nil, nil
}

func (p *Parser) BinaryExpr() (*ast.ComparisonExpr, error) {
	tok, _ := p.scanIgnoreWhitespace()
	switch tok {
	case token.EQ:
		return &ast.ComparisonExpr{}, nil
	case token.NEQ:
		return &ast.ComparisonExpr{Comparison: ast.NEQ}, nil
	case token.LT:
		return &ast.ComparisonExpr{}, nil
	case token.LTE:
		return &ast.ComparisonExpr{}, nil
	case token.GT:
		return &ast.ComparisonExpr{}, nil
	case token.GTE:
		return &ast.ComparisonExpr{}, nil
	}
	p.unscan()
	return nil, nil
}

func (p *Parser) Predicate() (ast.Expr, error) {

	var root ast.Expr
	var last ast.Expr
	var binary *ast.ComparisonExpr

	tok, _ := p.scanIgnoreWhitespace()
	p.unscan()

	for !tok.IsClause() && tok != token.EOF {

		if property, err := p.Property(); err == nil && property != nil {
			last = property
		} else if err != nil {
			return nil, err
		} else if binaryExpr, err := p.BinaryExpr(); err == nil && binaryExpr != nil {
			if root == nil {
				root = binaryExpr
			}

			binary = binaryExpr

			// if b, ok := binaryExpr.(ast.NotEqualExpr); ok {
			// 	binary = &b
			// }
		} else if err != nil {
			return nil, err
		} else if value, err := p.Value(); err == nil && value != nil {
			last = &ast.Ident{value}
		} else if err != nil {
			return nil, err
		}

		if binary != nil {
			if binary.X == nil {
				binary.X = last
			} else if binary.Y == nil {
				binary.Y = last
			}
			last = nil
		}

		tok, _ = p.scanIgnoreWhitespace()
		p.unscan()
	}

	return root, nil
}

func (p *Parser) Where() (ast.Stmt, error) {
	tok, _ := p.scanIgnoreWhitespace()
	if tok == token.WHERE {
		state := &ast.WhereStmt{}

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

func (p *Parser) Match() (ast.Stmt, error) {
	state := &ast.MatchStmt{}

	var lastVertex *ast.VertexPatn
	var lastEdge *ast.EdgePatn

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

func (p *Parser) OptionalMatch() (ast.Stmt, error) {
	state := &ast.OptionalMatchStmt{}
	return state, nil
}

func (p *Parser) Clause() (ast.Stmt, error) {
	tok, lit := p.scanIgnoreWhitespace()

	if !tok.IsClause() {
		return nil, fmt.Errorf("found %q, expected a clause", lit)
	}

	if tok == token.OPTIONAL {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == token.MATCH {
			tok = token.OPTIONAL_MATCH
		} else {
			return nil, fmt.Errorf("found %q, expected MATCH", lit)
		}
	} else if tok == token.DETACH {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == token.DELETE {
			tok = token.DETACH_DELETE
		} else {
			return nil, fmt.Errorf("found %q, expected DELETE", lit)
		}
	}

	switch tok {
	case token.MATCH:
		return p.Match()
	case token.OPTIONAL_MATCH:
		return p.OptionalMatch()
	}

	return nil, fmt.Errorf("No matching statement found %q", lit)
}

func (p *Parser) SubClause() (token.Token, bool) {
	tok, _ := p.scanIgnoreWhitespace()

	if tok.IsSubClause() {
		if tok == token.ON {
			tok, lit := p.scanIgnoreWhitespace()
			if tok == token.CREATE {
				return token.ON_CREATE, true
			} else if tok == token.MATCH {
				return token.ON_MATCH, true
			} else {
				panic(fmt.Sprintf("found %q, expected CREATE", lit))
			}
		}

		if tok == token.ORDER {
			tok, lit := p.scanIgnoreWhitespace()
			if tok == token.BY {
				return token.ORDER_BY, true
			} else {
				panic(fmt.Sprintf("found %q, expected BY", lit))
			}
		}

		return tok, true
	}

	p.unscan()
	return token.IDENT, false
}

// Parse parses a cypher Clauses statement.
func (p *Parser) Parse() (ast.Stmt, error) {
	return p.Clause()
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: scanner.NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok token.Token, lit string) {
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
func (p *Parser) scanIgnoreWhitespace() (tok token.Token, lit string) {
	tok, lit = p.scan()
	if tok == token.WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
