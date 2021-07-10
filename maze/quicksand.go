package maze

import "github.com/fatih/color"

type Quicksand struct {
	name string
}

func (*Quicksand) IsBlocked() bool {
	return false
}

func (path *Quicksand) print() {
	color.New(color.BgYellow).Print(" ")
}

func (path *Quicksand) Cost() int {
	return 15
}
