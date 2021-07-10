package maze

import "fmt"

type Maze struct {
	Definition [][]Structure
}

func (maze *Maze) Print() {
	for _, row := range maze.Definition {
		for _, structure := range row {
			structure.print()
		}
		fmt.Println("")
	}
}

func (maze Maze) GetStructure(node *Node) Structure {
	if (node.PositionX < 0 || node.PositionY < 0) ||
		(node.PositionX >= len(maze.Definition) || node.PositionY >= len(maze.Definition[node.PositionX])) {
		return nil
	}

	return maze.Definition[node.PositionX][node.PositionY]
}

func (maze *Maze) Clone() Maze {
	newMaze := Maze{}

	for _, structures := range maze.Definition {
		newLine := make([]Structure, len(structures))
		copy(newLine, structures)
		newMaze.Definition = append(newMaze.Definition, newLine)
	}

	return newMaze
}

func (maze *Maze) DrawPath(path *Path) {

	maze.Definition[path.Node.PositionX][path.Node.PositionY] = path

	if path.parent != nil {
		maze.DrawPath(path.parent)
	}
}
