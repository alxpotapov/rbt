package rbt
// Tree ...
type Tree struct {
	root *Node
}
// Insert - insert/update key and value to tree
func (t *Tree) Insert(key string, value interface{}) {
	t.root = t.root.Insert(key, value)
	t.root.color = black
}
// String - tree as string
func (t *Tree) String() string {
	return t.root.String()
}



