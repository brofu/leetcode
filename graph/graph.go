package graph

type Edge struct {
	To     int
	Weight int
}

type Graph interface {

	// add
	AddEdge(from, to, weight int)

	// delete
	DeleteEdge(from, to int)

	// query
	HasEdge(from, to int) bool
	Weight(from, to int) int
	Neighbors(from int) []Edge
}

type DWG struct {
	data [][]Edge
}

func NewDWG(n int) *DWG {

	return &DWG{
		data: make([][]Edge, n+1),
	}
}

func (dwg *DWG) AddEdge(from, to, weight int) {
	list := dwg.data[from]
	if len(list) == 0 {
		dwg.data[from] = []Edge{{To: to, Weight: weight}}
		return
	}
	for idx, edge := range list {
		if edge.To == to {
			edge.Weight = weight
			list[idx] = edge
			dwg.data[from] = list
			return
		}
	}
	list = append(list, Edge{To: to, Weight: weight})
	dwg.data[from] = list
}

func (dwg *DWG) DeleteEdge(from, to int) {
	list := dwg.data[from]
	if len(list) == 0 {
		return
	}

	index := -1
	hit := false
	for idx, edge := range list {
		if edge.To == to {
			hit = true
			index = idx
			break
		}
	}
	if hit {
		list = append(list[:index], list[index+1:]...)
		dwg.data[from] = list
	}
}

func (dwg *DWG) HasEdge(from, to int) bool {
	list := dwg.data[from]
	for _, edge := range list {
		if edge.To == to {
			return true
		}
	}
	return false
}

func (dwg *DWG) Weight(from, to int) int {
	list := dwg.data[from]
	for _, edge := range list {
		if edge.To == to {
			return edge.Weight
		}
	}
	return -1
}

func (dwg *DWG) Neighbors(from int) []Edge {
	return dwg.data[from]
}
