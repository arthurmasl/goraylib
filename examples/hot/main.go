package main

import (
	"fmt"
	"os"
	"plugin"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const libraryPath = "resources/library.so"

var lastMod time.Time

// 	@go build -buildmode=plugin -o resources/library.so cmd/library/main.go

func main() {
	rl.SetConfigFlags(rl.FlagWindowUnfocused)
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(1280, 1412, "goraylib")
	defer rl.CloseWindow()

	rl.SetWindowPosition(rl.GetMonitorWidth(0), 0)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		updateHot()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

func updateHot() {
	// check mod time
	fileInfo, err := os.Stat(libraryPath)
	if err != nil {
		panic(err)
	}
	currMod := fileInfo.ModTime()
	if time.Time.Equal(lastMod, currMod) {
		return
	}
	lastMod = currMod

	// hot reload
	fmt.Println("update")
	lib, err := plugin.Open(libraryPath)
	if err != nil {
		panic(err)
	}

	update, err := lib.Lookup("Update")
	if err != nil {
		panic(err)
	}

	update.(func())()
}
