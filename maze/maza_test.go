package maze

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	maze := &Maze{
		Definition: [][]Structure{
			{&Track{}, &Wall{}, &Wall{}, &Wall{}, &Wall{}, &Wall{}, &Track{}, &Track{}, &Track{}, &Track{}},
			{&Track{}, &Track{}, &Track{}, &Track{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Track{}},
			{&Track{}, &Track{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Track{}, &Track{}, &Track{}},
			{&Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Wall{}, &Track{}},
			{&Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Track{}},
			{&Track{}, &Wall{}, &Wall{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}},
			{&Track{}, &Wall{}, &Track{}, &Track{}, &Track{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Wall{}},
			{&Track{}, &Wall{}, &Track{}, &Wall{}, &Wall{}, &Wall{}, &Track{}, &Wall{}, &Track{}, &Track{}},
			{&Wall{}, &Track{}, &Track{}, &Track{}, &Track{}, &Track{}, &Track{}, &Wall{}, &Track{}, &Track{}},
			{&Track{}, &Track{}, &Track{}, &Track{}, &Wall{}, &Wall{}, &Wall{}, &Wall{}, &Track{}, &Track{}},
		},
	}

	maze.Print()
}
