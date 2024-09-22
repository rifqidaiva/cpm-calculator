package pert

import rl "github.com/gen2brain/raylib-go/raylib"

func handleInputKey() {
	if rl.IsKeyPressed(rl.KeyF) {
		if rl.IsWindowFullscreen() {
			rl.ToggleFullscreen()
			rl.SetWindowSize(windowWidth, windowHeight)
		} else {
			monitor := rl.GetCurrentMonitor()
			rl.SetWindowSize(rl.GetMonitorWidth(monitor), rl.GetMonitorHeight(monitor))
			rl.ToggleFullscreen()
		}
	}
}
