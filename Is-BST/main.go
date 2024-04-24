package main

func main() {}

func isBST(nodes []int64) bool {
	return true
}

type BST struct {
	Nodes []int
}

func (bst *BST) parent(nodeIdx int) int {
	return int(float64(nodeIdx-1) / float64(2))
}

func (bst *BST) left(nodeIdx int) int {
	return bst.Nodes[(2*nodeIdx)+1]
}

func (bst *BST) right(nodeIdx int) int {
	return bst.Nodes[(2*nodeIdx)+2]
}

func (bst *BST) isGreaterThanAllItsParents(nodeIdx int) bool {
	for nodeIdx > 0 {
		if bst.Nodes[nodeIdx] > bst.Nodes[int(float64(nodeIdx-1)/float64(2))] {
			nodeIdx = int(float64(nodeIdx-1) / float64(2))
		} else {
			return false
		}
	}
	return true
}
func (bst *BST) isSmallerThanAllItsParents(nodeIdx int) bool {
	for nodeIdx > 0 {
		if bst.Nodes[nodeIdx] < bst.Nodes[int(float64(nodeIdx-1)/float64(2))] {
			nodeIdx = int(float64(nodeIdx-1) / float64(2))
		} else {
			return false
		}
	}
	return true
}
