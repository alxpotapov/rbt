package rbt

type color int

const (
	red color = iota
	black
)

func (c color) String() string {
	switch c {
	case black:
		return "b"
	}
	return "r"
}
