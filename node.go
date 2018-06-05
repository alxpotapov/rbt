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
		switch {
		// дядя красный
		case n.children[offset.other()].red():
			return n.fixCase1(offset)
		// является "противоположным" сыном
		case n.children[offset].children[offset.other()].red():
			return n.fixCase2a(offset).fixCase2b(offset)
		// является "тем же" сыном
		case n.children[offset].children[offset].red():
			return n.fixCase2b(offset)
		}
	}
	return n
}

func (n *node) fixCase1(offset offset) *node {
	// recolour
	n.color = red
	n.children[left].color = black
	n.children[right].color = black
	return n
}

func (n *node) fixCase2a(offset offset) *node {
	// restructure
	// вращаем в сторону вокруг отца и делаем "того же" сына
	n.children[offset] = n.children[offset].rotate(offset)
	return n
}

func (n *node) fixCase2b(offset offset) *node {
	// restructure
	// вращаем в протовоположную сторону вокруг деда
	n = n.rotate(offset.other())
	n.color = black
	n.children[offset.other()].color = red
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
	fmt.Println("fixDoubleBlackViolation", n)
	if n.children[offset] == fakeBlackNode {
		fmt.Println("Create double black", n)
		n.color.increment()
		n.children[offset] = nil
		fmt.Println("Result", n)
		return n
	}

	if n.children[offset].color == doubleblack && n.children[offset.other()] != nil {
		return n.fixCaseA(offset).
			fixCaseB(offset).
			fixCaseC(offset).
			fixCaseD(offset)
	}

	return n
}

func (n *node) fixCaseA(offset offset) *node {
	if n.children[offset.other()].red() {
		fmt.Println("Case A", n)
		n = n.rotate(offset)
		n.color = black
		n.children[offset].color = red
		fmt.Println("Result", n)
	}
	return n
}

func (n *node) fixCaseB(offset offset) *node {
	if n.children[offset.other()].color == black {
		fmt.Println("Case B", n)
		n.color.increment()
		n.children[offset].color = black
		n.children[offset.other()].color = red
		fmt.Println("Result", n)
	}
	return n
}

func (n *node) fixCaseC(offset offset) *node {
	if n.children[offset.other()].color == black && n.children[offset.other()].children[offset].red() {
		fmt.Println("Case C", n)
		n.children[offset.other()] = n.children[offset.other()].rotate(offset.other())
		n.children[offset.other()].color = black
		n.children[offset.other()].children[offset.other()].color = red
		fmt.Println("Result", n)
	}
	return n
}

func (n *node) fixCaseD(offset offset) *node {
	if n.children[offset.other()].color == black && n.children[offset.other()].children[offset.other()].red() {
		fmt.Println("Case D", n)
		rootColor := n.color
		n = n.rotate(offset)
		n.color = rootColor
		n.children[offset.other()].color = black
		fmt.Println("Result", n)
	}
	return n
}

// String - вывод на экран
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
	n.children[left] = n.children[left].delete(n.key, comparer)
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
