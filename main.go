package main

import (
	"fmt"
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

var (
	reset      = false
	startPause = false
	skip       = false

	topConfBtn  = false
	topTimerBtn = true
)

func main() {

	Init()
	config := Load()

	// player_state := PAUSED
	// pomodoro_state := WORK
	duration := time.Duration(time.Minute * time.Duration(config.Work))

	menu := []Menu{Menu(topConfBtn), Menu(topTimerBtn)}

	state := State{
		Player:   PAUSED,
		Pomodoro: WORK,
		Color:    rl.Black,
		Duration: duration,
		Menu:     menu,
	}

	circle := Circle{
		startAngle: 0,
		endAngle:   0,
		unit:       float32(360 / (time.Duration(config.Work) * time.Minute).Seconds() / 24),
	}
	fmt.Println(float32(360 / time.Duration(config.Work).Seconds()))
	fmt.Println((time.Duration(config.Work) * time.Minute).Seconds())
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
		timerBtn := rg.Button(rl.NewRectangle(140, 10, 60, 30), "Timer")
		settingsBtn := rg.Button(rl.NewRectangle(205, 10, 60, 30), "Settings")

		rl.DrawLine(0, 50, 275, 50, rl.Black)
		switch {
		case bool(state.Menu[0]):
			DrawConfig(&state, &config, &circle)
		case bool(state.Menu[1]):
			DrawTimer(&state, &config, &circle)
		}
		rl.EndDrawing()

		// Should replace with state var
		CheckStates(config, &state, &circle) //&player_state, &pomodoro_state, &duration, &color)

		switch {
		case taskListBtn:
			fmt.Println("taskList")
		case configBtn:
			fmt.Println("config")
			state.Menu[0] = true
			state.Menu[1] = false
		case timerBtn:
			fmt.Println("timer")
			state.Menu[0] = false
			state.Menu[1] = true
		case settingsBtn:
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
