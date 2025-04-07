package uf

func equationsPossible(equations []string) bool {

	uf := NewPathCompactedUF(26)

	for _, item := range equations {
		if item[1] == '=' {
			uf.Connect(int(item[0]-'a'), int(item[3]-'a'))
		}
	}
	for _, item := range equations {
		if item[1] == '!' {
			if uf.Connected(int(item[0]-'a'), int(item[3]-'a')) {
				return false
			}
		}
	}
	return true
}
