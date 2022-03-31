package main

import (
	"fmt"
	"math"
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PAUSED    = iota + 1
	PLAYING   = iota + 1
	WORK      = iota + 1
	REST      = iota + 1
	LONG_REST = iota + 1
)

func main() {

	Init()
	config := Load()

	// player_state := PAUSED
	// pomodoro_state := WORK
	duration := time.Duration(time.Minute * time.Duration(config.Work))

	state := State{
		Player:   PAUSED,
		Pomodoro: WORK,
		Color:    rl.Black,
		Duration: duration,
	}

	// color := rl.Black

	rl.InitWindow(275, 450, "[raylib] Pomodoro")
	rl.SetTargetFPS(24)
	// rl.SetWindowState(rl.FlagWindowUndecorated)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Top Buttons
		taskListBtn := rg.Button(rl.NewRectangle(10, 10, 60, 30), "Task List")
		configBtn := rg.Button(rl.NewRectangle(75, 10, 60, 30), "Config")
		timerBrn := rg.Button(rl.NewRectangle(140, 10, 60, 30), "Timer")
		settingsBrn := rg.Button(rl.NewRectangle(205, 10, 60, 30), "Settings")

		rl.DrawLine(0, 50, 275, 50, rl.Black)
		// Middle elements
		// rl.DrawCircleSector(rl.NewVector2(138, 200), 100, 180, 270, 5, rl.Red)
		_hours := math.Floor(state.Duration.Minutes())
		_minutes := int(state.Duration.Seconds()) % 60
		rl.DrawText(fmt.Sprintf("%02.0f:%0.2d", _hours, _minutes), 86, 150, 42, state.Color)
		//DEBUG LINES
		//rl.DrawLine(138, 0, 138, 450, rl.Black)

		// Bottom Buttons
		reset := rg.Button(rl.NewRectangle(50, 375, 45, 45), "Reset")
		startPause := rg.Button(rl.NewRectangle(100, 360, 75, 75), "Start\nPause")
		skip := rg.Button(rl.NewRectangle(180, 375, 45, 45), "Skip")

		rl.EndDrawing()

		// Should replace with state var
		CheckStates(config, &state) //&player_state, &pomodoro_state, &duration, &color)

		switch {
		case taskListBtn:
			fmt.Println("taskList")
		case configBtn:
			fmt.Println("config")
		case timerBrn:
			fmt.Println("timer")
		case settingsBrn:
			fmt.Println("settings")
		case reset:
			state.Player = PAUSED
			ResetTimer(&state, config)
		case startPause:
			StartPause(&state)
			fmt.Println("start\\pause")
		case skip:
			fmt.Println("skip")
		}

		go Exit()
	}
}

func Exit() {
	if rl.GetCharPressed() == rl.KeyEscape {
		rl.CloseWindow()
	}
}
