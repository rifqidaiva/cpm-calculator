package pert

import rl "github.com/gen2brain/raylib-go/raylib"

type activity struct {
	startNumber  int
	finishNumber int
	letter       string
	duration     int
}

func (a activity) render() {
	startLine := rl.Vector2{}
	endLIne := rl.Vector2{}

	for i := range nodes {
		if nodes[i].number == a.startNumber {
			startLine = nodes[i].position
		}

		if nodes[i].number == a.finishNumber {
			endLIne = nodes[i].position
		}
	}

	rl.DrawLineV(startLine, endLIne, rl.Black)
}

func newActivity(startNumber int, finishNumber int, letter string, duration int) activity {
	return activity{
		startNumber:  startNumber,
		finishNumber: finishNumber,
		letter:       letter,
		duration:     duration,
	}
}
