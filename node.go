package rbt

import (
	"fmt"
)

type node struct {
	color    color
	key      interface{}
	value    interface{}
	children [2]*node
}

var fakeBlackNode = &node{color: black}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
}

func (n *node) find(key interface{}, comparer Comparer) (interface{}, bool) {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return nil, false
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.value, true
	}
	return n.children[offset].find(key, comparer)
}

func (n *node) insert(key, value interface{}, comparer Comparer) *node {
	// если не найден узел, то возвращаем новый
	if n == nil {
		return newNode(key, value)
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		// либо заменяем value в случае set(map),
		// либо добавляем в список в случае multiset(multimap)
		return n
	}
	n.children[offset] = n.children[offset].insert(key, value, comparer)
	return n.fixDoubleRedViolation(offset)
}

func (n *node) fixDoubleRedViolation(offset offset) *node {
	return n.
		case1(offset).
		case2(offset).
		case3(offset)
}

func (n *node) case1(offset offset) *node {
	if n.children[offset].red() && n.children[offset.other()].red() {
		//		fmt.Println("Case 1 (Both children are red): recolour")
		//		fmt.Println("- src", n)
		n.color = red
		n.children[offset].color = black
		n.children[offset.other()].color = black
		//		fmt.Println("- dst", n)
	}
	return n
}

func (n *node) case2(offset offset) *node {
	if n.children[offset].red() &&
		n.children[offset].children[offset.other()].red() {
		//		fmt.Println("Case 2 (The parent has an other child and are red): rotate", offset, "around child")
		//		fmt.Println("- src", n)
		n.children[offset] = n.children[offset].rotate(offset)
		//		fmt.Println("- dst", n)
	}
	return n
}

func (n *node) case3(offset offset) *node {
	if n.children[offset].red() &&
		n.children[offset].children[offset].red() {
		//		fmt.Println("Case 3 (The parent has a same child and are red): rotate", offset.other(), "around parent")
		//		fmt.Println("- src", n)
		n = n.rotate(offset.other())
		//		fmt.Println("- dst", n)
		//		fmt.Println("and recolour", n)
		//		fmt.Println("- src", n)
		n.color = black
		n.children[offset.other()].color = red
		//		fmt.Println("- dst", n)
	}
	return n
}

func (n *node) delete(key interface{}, comparer Comparer) *node {
	if n == nil {
		return nil
	}
	var offset offset
	switch comparer(n.key, key) {
	case IsGreater:
		offset = left
	case IsLesser:
		offset = right
	case AreEqual:
		return n.splice(comparer)
	}
	n.children[offset] = n.children[offset].delete(key, comparer)
	return n.fixDoubleBlackViolation(offset)
}

func (n *node) fixDoubleBlackViolation(offset offset) *node {
	return n.
		caseA(offset).
		caseB(offset).
		caseC(offset).
		caseD(offset)
}

func (n *node) caseA(offset offset) *node {
	if n.blackToken(offset) &&
		n.children[offset.other()].black() &&
		n.children[offset.other()].children[offset.other()].red() {

		//		fmt.Println("Case A (The sibling is black, but hi has other red child): rotate", offset, "around root")
		//		fmt.Println("- src", n)
		rootColor := n.color
		n = n.rotate(offset)
		//		fmt.Println("- dst", n)
		//		fmt.Println("and recolour")
		//		fmt.Println("- src", n)
		n.color = rootColor
		n.children[offset].color, n.children[offset.other()].color = black, black
		if n.children[offset].children[offset] == fakeBlackNode {
			n.children[offset].children[offset] = nil
		} else {
			n.children[offset].children[offset].color = black
		}
		//		fmt.Println("- dst", n)

	}
	return n

}

func (n *node) caseB(offset offset) *node {
	if n.blackToken(offset) &&
		n.children[offset.other()].black() &&
		n.children[offset.other()].children[offset].red() {

		//		fmt.Println("Case B (The sibling is black, but he has same red children): rotate", offset, "around sibling")
		//		fmt.Println("- src", n)
		n.children[offset.other()] = n.children[offset.other()].rotate(offset)
		//		fmt.Println("- dst", n)
		//		fmt.Println("and recolour", n)
		//		fmt.Println("- src", n)
		n.children[offset.other()].color = black
		n.children[offset.other()].children[offset.other()].color = red
		//		fmt.Println("- dst", n)
		return n.caseA(offset)
	}
	return n
}

func (n *node) caseC(offset offset) *node {
	if n.blackToken(offset) &&
		n.children[offset.other()].black() {

		//		fmt.Println("Case C (The sibling is black): recolour")
		//		fmt.Println("- src", n)
		n.color.increment()
		n.children[offset.other()].color = red
		if n.children[offset] == fakeBlackNode {
			n.children[offset] = nil
		} else {
			n.children[offset].color = black
		}
		//		fmt.Println("- dst", n)
	}
	return n
}

func (n *node) caseD(offset offset) *node {
	if n.blackToken(offset) &&
		n.children[offset.other()].red() {

		//		fmt.Println("Case D (The sibling is red): rotate", offset, "around root")
		//		fmt.Println("- src", n)
		n = n.rotate(offset)
		//		fmt.Println("- dst", n)
		//		fmt.Println("and recolour")
		//		fmt.Println("- src", n)
		n.color = black
		n.children[offset].color = red
		//		fmt.Println("- dst", n)
		n.children[offset] = n.children[offset].fixDoubleBlackViolation(offset)
	}
	return n
}

// String ...
func (n *node) String() string {
	if n == nil {
		return ""
	}
	s := ""
	s += n.children[left].String()
	s += fmt.Sprintf("%v:%s", n.key, n.color)
	s += n.children[right].String()
	return "(" + s + ")"
}

func (n *node) splice(comparer Comparer) *node {
	if n.children[left] == nil && n.children[right] == nil {
		if n.red() {
			return nil
		}
		return fakeBlackNode
	}
	if n.children[right] == nil {
		n.children[left].color = black
		return n.children[left]
	}
	if n.children[left] == nil {
		n.children[right].color = black
		return n.children[right]
	}
	tempNode := n.children[left].findMax()
	n.key = tempNode.key
	n.value = tempNode.value
	n.children[left] = n.children[left].delete(tempNode.key, comparer)
	return n.fixDoubleBlackViolation(left)
}

// findMax ...
func (n *node) findMax() *node {
	if n.children[right] != nil {
		n.children[right] = n.children[right].findMax()
	}
	return n
}

func (n *node) red() bool {
	return n != nil && n.color == red
}

func (n *node) black() bool {
	return n != nil && n.color == black
}

func (n *node) blackToken(offset offset) bool {
	return n.children[offset] != nil &&
		(n.children[offset] == fakeBlackNode ||
			n.children[offset].color == doubleBlack
			)
}

func (n *node) rotate(offset offset) *node {
	root := n.children[offset.other()]
	n.children[offset.other()] = root.children[offset]
	root.children[offset] = n
	return root
}

func (n *node) blackHeight() int {
	if n == nil {
		return 1
	}
	leftBlackHeight := n.children[left].blackHeight()
	if leftBlackHeight == 0 {
		return leftBlackHeight
	}
	rightBlackHeight := n.children[right].blackHeight()
	if rightBlackHeight == 0 {
		return rightBlackHeight
	}
	if leftBlackHeight != rightBlackHeight {
		return 0
	} else {
		if n.black() {
			leftBlackHeight++
		}
		return leftBlackHeight
	}
}
