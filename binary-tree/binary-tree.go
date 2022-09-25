package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	//end of the tree
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true

}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(30)
	tree.Insert(300)
	fmt.Println(tree.Search(30))
	fmt.Println(tree.Search(320))
}
