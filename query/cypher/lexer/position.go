package lexer

import "fmt"

// Position describes an  source position
type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (character count)
}

// IsValid returns true if the position is valid.
func (p *Position) IsValid() bool { return p.Line > 0 }

// String returns a string in one of several forms:
//
//	line:column         valid position without file name
func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

var _ error = (*PositionError)(nil)

// PositionError is a parse error that contains a position.
type PositionError struct {
	Pos Position
	Err error
}

func (e *PositionError) Error() string {
	return fmt.Sprintf("At %s: %s", e.Pos, e.Err)
}

func (p Position) Errorf(format string, a ...interface{}) *PositionError {
	return &PositionError{Pos: p, Err: fmt.Errorf(format, a)}
}
