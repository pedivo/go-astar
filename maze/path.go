package maze

import (
	"github.com/fatih/color"
	"sort"
)

type Path struct {
	name      string
	Node      *Node
	parent    *Path
	TotalCost int
}

type ByCost []Path

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCost) Less(i, j int) bool { return a[i].TotalCost < a[j].TotalCost }

func (*Path) IsBlocked() bool {
	return true
}

func (*Path) Cost() int {
	return 0
}

func (path *Path) print() {
	color.New(color.BgHiWhite).Print(" ")
}

func (path Path) FindOpenNeighbors(maze *Maze) []Path {
	var openPaths []Path

	for mazeRowDelta := -1; mazeRowDelta <= 1; mazeRowDelta++ {
		for mazeColumnDelta := -1; mazeColumnDelta <= 1; mazeColumnDelta++ {
			if mazeRowDelta == 0 && mazeColumnDelta == 0 {
				continue
			}

			nodeDelta := &Node{
				PositionX: path.Node.PositionX + mazeRowDelta,
				PositionY: path.Node.PositionY + mazeColumnDelta,
			}

			structure := maze.GetStructure(nodeDelta)

			if structure == nil || structure.IsBlocked() || path.Node.Equals(nodeDelta) {
				continue
			}

			openPaths = append(
				openPaths,
				Path{
					Node:      nodeDelta,
					parent:    &path,
					TotalCost: path.TotalCost + structure.Cost(),
				},
			)
		}
	}

	sort.Sort(ByCost(openPaths))

	return openPaths
}

func (*Path) SortPaths(path []Path) []Path {
	sort.Sort(ByCost(path))
	return path
}
