package pert

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 1080
	windowHeight = 720
)

var (
	comicSans  rl.Font
	nodes      []node
	activities []activity
)

func CreatePert(pathListData [][]string) {
	rl.InitWindow(windowWidth, windowHeight, "PERT Diagram")
	rl.SetTargetFPS(60)
	comicSans = rl.LoadFont("assets/fonts/opensans_regular.ttf")

	defer rl.CloseWindow()

	nodes = []node{
		newNode(rl.Vector2{X: 100, Y: 100}, 70, 1),
		newNode(rl.Vector2{X: 300, Y: 100}, 70, 2),
		newNode(rl.Vector2{X: 500, Y: 100}, 70, 3),
		newNode(rl.Vector2{X: 700, Y: 100}, 70, 4),
		newNode(rl.Vector2{X: 900, Y: 100}, 70, 5),
	}

	activities = []activity{
		newActivity(1, 2, "A", 10),
		newActivity(1, 3, "A", 10),
		newActivity(1, 4, "A", 10),
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

		for i := range activities {
			activities[i].render()
		}

		for i := range nodes {
			nodes[i].render()
		}

		rl.EndDrawing()
	}

	rl.UnloadFont(comicSans)
}
