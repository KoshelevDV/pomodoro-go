package main

import (
	"fmt"
	"math"
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Circle struct {
	startAngle float32
	endAngle   float32
	unit       float32
}

func DrawTimer(state *State, config *Config, circle *Circle) {
	// Middle elements
	fmt.Println(circle.endAngle, circle.unit)
	rl.DrawRing(rl.NewVector2(134, 170), 70, 74, circle.startAngle, circle.endAngle, 100, rl.Blue)
	if state.Player == PLAYING {
		circle.endAngle += circle.unit
	}

	_hours := math.Floor(state.Duration.Minutes())
	_minutes := int(state.Duration.Seconds()) % 60
	rl.DrawText(fmt.Sprintf("%02.0f:%0.2d", _hours, _minutes), 86, 150, 42, state.Color)
	// Bottom Buttons
	reset = rg.Button(rl.NewRectangle(50, 375, 45, 45), "Reset")
	startPause = rg.Button(rl.NewRectangle(100, 360, 75, 75), "Start\nPause")
	skip = rg.Button(rl.NewRectangle(180, 375, 45, 45), "Skip")
}

func DrawConfig(state *State, config *Config, circle *Circle) {
	rl.DrawText("Work", 14, 80, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("%3d", config.Work), 185, 80, 20, rl.Black)
	rl.DrawText(" mins", 215, 80, 20, rl.Black)
	config.Work = int(rg.Slider(rl.NewRectangle(10, 100, 255, 20), float32(config.Work), 5, 120))

	rl.DrawText("Rest", 14, 130, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("%3d", config.Rest), 185, 130, 20, rl.Black)
	rl.DrawText(" mins", 215, 130, 20, rl.Black)
	config.Rest = int(rg.Slider(rl.NewRectangle(10, 150, 255, 20), float32(config.Rest), 5, 120))

	rl.DrawText("Long Rest", 14, 180, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("%3d", config.Long_rest), 185, 180, 20, rl.Black)
	rl.DrawText(" mins", 215, 180, 20, rl.Black)
	config.Long_rest = int(rg.Slider(rl.NewRectangle(10, 200, 255, 20), float32(config.Long_rest), 5, 120))

	rl.DrawText("Sessions", 14, 230, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("%3d", config.WorkPeriod), 153, 230, 20, rl.Black)
	rl.DrawText(" rounds", 183, 230, 20, rl.Black)
	config.WorkPeriod = int(rg.Slider(rl.NewRectangle(10, 250, 255, 20), float32(config.WorkPeriod), 0, 12))

	if rg.Button(rl.NewRectangle(150, 400, 100, 30), "Save") {
		state.Duration = time.Duration(config.Work) * time.Minute
		state.Duration = time.Duration(config.Rest) * time.Minute
		state.Duration = time.Duration(config.Long_rest) * time.Minute

		circle.endAngle = 0

		circle.unit = float32(360 / state.Duration.Seconds() / 24)

	}
}
