package others

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type FrequencyNode struct {
	Node
	Frequency int
	Pre       *FrequencyNode
	Next      *FrequencyNode
}
