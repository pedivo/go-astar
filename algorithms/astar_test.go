package algorithms

import (
	"pedro.nevesmacha.do/maze"
	"testing"
)

func TestBasicMazeSolver(t *testing.T) {
	currentMaze := maze.Maze{
		Definition: [][]maze.Structure{
			{&maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
		},
	}

	aStar := AStar{}
	_, resultMaze := aStar.aStar(currentMaze, maze.Node{0, 0}, maze.Node{9, 9})

	resultMaze.Print()
}

func TestMazeSolverComplexStructures(t *testing.T) {
	currentMaze := maze.Maze{
		Definition: [][]maze.Structure{
			{&maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Quicksand{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}},
			{&maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
			{&maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Track{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Wall{}, &maze.Track{}, &maze.Track{}},
		},
	}

	aStar := AStar{}
	_, resultMaze := aStar.aStar(currentMaze, maze.Node{0, 0}, maze.Node{9, 9})

	resultMaze.Print()
}
