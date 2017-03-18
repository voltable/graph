package cypher

import (
	"fmt"
	"io"
	"strconv"
)

const emptyString = ""

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
			panic(fmt.Sprintf("found %q, expected field", lit))
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
			panic(fmt.Sprintf("found %q, expected field", lit))
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok == QUOTATION {
			// We found a double quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != QUOTATION {
				panic(fmt.Sprintf("found %q, expected field", lit))
			}
		} else if tok != IDENT && tok == SINGLEQUOTATION {
			// We found a single quoted string
			tok, lit = p.scanIgnoreWhitespace()
			properties[prop] = lit
			tok, lit = p.scanIgnoreWhitespace()
			if tok != IDENT && tok != SINGLEQUOTATION {
				panic(fmt.Sprintf("found %q, expected field", lit))
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

func (p *Parser) Node() (*MatchVertexStatement, bool) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != IDENT && tok == LPAREN {
		stmt := &MatchVertexStatement{}
		stmt.Properties = make(map[string]interface{})

		tok, lit = p.scanIgnoreWhitespace()
		stmt.Variable = lit

		if label, ok := p.Label(); ok {
			stmt.Label = label
		}

		if properties, ok := p.Properties(); ok {
			stmt.Properties = properties
		}

		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT && tok != RPAREN {
			panic(fmt.Sprintf("found %q, expected field", lit))
		}

		return stmt, true
	}

	p.unscan()
	return nil, false

}

func (p *Parser) Match() (*MatchVertexStatement, error) {

	if tok, lit := p.scanIgnoreWhitespace(); tok != MATCH {
		return nil, fmt.Errorf("found %q, expected MATCH", lit)
	}

	// Next we should loop over all the pattern.
	for {

		if node, ok := p.Node(); ok {
			return node, nil
		}

		break
	}

	return nil, nil
}

// Parse parses a cypher MATCH statement.
func (p *Parser) Parse() (*MatchVertexStatement, error) {
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
