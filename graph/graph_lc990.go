package graph

func equationsPossible(equations []string) bool {

	uf := NewUF(26)

	for _, equ := range equations {
		if equ[1] == '=' {
			a := int(equ[0] - 'a')
			b := int(equ[3] - 'a')
			uf.Union(a, b)
		}
	}

	for _, equ := range equations {
		if equ[1] == '!' {
			a := int(equ[0] - 'a')
			b := int(equ[3] - 'a')
			if uf.Connected(a, b) {
				return false
			}
		}
	}

	return true
}
