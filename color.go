package rbt

type color int

const (
	red color = iota
	black
	doubleblack
)

func (c *color) increment() {
	*c++
}

func (c color) String() string {
	switch c {
	case black:
		return "b"
	case doubleblack:
		return "bb"
	}
	return "r"
}
