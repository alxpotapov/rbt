package rbt

// Comparison - результат сравнения
type Comparison int8

const (
	// IsLesser - меньше
	IsLesser Comparison = iota - 1
	// AreEqual - равны
	AreEqual
	// IsGreater - больше
	IsGreater
)

// Comparer - прототип функции сравнения
type Comparer func(interface{}, interface{}) Comparison
