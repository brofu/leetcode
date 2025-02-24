package tree

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {

	tree := make(map[string]string)

	for _, list := range regions {
		parent := list[0]
		for _, region := range list[1:] {
			tree[region] = parent
		}
	}

	region1p, region2p := region1, region2

	for region1p != region2p {
		if val, ok := tree[region1p]; ok {
			region1p = val
		} else {
			region1p = region2
		}

		if val, ok := tree[region2p]; ok {
			region2p = val
		} else {
			region2p = region1
		}
	}

	return region1p
}

func findSmallestRegionPV2(regions [][]string, region1 string, region2 string) string {

	parentMap := make(map[string]string)

	for _, list := range regions {
		for idx := 1; idx < len(list); idx++ {
			parentMap[list[idx]] = list[0]
		}
	}
	a, b := region1, region2

	for a != b {
		if a == "" {
			a = region2
		} else {
			a, _ = parentMap[a]
		}

		if b == "" {
			b = region1
		} else {
			b, _ = parentMap[b]
		}
	}
	return a
}
