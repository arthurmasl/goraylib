package main

import (
	"fmt"
	"os"
	"plugin"
	"time"
)

const resourcesPath = "examples/hot/resources"

var (
	lastMod  time.Time
	updateFn func() bool
	initFn   func()
	lib      *plugin.Plugin
	libPath  string
)

func main() {
	time.Sleep(time.Second)
	initWindow()

	for {
		exit := updateHot()
		if exit {
			break
		}
	}
}

func loadLib() bool {
	// check mod time
	info, err := os.Stat(resourcesPath)
	if err != nil {
		panic(err)
	}
	currMod := info.ModTime()
	if time.Time.Equal(lastMod, currMod) {
		return false
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
		return false
	}
	libPath = fmt.Sprintf("%v/%v", resourcesPath, files[0])
	fmt.Println("found", libPath)

	lib, err = plugin.Open(libPath)
	os.Remove(libPath)
	return err == nil
}

func initWindow() {
	loadLib()
	if lib == nil {
		panic("Library not loaded")
	}

	init, err := lib.Lookup("Init")
	if err != nil {
		panic(err)
	}

	init.(func())()
	fmt.Println("Window created")
}

func updateHot() bool {
	if updateFn != nil {
		exit := updateFn()
		if exit {
			return true
		}
	}

	isModified := loadLib()
	if !isModified || lib == nil {
		return false
	}

	// hot reload
	update, err := lib.Lookup("Update")
	if err != nil {
		panic(err)
	}

	updateFn = update.(func() bool)
	exit := updateFn()
	return exit
}
