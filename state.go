package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	Player   int
	Pomodoro int
	Color    rl.Color
	Duration time.Duration
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

// Reset timer to default config duration.
// NEED TO DO. BAD IMPLEMENTATION
func ResetTimer(state *State) {
	(*state).Duration = time.Duration(time.Minute * time.Duration(work))
}
