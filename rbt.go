package rbt

import (
	"strings"
	"fmt"
)

// Callback ...
type Callback func(node *Node)

type color int

const (
	black color = iota
	red
)
// String - color as string
func (c color) String() string {
	if c == black {
		return "b"
	}
	return "r"
}

type child int

const (
	left  child  = iota
	right
)

func (c child) other() child {
	if c == right {
		return left
	}
	return right
}
// String - child as string
func (c child) String() string {
	if c == right {
		return "r"
	}
	return "l"
}

// Node contains the key as string, value as interface{}
type Node struct {
	color color
	child []*Node
	Key   string
	Value interface{}
}

func newNode(key string, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
		child: make([]*Node, 2),
		color: red,
	}
}

// Insert a key and value into a sorted tree
func (n *Node) Insert(key string, value interface{}) *Node {
	if n == nil {
		return newNode(key, value)
	}
	comp := strings.Compare(n.Key, key)
	switch {
	case comp > 0:
		n.child[left] = n.child[left].Insert(key, value)
		return n.fixViolation(left)
	case comp < 0:
		n.child[right] = n.child[right].Insert(key, value)
		return n.fixViolation(right)
	default:
		n.Value = value
	}
	return n
}

// String - show tree as a string
func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.child[left] != nil {
		s += n.child[left].String() + " "
	}
	s += fmt.Sprintf("key=%s%s;val=%v", n.Key, n.color, n.Value)
	if n.child[right] != nil {
		s += " " + n.child[right].String()
	}
	return "(" + s + ")"
}

func (n *Node) fixViolation(direction child) *Node {
	if n.child[direction].isRed() {
		if n.child[direction.other()].isRed() {
			n.flipColor()
			return n
		}
		if n.child[direction].child[direction].isRed() {
			return n.singleRotate(direction.other())
		}
		if n.child[direction].child[direction.other()].isRed() {
			return n.doubleRotate(direction.other())
		}
	}
	return n
}

func (n *Node) flipColor() {
	n.color = red
	n.child[right].color, n.child[left].color = black, black
}

func (n *Node) singleRotate(direction child) *Node {
	node := n.child[direction.other()]
	n.child[direction.other()] = node.child[direction]
	node.child[direction] = n
	n.color, node.color= red, black
	return node
}

func (n *Node) doubleRotate(direction child) *Node {
	n.child[direction.other()] = n.child[direction.other()].singleRotate(direction.other())
	node := n.singleRotate(direction)
	return node
}

func (n *Node) isRed() bool {
	return (n != nil) && (n.color == red)
}


func(n *Node) checkViolation()error {
	if n == nil {
		return fmt.Errorf("tree is empty")
	}
	// property 2
	if n.color != black {
		return fmt.Errorf("property 2 is violated: the root of the tree should be black")
	}
	// property 3
	node := newNode("key", "value")
	if (node.child[left] == nil &&  node.child[left].isRed()) || (node.child[right] == nil && node.child[right].isRed()) {
		return fmt.Errorf("property 3 is violated: every leaf (nil) is black")
	}
	// property 4
	if !n.checkColor(n) {
		return fmt.Errorf("property 4 is violated: if a node is red, then both its children should be black")
	}
	// property 5
	lh := n.getHeight(n, left)
	rh := n.getHeight(n, right)
	if lh != rh {
		return fmt.Errorf("property 5 is violated: from node to leaves different height l= %d, r = %d", lh, rh)
	}
	return nil
}

func(n *Node)getHeight(node*Node, direction child) int {
	if node == nil{
		return 0
	}
	if node.color==black {
		return 1 + n.getHeight(node.child[direction], direction)
	}
	return n.getHeight(node.child[direction], direction)
}


func(n *Node)checkColor(node*Node) bool {
	check := true
	if node == nil{
		return check
	}
	if node.color==red {
		return  !node.child[right].isRed() && !node.child[left].isRed()
	}
	if node.child[left] != nil {
		check = check && n.checkColor(node.child[left])
	}
	if node.child[right] != nil {
		check = check && n.checkColor(node.child[right])
	}
	return check
}

