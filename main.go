package main

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Printf("%s", "HELO")
	rl.InitWindow(275, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(24)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		taskList := rg.Button(rl.NewRectangle(10, 10, 60, 30), "Task List")
		config := rg.Button(rl.NewRectangle(75, 10, 60, 30), "Config")
		timer := rg.Button(rl.NewRectangle(140, 10, 60, 30), "Timer")
		settings := rg.Button(rl.NewRectangle(205, 10, 60, 30), "Settings")

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
		}

		go Exit()
	}
}

func Exit() {
	if rl.GetCharPressed() == rl.KeyEscape {
		rl.CloseWindow()
	}
}
