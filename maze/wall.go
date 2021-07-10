package maze

import (
	"github.com/fatih/color"
)

type Wall struct {
}

func (wall *Wall) IsBlocked() bool {
	return true
}

func (*Wall) print() {
	color.New(color.BgRed).Print(" ")
}

func (*Wall) Cost() int {
	return 99999999
}
