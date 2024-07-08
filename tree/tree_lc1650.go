package tree

func lowestCommonAncestorIII(p, q *NodeWithParent) *NodeWithParent {

	if p.Parent == nil {
		return p
	}

	if q.Parent == nil {
		return q
	}

	pAncestors, qAncestors := []*NodeWithParent{}, []*NodeWithParent{}

	for ; p != nil; p = p.Parent {
		pAncestors = append(pAncestors, p)
	}

	for ; q != nil; q = q.Parent {
		qAncestors = append(qAncestors, q)
	}

	i, j := len(pAncestors)-1, len(qAncestors)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if pAncestors[i] != qAncestors[j] {
			break
		}
	}

	return pAncestors[i+1]
}

func lowestCommonAncestorIIIV2(p, q *NodeWithParent) *NodeWithParent {

	pp, pq := p, q

	for pp.Val != pq.Val {

		pp = pp.Parent
		pq = pq.Parent

		if pp == nil {
			pp = q
		}

		if pq == nil {
			pq = p
		}
	}

	return pp
}

func lowestCommonAncestorIIIV2PV1(p, q *NodeWithParent) *NodeWithParent {
	pp, pq := p, q

	for pp.Val != pq.Val {

		pp = pp.Parent
		pq = pq.Parent
		if pp == nil {
			pp = q
		}

		if pq == nil {
			pq = p
		}
	}
	return pp
}
