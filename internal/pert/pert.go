package pert

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 1080
	windowHeight = 720
)

func CreatePert(pathListData [][]string) {
	rl.InitWindow(windowWidth, windowHeight, "PERT Diagram")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	nodes := []node{
		newNode(rl.Vector2{X: 100, Y: 100}, 50),
		newNode(rl.Vector2{X: 300, Y: 100}, 50),
		newNode(rl.Vector2{X: 500, Y: 100}, 50),
	}

	for !rl.WindowShouldClose() {
		handleInputKey()
		handleNodeSelection(&nodes)

		// Update from top to bottom to find the topmost selected node.
		for i := len(nodes) - 1; i >= 0; i-- {
			if dragging := nodes[i].update(); dragging {
				break
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for i := range nodes {
			nodes[i].render()
		}

		rl.EndDrawing()
	}
}
