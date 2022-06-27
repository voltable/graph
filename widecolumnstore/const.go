package widecolumnstore

var (
	// US unit separator can be used as delimiters to mark fields of data structures.
	// If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = []byte(string('\u241F'))
)
