package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu bool

type State struct {
	Player   int
	Pomodoro int
	Color    rl.Color
	Duration time.Duration
	Menu     []Menu
}

func CheckStates(config Config, state *State, circle *Circle) { //player_state *int, pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch (*state).Player {
	case PAUSED:
	case PLAYING:
		(*state).Duration = (*state).Duration - time.Second/24
		if (*state).Duration <= 0 {
			(*state).Player = PAUSED
			CheckPomodoroStates(config, state, circle) //pomodoro_state, duration, color)
		}
	}
}

func CheckPomodoroStates(config Config, state *State, circle *Circle) { //pomodoro_state *int, duration *time.Duration, color *rl.Color) {
	switch (*state).Pomodoro {
	case WORK:
		config.WorkCount += 1
		if config.WorkCount%4 == 0 {
			CPD(state, LONG_REST, rl.Yellow, time.Duration(config.Long_rest)*time.Minute)
			circle.endAngle = 0
			circle.unit = float32(360 / state.Duration.Seconds() / 24)
		} else {
			CPD(state, REST, rl.Green, time.Duration(config.Rest)*time.Minute)
			circle.endAngle = 0
			circle.unit = float32(360 / state.Duration.Seconds() / 24)
		}
	case REST:
		fallthrough
	case LONG_REST:
		CPD(state, WORK, rl.Black, time.Duration(config.Work)*time.Minute)
		circle.endAngle = 0
		circle.unit = float32(360 / state.Duration.Seconds() / 24)
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

// Reset timer to default config duration.
// NEED TO DO. BAD IMPLEMENTATION
func ResetTimer(state *State, config Config) {
	(*state).Duration = time.Duration(time.Minute * time.Duration(config.Work))
}
