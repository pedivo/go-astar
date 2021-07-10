package algorithms

import (
	"pedro.nevesmacha.do/maze"
)

type AStar struct {
	openPaths []maze.Path
	usedPaths []maze.Node
}

func (aStar *AStar) aStar(originalMaze maze.Maze, startNode maze.Node, endNode maze.Node) (int, maze.Maze) {
	currentMaze := originalMaze.Clone()
	startStructure := currentMaze.GetStructure(&startNode)

	if startStructure.IsBlocked() {
		println("Maze starting in a blocked structure")
		return 0, currentMaze
	}

	startPath := maze.Path{
		Node:      &startNode,
		TotalCost: 0,
	}

	aStar.openPaths = startPath.FindOpenNeighbors(&currentMaze)
	startPath = aStar.openPaths[0]
	aStar.addTrack(&currentMaze, &startPath)

	optimizedPath := aStar.findBestPath(&currentMaze, startPath, endNode)

	originalMaze.DrawPath(&optimizedPath)

	return 0, originalMaze
}

func (aStar *AStar) findBestPath(currentMaze *maze.Maze, currentPath maze.Path, endNode maze.Node) maze.Path {
	aStar.openPaths = append(aStar.openPaths, currentPath.FindOpenNeighbors(currentMaze)...)
	aStar.openPaths = currentPath.SortPaths(aStar.openPaths)

	aStar.addTrack(currentMaze, &aStar.openPaths[0])

	if isEndNode(aStar.openPaths[0], endNode) {
		return aStar.openPaths[0]
	}

	return aStar.findBestPath(currentMaze, aStar.openPaths[0], endNode)
}

func isEndNode(path maze.Path, endNode maze.Node) bool {
	return path.Node.Equals(&endNode)
}

func (aStar *AStar) addTrack(currentMaze *maze.Maze, path *maze.Path) {
	var nonUsesOpenPaths []maze.Path

	for index, openPath := range aStar.openPaths {
		if index == 0 {
			continue
		}

		var isPathUsed = false
		for _, usedPaths := range aStar.usedPaths {
			if usedPaths.Equals(openPath.Node) {
				isPathUsed = true
			}
		}

		for _, nonUsesOpenPath := range nonUsesOpenPaths {
			if nonUsesOpenPath.Node.Equals(openPath.Node) {
				isPathUsed = true
			}
		}

		if !isPathUsed {
			nonUsesOpenPaths = append(nonUsesOpenPaths, openPath)
		}
	}

	aStar.openPaths = nonUsesOpenPaths

	currentMaze.Definition[path.Node.PositionX][path.Node.PositionY] = path
}
