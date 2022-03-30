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

var (
	work       = 25
	shortBreak = 5
	longBreak  = 15
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
			ResetTimer(&state)
		case startPause:
			StartPause(&state)
			fmt.Println("start\\pause")
		case skip:
			fmt.Println("skip")
		}

		go Exit()
	}
}

func CheckStates(config Config, state *State) { //player_state *int, pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch (*state).Player {
	case PAUSED:
	case PLAYING:
		(*state).Duration = (*state).Duration - time.Second/24
		if (*state).Duration <= 0 {
			(*state).Player = PAUSED
			CheckPomodoroStates(config, state) //pomodoro_state, duration, color)
		}
	}
}

func CheckPomodoroStates(config Config, state *State) { //pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch (*state).Pomodoro {
	case WORK:
		config.WorkCount += 1
		if config.WorkCount%4 == 0 {
			CPD(state, LONG_REST, rl.Yellow, time.Duration(config.Long_rest)*time.Minute)
		} else {
			CPD(state, REST, rl.Green, time.Duration(config.Rest)*time.Minute)
		}
	case REST:
		fallthrough
	case LONG_REST:
		CPD(state, WORK, rl.Black, time.Duration(config.Work)*time.Minute)
	}
}

func CPD(state *State, pomodoro_state int, color rl.Color, duration time.Duration) {
	(*state).Pomodoro = pomodoro_state
	(*state).Color = color
	(*state).Duration = duration
}

func StartPause(state *State) {
	if (*state).Player == PAUSED {
		(*state).Player = PLAYING
	} else {
		(*state).Player = PAUSED
	}
}

type State struct {
	Player   int
	Pomodoro int
	Color    rl.Color
	Duration time.Duration
}

// Reset timer to default config duration.
// NEED TO DO. BAD IMPLEMENTATION
func ResetTimer(state *State) {
	(*state).Duration = time.Duration(time.Minute * time.Duration(work))
}

func Exit() {
	if rl.GetCharPressed() == rl.KeyEscape {
		rl.CloseWindow()
	}
}
