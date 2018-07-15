package ir

var (
	vertex                 = []byte("v")
	label                  = []byte("l")
	properties             = []byte("p")
	relationship           = []byte("r")
	relationshipproperties = []byte("k")
	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = []byte(string('\u241F'))
)
