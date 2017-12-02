package scanner

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/RossMerr/Caudex.Graph/query/cypher/lexer"
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
func (s *Scanner) Scan() (tok lexer.Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if tok, lit := s.scanCharacter(ch); tok != lexer.ILLEGAL {
		return tok, lit
	}

	s.unread()
	return s.scanIdent()
}

func (s *Scanner) scanCharacter(ch rune) (tok lexer.Token, lit string) {
	// Otherwise read the individual character.

	switch ch {
	case eof:
		return lexer.EOF, ""
	case '`':
		return lexer.GRAVE, string(ch)
	case '(':
		return lexer.LPAREN, string(ch)
	case ')':
		return lexer.RPAREN, string(ch)
	case ',':
		return lexer.COMMA, string(ch)
	case ':':
		return lexer.COLON, string(ch)
	case '.':
		return lexer.DOT, string(ch)
	case '|':
		return lexer.PIPE, string(ch)
	case '[':
		return lexer.LSQUARE, string(ch)
	case ']':
		return lexer.RSQUARE, string(ch)
	case '+':
		return lexer.ADD, string(ch)
	case '-':
		return lexer.SUB, string(ch)
	case '*':
		return lexer.MUL, string(ch)
	case '/':
		return lexer.DIV, string(ch)
	case '%':
		return lexer.MOD, string(ch)
	case '^':
		return lexer.POW, string(ch)
	case '=':
		next := s.read()
		if next == '~' {
			return lexer.EQREGEX, "=~"
		}
		s.unread()
		return lexer.EQ, string(ch)
	case '{':
		return lexer.LCURLY, string(ch)
	case '}':
		return lexer.RCURLY, string(ch)
	case '"':
		return lexer.QUOTATION, string(ch)
	case '\'':
		return lexer.SINGLEQUOTATION, string(ch)
	case '<':
		next := s.read()
		if next == '=' {
			return lexer.LTE, "<="
		}
		if next == '>' {
			return lexer.NEQ, "<>"
		}
		s.unread()
		return lexer.LT, string(ch)
	case '>':
		next := s.read()
		if next == '=' {
			return lexer.GTE, ">="
		}
		s.unread()
		return lexer.GT, string(ch)
	}
	return lexer.ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok lexer.Token, lit string) {
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

	return lexer.WS, buf.String()
}

func (s *Scanner) isCharacter(ch rune) bool {
	tok, _ := s.scanCharacter(ch)
	return tok != lexer.ILLEGAL
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok lexer.Token, lit string) {
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

	if tok = lexer.Clause(lit); tok != lexer.IDENT {
		return tok, buf.String()
	}

	if tok = lexer.SubClause(lit); tok != lexer.IDENT {
		return tok, buf.String()
	}

	if tok = lexer.Boolean(lit); tok != lexer.IDENT {
		return tok, buf.String()
	}

	if tok = lexer.Comparison(lit); tok != lexer.IDENT {
		return tok, buf.String()
	}

	switch strings.ToUpper(buf.String()) {
	case "IS":
		return lexer.IS, buf.String()
	case "NULL":
		return lexer.NULL, buf.String()
	}

	// Otherwise return as a regular identifier.
	return lexer.IDENT, buf.String()
}
