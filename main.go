package main

import (
	"fmt"
	"math"
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PAUSED  = iota + 1
	PLAYING = iota + 1
)

var (
	work       = 25
	shortBreak = 5
	longBreak  = 15
)

func main() {

	state := PLAYING
	duration := time.Duration(time.Minute * time.Duration(work))

	rl.InitWindow(275, 450, "[raylib] Pomodoro")
	rl.SetTargetFPS(24)
	// rl.SetWindowState(rl.FlagWindowUndecorated)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawLine(0, 50, 275, 50, rl.Black)
		// Top Buttons
		taskList := rg.Button(rl.NewRectangle(10, 10, 60, 30), "Task List")
		config := rg.Button(rl.NewRectangle(75, 10, 60, 30), "Config")
		timer := rg.Button(rl.NewRectangle(140, 10, 60, 30), "Timer")
		settings := rg.Button(rl.NewRectangle(205, 10, 60, 30), "Settings")

		// Middle elements
		rl.DrawCircleSector(rl.NewVector2(138, 200), 100, 180, 270, 5, rl.Red)
		rl.DrawText(fmt.Sprintf("%02.0f:%0.2d", math.Floor(duration.Minutes()), int(duration.Seconds())%60), 86, 150, 42, rl.Black)
		//DEBUG LINES
		//rl.DrawLine(138, 0, 138, 450, rl.Black)

		// Bottom Buttons
		reset := rg.Button(rl.NewRectangle(50, 375, 45, 45), "Reset")
		startPause := rg.Button(rl.NewRectangle(100, 360, 75, 75), "Start\nPause")
		skip := rg.Button(rl.NewRectangle(180, 375, 45, 45), "Skip")

		switch state {
		case PAUSED:
		case PLAYING:
			duration = duration - time.Second/24
			if duration == 0 {
				state = PAUSED
			}
		}

		rl.EndDrawing()

		switch {
		case taskList:
			fmt.Println("taskList")
		case config:
			fmt.Println("config")
		case timer:
			fmt.Println("timer")
		case settings:
			fmt.Println("settings")
		case reset:
			state = PAUSED
			ResetTimer(&duration)
		case startPause:
			StartPause(&state)
			fmt.Println("start\\pause")
		case skip:
			fmt.Println("skip")
		}

		go Exit()
	}
}

func StartPause(state *int) {
	if *state == PAUSED {
		*state = PLAYING
	} else {
		*state = PAUSED
	}
}

func ResetTimer(duration *time.Duration) {
	*duration = time.Duration(time.Minute * time.Duration(work))
}

func Exit() {
	if rl.GetCharPressed() == rl.KeyEscape {
		rl.CloseWindow()
	}
}
