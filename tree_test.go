package rbt

import (
	"testing"
)

var tree = NewTree(func(f, s interface{}) Comparison {
	intF, _ := f.(string)
	intS, _ := s.(string)
	switch {
	case intF < intS:
		return IsLesser
	case intF > intS:
		return IsGreater
	}
	return AreEqual
})

func TestTreeInsert(t *testing.T) {
	tree.Insert("A", nil)
	t.Log(tree)
	tree.Insert("L", nil)
	t.Log(tree)
	tree.Insert("G", nil)
	t.Log(tree)
	tree.Insert("O", nil)
	t.Log(tree)
	tree.Insert("R", nil)
	t.Log(tree)
	tree.Insert("I", nil)
	t.Log(tree)
	tree.Insert("T", nil)
	t.Log(tree)
	tree.Insert("H", nil)
	t.Log(tree)
	tree.Insert("M", nil)
	t.Log(tree)
}

func TestTreeFind(t *testing.T) {
	t.Log("Find A")
	if _, found := tree.Find("A"); found {
		t.Log("found")
	}
	t.Log("Find X")
	if _, found := tree.Find("X"); !found {
		t.Log("not found")
	}
	t.Log("Find L")
	if _, found := tree.Find("L"); found {
		t.Log("found")
	}
	t.Log("Find M")
	if _, found := tree.Find("M"); found {
		t.Log("found")
	}
	t.Log("Find W")
	if _, found := tree.Find("W"); !found {
		t.Log("not found")
	}

}

//func TestTreeDelete(t *testing.T) {
//	t.Log("Delete A")
//	tree.Delete("A")
//	t.Log(tree)
//	t.Log("Delete L")
//	tree.Delete("L")
//	t.Log(tree)
//	t.Log("Delete G")
//	tree.Delete("G")
//	t.Log(tree)
//	t.Log("Delete O")
//	tree.Delete("O")
//	t.Log(tree)
//	t.Log("Delete R")
//	tree.Delete("R")
//	t.Log(tree)
//	t.Log("Delete I")
//	tree.Delete("I")
//	t.Log(tree)
//	t.Log("Delete T")
//	tree.Delete("T")
//	t.Log(tree)
//	t.Log("Delete H")
//	tree.Delete("H")
//	t.Log(tree)
//	t.Log("Delete M")
//	tree.Delete("M")
//	t.Log(tree)
//}

func TestTreeClear(t *testing.T) {
	tree.Insert("A", nil)
	t.Log(tree)
	tree.Insert("L", nil)
	t.Log(tree)
	tree.Clear()
	t.Log(tree)

}

func TestTreeEmpty(t *testing.T) {
	if empty := tree.Empty(); empty {
		t.Log("tree is empty")
	}
	t.Log("Insert A")
	tree.Insert("A", nil)
	if empty := tree.Empty(); !empty {
		t.Log("tree is not empty")
	}
	t.Log("Delete A")
	tree.Delete("A")
	if empty := tree.Empty(); empty {
		t.Log("tree is empty")
	}
}
