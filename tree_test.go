package rbt

import (
	"testing"
)


func TestNode_Insert(t *testing.T) {
	tree := &Tree{}
	tree.Insert("00", 0)
	tree.Insert("01", 1)
	tree.Insert("02", 2)
	tree.Insert("03", 3)
	tree.Insert("04", 4)
	tree.Insert("05", 5)
	tree.Insert("06", 6)
	tree.Insert("07", 7)
	tree.Insert("08", 8)
	tree.Insert("09", 9)
	tree.Insert("00", 0)
	tree.Insert("10", 10)
	tree.Insert("11", 11)
	if err := tree.root.checkViolation(); err != nil {
		t.Fatal(err)
	}
	t.Log(tree.String())
}
