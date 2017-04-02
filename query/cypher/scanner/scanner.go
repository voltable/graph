package scanner

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/RossMerr/Caudex.Graph/query/cypher/token"
)

var eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok token.Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if tok, lit := s.scanCharacter(ch); tok != token.ILLEGAL {
		return tok, lit
	}

	s.unread()
	return s.scanIdent()
}

func (s *Scanner) scanCharacter(ch rune) (tok token.Token, lit string) {
	// Otherwise read the individual character.

	switch ch {
	case eof:
		return token.EOF, ""
	case '(':
		return token.LPAREN, string(ch)
	case ')':
		return token.RPAREN, string(ch)
	case ',':
		return token.COMMA, string(ch)
	case ':':
		return token.COLON, string(ch)
	case '.':
		return token.DOT, string(ch)
	case '|':
		return token.PIPE, string(ch)
	case '[':
		return token.LSQUARE, string(ch)
	case ']':
		return token.RSQUARE, string(ch)
	case '+':
		return token.ADD, string(ch)
	case '-':
		return token.SUB, string(ch)
	case '*':
		return token.MUL, string(ch)
	case '/':
		return token.DIV, string(ch)
	case '%':
		return token.MOD, string(ch)
	case '^':
		return token.POW, string(ch)
	case '=':
		next := s.read()
		if next == '~' {
			return token.EQREGEX, "=~"
		}
		s.unread()
		return token.EQ, string(ch)
	case '{':
		return token.LCURLY, string(ch)
	case '}':
		return token.RCURLY, string(ch)
	case '"':
		return token.QUOTATION, string(ch)
	case '\'':
		return token.SINGLEQUOTATION, string(ch)
	case '<':
		next := s.read()
		if next == '=' {
			return token.LTE, "<="
		}
		if next == '>' {
			return token.NEQ, "<>"
		}
		s.unread()
		return token.LT, string(ch)
	case '>':
		next := s.read()
		if next == '=' {
			return token.GTE, ">="
		}
		s.unread()
		return token.GT, string(ch)
	}
	return token.ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return token.WS, buf.String()
}

func (s *Scanner) isCharacter(ch rune) bool {
	tok, _ := s.scanCharacter(ch)
	return tok != token.ILLEGAL
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) && ch != '_' || s.isCharacter(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	lit = buf.String()

	if tok = token.Clause(lit); tok != token.IDENT {
		return tok, buf.String()
	}

	if tok = token.SubClause(lit); tok != token.IDENT {
		return tok, buf.String()
	}

	if tok = token.Boolean(lit); tok != token.IDENT {
		return tok, buf.String()
	}

	if tok = token.Comparison(lit); tok != token.IDENT {
		return tok, buf.String()
	}

	switch strings.ToUpper(buf.String()) {
	case "IS":
		return token.IS, buf.String()
	case "NULL":
		return token.NULL, buf.String()
	}

	// Otherwise return as a regular identifier.
	return token.IDENT, buf.String()
}
