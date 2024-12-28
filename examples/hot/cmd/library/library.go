package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Init() {
	rl.SetConfigFlags(rl.FlagWindowUnfocused)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1280, 1412, "goraylib")

	rl.SetWindowPosition(rl.GetMonitorWidth(0), 0)
	rl.SetTargetFPS(60)
}

func Update() bool {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangle(100, 100, 100, 100, rl.Red)

	rl.DrawFPS(10, 10)
	rl.EndDrawing()

	exit := rl.WindowShouldClose()
	if exit {
		rl.CloseWindow()
	}
	return exit
}
