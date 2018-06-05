package rbt

import (
	"fmt"
)

type node struct {
	color      color
	blacktoken bool
	key        interface{}
	value      interface{}
	children   [2]*node
}

var fakeBlackNode = &node{color: black}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
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
	return n.fixDoubleRed(offset)
}

func (n *node) fixDoubleRed(offset offset) *node {
	if n.children[offset].red() {
		return n.fixCase1(offset).fixCase2(offset)
	}
	return n
}

func (n *node) fixCase1(offset offset) *node {
	if n.children[offset].children[offset.other()].red() {
		n.children[offset] = n.children[offset].rotate(offset)
	}
	return n
}

func (n *node) fixCase2(offset offset) *node {
	if n.children[offset].children[offset].red() {
		n = n.rotate(offset.other())
		n.children[offset].color = black
	}
	return n
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

func (n *node) delete(key interface{}, comparer Comparer) *node {
	// если не найден узел, то возвращаем nil
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
	fmt.Println(n, offset)
	if n.children[offset] == nil {
		return n
	}

	if (n.children[offset].blacktoken || n.children[offset] == fakeBlackNode) && n.children[offset.other()] != nil {
		n = n. //fixCaseC(offset).
			fixCaseA(offset).
			fixCaseB(offset)
	}
	return n
}

func (n *node) fixCaseA(offset offset) *node {
	// The sibling of the doubly-black node is black and one nephew is red
	if (n.children[offset].blacktoken || n.children[offset] == fakeBlackNode) && n.children[offset.other()] != nil {
		if !n.children[offset.other()].red() {

			if n.children[offset.other()].children[offset].red() {
				fmt.Println("Case A", n, offset)
				n.children[offset.other()] = n.children[offset.other()].rotate(offset)
				n.children[offset.other()].color = black
				n.children[offset.other()].children[offset.other()].color = red
				fmt.Println("Result", n)
			}

			if n.children[offset.other()].children[offset.other()].red() {
				fmt.Println("Case A", n, offset)
				rootColor := n.color
				n = n.rotate(offset)
				n.color = rootColor
				n.children[offset].color = black
				if n.children[offset].children[offset] == fakeBlackNode {
					n.children[offset].children[offset] = nil
				}
				n.children[offset.other()].color = black

				//			n.children[offset].blacktoken = false
				fmt.Println("Result", n)
			}

		}
	}
	return n
}

func (n *node) fixCaseB(offset offset) *node {
	// The sibling and both nephews of the doubly-black node are black
	if (n.children[offset].blacktoken || n.children[offset] == fakeBlackNode) && n.children[offset.other()] != nil {
		if !n.children[offset.other()].red() {
			fmt.Println("Case B", n)
			if n.children[offset] == fakeBlackNode {
				n.children[offset] = nil
			} else {
				n.children[offset].blacktoken = false
			}

			if n.red() {
				n.color = black
			} else {
				n.blacktoken = true
			}
			n.children[offset.other()].color = red
			fmt.Println("Result", n)
		}
	}
	return n
}

func (n *node) fixCaseC(offset offset) *node {
	if n.children[offset.other()].red() {
		fmt.Println("Case C", n)
		n = n.rotate(offset)
		n.color = black
		n.children[offset].color = red
		fmt.Println("Result", n)
	}
	return n
}

//func (n *node) fixCaseA(offset offset) *node {
//	if n.children[offset.other()].red() {
//		fmt.Println("Case A", n)
//		n = n.rotate(offset)
//		n.color = black
//		n.children[offset].color = red
//		fmt.Println("Result", n)
//	}
//	return n
//}

//func (n *node) fixCaseB(offset offset) *node {
//	if n.children[offset.other()].color == black && !n.children[offset.other()].children[offset].red() && !n.children[offset.other()].children[offset.other()].red() {
//		fmt.Println("Case B", n)
//		n.blacktoken = true
//		n.children[offset].blacktoken = false
//		n.children[offset] = nil
//		//n.children[offset].color = black
//		n.children[offset.other()].color = red
//		fmt.Println("Result", n)
//	}
//	return n
//}

//func (n *node) fixCaseC(offset offset) *node {
//	if n.children[offset.other()].color == black && n.children[offset.other()].children[offset].red() {
//		fmt.Println("Case C", n)
//		n.children[offset.other()] = n.children[offset.other()].rotate(offset.other())
//		n.children[offset.other()].color = black
//		n.children[offset.other()].children[offset.other()].color = red
//		fmt.Println("Result", n)
//	}
//	return n
//}

//func (n *node) fixCaseD(offset offset) *node {
//	if n.children[offset.other()].color == black && n.children[offset.other()].children[offset.other()].red() {
//		fmt.Println("Case D", n)
//		rootColor := n.color
//		n = n.rotate(offset)
//		n.color = rootColor
//		n.children[offset.other()].color = black
//		fmt.Println("Result", n)
//	}
//	return n
//}

// String - вывод на экран
func (n *node) String() string {
	if n == nil {
		return ""
	}
	s := ""
	s += n.children[left].String()
	s += fmt.Sprintf("%v:%s:%v", n.key, n.color, n.blacktoken)
	s += n.children[right].String()
	return "(" + s + ")"
}

func (n *node) splice(comparer Comparer) *node {
	//Удалить узел и вернуть nil
	if n.children[left] == nil && n.children[right] == nil {
		// удаляемый узел может быть как красным, так и черным
		// если узел черный, то делаем его фиктивным
		if n.red() {
			return nil
		}
		return fakeBlackNode
	}
	//Удалить узел и вернуть его левую подветвь
	if n.children[right] == nil {
		// остающийся узел всегда красный и его перекрашиваем
		n.children[left].color = black
		return n.children[left]
	}
	//Удалить узел и вернуть его правую подветвь
	if n.children[left] == nil {
		// остающийся узел всегда красный и его перекрашиваем
		n.children[right].color = black
		return n.children[right]
	}
	//Заменить значение текущего узла на максимум левой подветви
	//Удалить максимум левой подветви
	//Вернуть собранное значение
	tempNode := n.children[left].findMax()
	n.key = tempNode.key
	n.value = tempNode.value
	fmt.Println("delete proc", n.children[left])
	n.children[left] = n.children[left].delete(tempNode.key, comparer)
	return n.fixDoubleBlackViolation(left)
}

// findMax - вернуть узел с максимальным значением из левой подветви
func (n *node) findMax() *node {
	if n.children[right] != nil {
		n.children[right] = n.children[right].findMax()
	}
	return n
}

func (n *node) red() bool {
	return n != nil && n.color == red
}

func (n *node) rotate(offset offset) *node {
	root := n.children[offset.other()]
	n.children[offset.other()] = root.children[offset]
	root.children[offset] = n
	return root
}
