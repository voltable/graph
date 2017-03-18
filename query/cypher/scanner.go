package cypher

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

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
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if tok, lit := s.scanCharacter(ch); tok != ILLEGAL {
		return tok, lit
	}

	s.unread()
	return s.scanIdent()
}

func (s *Scanner) scanCharacter(ch rune) (tok Token, lit string) {
	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '(':
		return LPAREN, string(ch)
	case ')':
		return RPAREN, string(ch)
	case ',':
		return COMMA, string(ch)
	case ':':
		return COLON, string(ch)
	case '.':
		return DOT, string(ch)
	case '|':
		return PIPE, string(ch)
	case '[':
		return LSQUARE, string(ch)
	case ']':
		return RSQUARE, string(ch)
	case '+':
		return ADD, string(ch)
	case '-':
		return SUB, string(ch)
	case '*':
		return MUL, string(ch)
	case '/':
		return DIV, string(ch)
	case '%':
		return MOD, string(ch)
	case '^':
		return POW, string(ch)
	case '=':
		return EQ, string(ch)
	case '<':
		return LT, string(ch)
	case '>':
		return GT, string(ch)
	case '{':
		return LCURLY, string(ch)
	case '}':
		return RCURLY, string(ch)
	case '"':
		return QUOTATION, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
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

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) && ch != '_' {

			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	lit = buf.String()

	if tok = Keyword(lit); tok != IDENT {
		return tok, buf.String()
	}

	if tok = Boolean(lit); tok != IDENT {
		return tok, buf.String()
	}

	if tok = Comparison(lit); tok != IDENT {
		return tok, buf.String()
	}

	switch strings.ToUpper(buf.String()) {

	case "<>":
		return NEQ, buf.String()
	case "<=":
		return LTE, buf.String()
	case ">=":
		return GTE, buf.String()
	case "IS":
		return IS, buf.String()
	case "NULL":
		return NULL, buf.String()
	case "=~":
		return EQREGEX, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}
