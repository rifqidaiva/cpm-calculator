package pert

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type node struct {
	position   rl.Vector2
	radius     float32
	color      color.RGBA
	dragging   bool
	dragOffset rl.Vector2
}

func (n *node) update() bool {
	mousePosition := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if rl.CheckCollisionPointCircle(mousePosition, n.position, n.radius) {
			n.dragging = true
			n.dragOffset = rl.Vector2Subtract(mousePosition, n.position)
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		n.dragging = false
	}

	if n.dragging {
		n.position = rl.Vector2Subtract(mousePosition, n.dragOffset)
	}

	return n.dragging
}

func (n node) render() {
	rl.DrawCircleV(n.position, n.radius, n.color)
}

func newNode(position rl.Vector2, radius float32) node {
	return node{
		position: position,
		radius:   radius,
		color: color.RGBA{
			uint8(rl.GetRandomValue(0, 255)),
			uint8(rl.GetRandomValue(0, 255)),
			uint8(rl.GetRandomValue(0, 255)),
			255,
		},
	}
}

// Function to handle node selection and bring the selected node to the top.
func handleNodeSelection(nodes *[]node) {
	mousePosition := rl.GetMousePosition()

	for i := len(*nodes) - 1; i >= 0; i-- {
		if rl.CheckCollisionPointCircle(mousePosition, (*nodes)[i].position, (*nodes)[i].radius) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			selectedNode := (*nodes)[i]
			*nodes = append((*nodes)[:i], (*nodes)[i+1:]...)
			*nodes = append(*nodes, selectedNode)
			break
		}
	}
}
