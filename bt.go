/*
	Print all possible path form root to leaf in a BST
				  a
				/   \
			 b     c
			/ \   / \
		 d   e f   g

	Print:
	abd
	abe
	acf
	acg
*/

package main

import "fmt"

type Node struct {
	val   string
	left  *Node
	right *Node
}

func setup() Node {
	n := Node{"a", nil, nil}

	n.left = &Node{"c", nil, nil}
	n.left.left = &Node{"d", nil, nil}
	n.left.left.right = &Node{"h", nil, nil}
	n.left.right = &Node{"e", nil, nil}

	n.right = &Node{"b", nil, nil}
	n.right.left = &Node{"g", nil, nil}
	n.right.right = &Node{"f", nil, nil}

	return n
}

func printHelper(node Node, path string) {
	val := path + node.val

	if node.left != nil {
		printHelper(*node.left, val)
	}

	if node.right != nil {
		printHelper(*node.right, val)
	}

	if node.left == nil && node.right == nil {
		fmt.Println(val)
	}
}

func printPath(node Node) {
	printHelper(node, "")
}

func printTopToBottom(node *Node, prevs map[*Node]*Node) {
	cur := node
	stk := []string{cur.val}

	for cur != nil {
		cur = prevs[cur]
		if cur != nil {
			stk = append(stk, cur.val)
		}
	}

	i := len(stk)
	for i > 0 {
		i--
		fmt.Print(stk[i])
	}

	fmt.Println()
}

func printNoRecursion(node Node) {
	m := map[*Node]*Node{}
	nodes := []*Node{&node}
	leafs := []*Node{}
	cur := 0
	var n *Node

	for cur < len(nodes) {
		n = nodes[cur]
		if n.left != nil {
			m[n.left] = n
			nodes = append(nodes, n.left)
		}
		if n.right != nil {
			m[n.right] = n
			nodes = append(nodes, n.right)
		}

		if n.left == nil && n.right == nil {
			leafs = append(leafs, n)
		}
		cur++
	}

	for i := 0; i < len(leafs); i++ {
		printTopToBottom(leafs[i], m)
	}
}

func main() {
	node := setup()

	fmt.Println("Print reccursive")
	printPath(node)

	fmt.Println("Print flat")
	printNoRecursion(node)
}
