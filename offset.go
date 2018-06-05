package rbt

type offset int

const (
	left offset = iota
	right
)

func (o offset) other() offset {
	return o ^ right
}
