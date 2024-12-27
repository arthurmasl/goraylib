package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var pos = rl.Vector2{}

func main() {
	rl.SetConfigFlags(rl.FlagWindowUnfocused)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1280, 1412, "goraylib")
	defer rl.CloseWindow()

	rl.SetWindowPosition(rl.GetMonitorWidth(0), 0)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}

func update() {
	pos = rl.GetMousePosition()
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangle(int32(pos.X), int32(pos.Y), 100, 100, rl.Red)

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}
