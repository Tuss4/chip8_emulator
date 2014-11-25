package main

import (
	"github.com/go-gl/glfw"
)

func main() {
	test_chan := make(chan int)
	glfw.Init()
	glfw.OpenWindow(640, 380, 0, 0, 0, 8, 8, 8, glfw.Windowed)
	<-test_chan
}
