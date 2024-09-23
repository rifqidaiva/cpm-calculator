package pert

import (
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type node struct {
	position   rl.Vector2
	radius     float32
	color      color.RGBA
	dragging   bool
	dragOffset rl.Vector2
	number     int
	early      int
	late       int
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
	rl.DrawCircleLines(int32(n.position.X), int32(n.position.Y), n.radius, rl.Black)
	rl.DrawLineV(
		rl.Vector2Add(n.position, rl.NewVector2(-n.radius, 0)),
		rl.Vector2Add(n.position, rl.NewVector2(n.radius, 0)),
		rl.Black,
	)
	rl.DrawLineV(
		n.position,
		rl.Vector2Add(n.position, rl.NewVector2(0, n.radius)),
		rl.Black,
	)

	numberText := strconv.Itoa(n.number)
	earlyText := strconv.Itoa(n.early)
	lateText := strconv.Itoa(n.late)

	numberCenter := rl.Vector2Add(n.position, rl.NewVector2(0, -n.radius/2))
	earlyCenter := rl.Vector2Add(n.position, rl.NewVector2(-n.radius/2, n.radius/2))
	lateCenter := rl.Vector2AddValue(n.position, n.radius/2)

	numberPosition := getCenterText(comicSans, numberText, numberCenter, 24, 0)
	earlyPosition := getCenterText(comicSans, earlyText, earlyCenter, 24, 0)
	latePosition := getCenterText(comicSans, lateText, lateCenter, 24, 0)

	rl.DrawTextEx(comicSans, numberText, numberPosition, 24, 0, rl.Black)
	rl.DrawTextEx(comicSans, earlyText, earlyPosition, 24, 0, rl.Black)
	rl.DrawTextEx(comicSans, lateText, latePosition, 24, 0, rl.Black)
}

func newNode(position rl.Vector2, radius float32, number int) node {
	return node{
		position: position,
		radius:   radius,
		color:    rl.RayWhite,
		number:   number,
		early:    123,
		late:     123,
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

func getCenterText(font rl.Font, text string, centerPosition rl.Vector2, fontSize float32, spacing float32) rl.Vector2 {
	measuredText := rl.MeasureTextEx(font, text, fontSize, spacing)
	textPosition := rl.Vector2Subtract(centerPosition, rl.NewVector2(measuredText.X/2, measuredText.Y/2))

	return textPosition
}
