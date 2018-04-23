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
	red   color = iota
)

type child int

const (
	left  child = 0
	right child = 1
)

func (c color) String() string {
	if c == black {
		return "b"
	}
	return "r"
}

func (c child) other() child {
	if c == right {
		return left
	}
	return right
}

func (c child) String() string {
	if c == right {
		return "r"
	}
	return "l"
}

// Node contains the key as string, value as interface{} and 2 pointers to his children
type Node struct {
	color color
	child []*Node
	Key   string
	Value interface{}
}

func NewNode(key string, value interface{}) *Node {
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
		return NewNode(key, value)
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


