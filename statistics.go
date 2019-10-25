package graph

type Statistics struct {
	Rows   int
	DbHits DbHits
}

func NewStatistics() Statistics {
	return Statistics{
		DbHits: DbHits{},
	}
}

type DbHits struct {
	CreateNodes         int `json:"+nodes"`
	CreateRelationships int `json:"+relationships"`
	CreateLabels        int `json:"+labels"`
	CreateTypes         int `json:"+types"`
	CreateProperties    int `json:"+properties"`
}

