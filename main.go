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

	player_state := PAUSED
	pomodoro_state := WORK
	duration := time.Duration(time.Minute * time.Duration(config.Work))

	color := rl.Black

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
		_hours := math.Floor(duration.Minutes())
		_minutes := int(duration.Seconds()) % 60
		rl.DrawText(fmt.Sprintf("%02.0f:%0.2d", _hours, _minutes), 86, 150, 42, color)
		//DEBUG LINES
		//rl.DrawLine(138, 0, 138, 450, rl.Black)

		// Bottom Buttons
		reset := rg.Button(rl.NewRectangle(50, 375, 45, 45), "Reset")
		startPause := rg.Button(rl.NewRectangle(100, 360, 75, 75), "Start\nPause")
		skip := rg.Button(rl.NewRectangle(180, 375, 45, 45), "Skip")

		rl.EndDrawing()

		CheckStates(config, &player_state, &pomodoro_state, &duration, &color)

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
			player_state = PAUSED
			ResetTimer(&duration)
		case startPause:
			StartPause(&player_state)
			fmt.Println("start\\pause")
		case skip:
			fmt.Println("skip")
		}

		go Exit()
	}
}

func CheckStates(config Config, player_state *int, pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch *player_state {
	case PAUSED:
	case PLAYING:
		*duration = *duration - time.Second/24
		if *duration <= 0 {
			*player_state = PAUSED
			CheckPomodoroStates(config, pomodoro_state, duration, color)
		}
	}
}

func CheckPomodoroStates(config Config, pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch *pomodoro_state {
	case WORK:
		config.WorkCount += 1
		if config.WorkCount%4 == 0 {
			CPD(pomodoro_state, LONG_REST,
				color, rl.Yellow,
				duration, time.Duration(config.Long_rest)*time.Minute)
		} else {
			CPD(pomodoro_state, REST,
				color, rl.Green,
				duration, time.Duration(config.Rest)*time.Minute)
		}
	case REST:
		fallthrough
	case LONG_REST:
		CPD(pomodoro_state, WORK,
			color, rl.Black,
			duration, time.Duration(config.Work)*time.Minute)
	}
}

func CPD(pomodoro_state *int, p_s int, color *rl.Color, c rl.Color, duration *time.Duration, d time.Duration) {
	*pomodoro_state = p_s
	*color = c
	*duration = d
}

func StartPause(state *int) {
	if *state == PAUSED {
		*state = PLAYING
	} else {
		*state = PAUSED
	}
}

// Reset timer to default config duration.
// NEED TO DO. BAD IMPLEMENTATION
func ResetTimer(duration *time.Duration) {
	*duration = time.Duration(time.Minute * time.Duration(work))
}

func Exit() {
	if rl.GetCharPressed() == rl.KeyEscape {
		rl.CloseWindow()
	}
}
