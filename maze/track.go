package maze

import "fmt"

type Track struct {
	name string
}

func (*Track) IsBlocked() bool {
	return false
}

func (path *Track) print() {
	fmt.Print(" ")
}

func (path *Track) Cost() int {
	return 1
}
