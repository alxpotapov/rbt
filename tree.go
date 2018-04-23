package rbt

import (
	"fmt"
)

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key string, value interface{}) {
	t.root = t.root.Insert(key, value)
	t.root.color = black
}

func (t *Tree) String() string {
	if t.root == nil {
		return "()"
	}
	return t.root.String()
}

func(t *Tree) checkViolation()error {
	if t.root == nil {
		return fmt.Errorf("tree is empty")
	}
	// property 2
	if t.root.color != black {
		return fmt.Errorf("property 2 is violated: the root of the tree should be black")
	}
	// property 3
	n := NewNode("key", "value")
	if (n.child[left] == nil &&  n.child[left].isRed()) || (n.child[right] == nil && n.child[right].isRed()) {
		return fmt.Errorf("property 3 is violated: every leaf (nil) is black")
	}
	// property 4
	if !t.checkColor(t.root) {
		return fmt.Errorf("property 4 is violated: if a node is red, then both its children should be black")
	}
	// property 5
	lh := t.getHeight(t.root, left)
	rh := t.getHeight(t.root, right)
	if lh != rh {
		return fmt.Errorf("property 5 is violated: from node to leaves different height l= %d, r = %d", lh, rh)
	}
	return nil
}

func(t *Tree)getHeight(node*Node, direction child) int {
	if node == nil{
		return 0
	}
	if node.color==black {
		return 1 + t.getHeight(node.child[direction], direction)
	}
	return t.getHeight(node.child[direction], direction)
}


func(t *Tree)checkColor(node*Node) bool {
	check := true
	if node == nil{
		return check
	}
	if node.color==red {
		return  !node.child[right].isRed() && !node.child[left].isRed()
	}
	if node.child[left] != nil {
		check = check && t.checkColor(node.child[left])
	}
	if node.child[right] != nil {
		check = check && t.checkColor(node.child[right])
	}
	return check
}
