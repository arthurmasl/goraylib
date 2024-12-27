package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// lib, err := plugin.Open("resources/library.so")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// f, err := lib.Lookup("F")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// f.(func())()

	rl.SetConfigFlags(rl.FlagWindowUnfocused)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1280, 1412, "goraylib")
	defer rl.CloseWindow()

	rl.SetWindowPosition(rl.GetMonitorWidth(0), 0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
