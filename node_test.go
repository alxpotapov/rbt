package rbt

import (
	"testing"
)

var n *node
var comparer = func(f, s interface{}) Comparison {
	intF, ok := f.(string)
	if !ok {
		return IsLesser
	}
	intS, ok := s.(string)
	if !ok {
		return IsLesser
	}
	switch {
	case intF < intS:
		return IsLesser
	case intF > intS:
		return IsGreater
	}
	return AreEqual
}

func TestInsert(t *testing.T) {

	n = n.insert("A", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("L", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("G", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("O", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("R", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("I", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("T", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("H", nil, comparer)
	n.color = black
	t.Log(n)
	n = n.insert("M", nil, comparer)
	n.color = black
	t.Log(n)
}

//func TestFind(t *testing.T) {
//	t.Log("Find A")
//	if _, found := n.find("A", comparer); !found {
//		t.Fatal("expected found")
//	}
//	t.Log("Find X")
//	if _, found := n.find("X", comparer); found {
//		t.Fatal("expected not found")
//	}
//	t.Log("Find L")
//	if _, found := n.find("L", comparer); !found {
//		t.Fatal("expected found")
//	}
//	t.Log("Find M")
//	if _, found := n.find("M", comparer); !found {
//		t.Fatal("expected found")
//	}
//	t.Log("Find W")
//	if _, found := n.find("W", comparer); found {
//		t.Fatal("expected not found")
//	}

//}

func TestDelete(t *testing.T) {
	t.Log("Delete A")
	n = n.delete("A", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete L")
	n = n.delete("L", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete G")
	n = n.delete("G", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete O")
	n = n.delete("O", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete R")
	n = n.delete("R", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete I")
	n = n.delete("I", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete T")
	n = n.delete("T", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete H")
	n = n.delete("H", comparer)
	n.color = black
	t.Log(n)
	t.Log("Delete M")
	n = n.delete("M", comparer)
	n.color = black
	t.Log(n)
}
