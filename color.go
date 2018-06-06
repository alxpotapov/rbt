package rbt

type color int

const (
	red color = iota
	black
	doubleBlack
)

func (c color) String() string {
	switch c {
	case black:
		return "b"
	case doubleBlack:
		return "bb"
	}
	return "r"
}

func (c *color) increment() {
	*c++
}
