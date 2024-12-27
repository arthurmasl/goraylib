package main

import "fmt"

// go build -buildmode=plugin -o resources/library.so cmd/library/main.go
func Update() {
	fmt.Println("hello")
}
