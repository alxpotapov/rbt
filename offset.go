package rbt

type offset int

const (
	left offset = iota
	right
)

func (o offset) other() offset {
	return o ^ right
}

func (o offset) String() string {
	if o == left {
		return "left"
	}
	return "right"
}
