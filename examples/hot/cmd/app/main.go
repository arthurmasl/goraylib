package main

import (
	"fmt"
	"os"
	"plugin"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const resourcesPath = "examples/hot/resources"

var (
	lastMod  time.Time
	updateFn func()
)

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
	if updateFn != nil {
		updateFn()
	}

	// check mod time
	info, err := os.Stat(resourcesPath)
	if err != nil {
		panic(err)
	}
	currMod := info.ModTime()
	if time.Time.Equal(lastMod, currMod) {
		return
	}
	lastMod = currMod
	fmt.Println("dir modified")

	// get plugin
	dir, err := os.Open(resourcesPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	files, err := dir.Readdirnames(1)
	if err != nil {
		return
	}
	pluginPath := fmt.Sprintf("%v/%v", resourcesPath, files[0])
	fmt.Println("found", pluginPath)

	// hot reload
	lib, err := plugin.Open(pluginPath)
	if err != nil {
		return
	}

	update, err := lib.Lookup("Update")
	if err != nil {
		panic(err)
	}

	updateFn = update.(func())
	updateFn()

	// remove plugin
	fmt.Println("remove", pluginPath)
	err = os.Remove(pluginPath)
	if err != nil {
		panic(err)
	}
}
