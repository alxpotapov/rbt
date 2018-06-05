# BST
## Usage
```
go get github.com/alxpotapov/bst
```
## A simple implementation of basic binary search methods:
 * insert key and its value
 * find by key
 * delete by key
### Create tree
```
tree := NewTree(func(f, s interface{}) Comparison {
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

```
### Insert nodes
```
tree.Insert("A", "AA")
tree.Insert("L", "LL")
tree.Insert("G", "GG")
fmt.Println(tree)
```
### Find value by key
```
if value, found := tree.Find("L"); found {
	// ...
}
```
### Delete value by key
```
tree.Delete("A")
fmt.Println(tree)
```
