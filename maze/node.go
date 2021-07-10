package maze

import "fmt"

type Node struct {
	PositionX int
	PositionY int
}

func (node Node) print() {
	fmt.Printf("X:%d Y:%d", node.PositionX, node.PositionY)
}

func (node Node) Equals(deltaNode *Node) bool {
	return node.PositionX == deltaNode.PositionX &&
		node.PositionY == deltaNode.PositionY

}
