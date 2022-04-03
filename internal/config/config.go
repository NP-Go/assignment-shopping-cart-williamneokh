package config

type ItemInformation struct {
	CatType  int
	Quantity int
	Cost     float64
}

var Category []string
var NameMap map[string]ItemInformation

func LoadMap() {
	NameMap = make(map[string]ItemInformation)
}
