package lexer

import "fmt"

var _ error = (*PositionError)(nil)

// PositionError is a parse error that contains a position.
type PositionError struct {
	Pos Position
	Err error
}

func (e *PositionError) Error() string {
	return fmt.Sprintf("At %s: %s", e.Pos, e.Err)
}
