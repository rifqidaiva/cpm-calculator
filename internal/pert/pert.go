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

	for !rl.WindowShouldClose() {
		handleInputKey()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("PERT Diagram", 190, 200, 20, rl.LightGray)
		rl.DrawCircleV(rl.NewVector2(100, 100), 50, rl.Black)

		rl.EndDrawing()
	}
}
