package graph

type Statistics struct {
	Rows int
	DbHits DbHits
}

func NewStatistics() Statistics {
	return Statistics{
		DbHits: DbHits{},
	}
}

type DbHits struct {
	CreateNode int  `json:"+nodes"`
}

