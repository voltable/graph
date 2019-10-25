package ir

const EmptyString = ""

var (
	ID            = []byte("i")
	Label         = []byte("l")
	Properties    = []byte("p")
	Vertex        = []byte("v")
	Edge          = []byte("r")
	EdgeType= []byte("t")
	EdgePoperties = []byte("k")

	// T = transpose
	TID             = []byte("ti")
	TLabel          = []byte("tl")
	TProperties     = []byte("tp")
	TVertex         = []byte("tv")
	TEdge           = []byte("tr")
	TEdgeType= []byte("t")
	TEdgeProperties = []byte("tk")

	// This Key can be used to fetch all nodes, it's the prefix for just the transpose ID
	//AllNodesKey = widecolumnstore.NewKey(TID, &widecolumnstore.Column{}).Marshal()
)
