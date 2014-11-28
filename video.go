// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	"fmt"
	"github.com/go-gl/glfw"
	"os"
)

var running bool

func main() {
	var err error

	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.Terminate()

	if err = glfw.OpenWindow(640, 320, 2040, 2040, 2040, 0, 0, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)

	glfw.SetWindowTitle("Chip 8 Emulator Window")

	glfw.SetWindowCloseCallback(onClose)
	glfw.SetKeyCallback(onKey)

	running = true

	for running {
		glfw.SwapBuffers()

		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1
	}
}

func onKey(key, state int) {
	if key == glfw.KeyEsc {
		running = state == 0
	}
}

func onClose() int {
	fmt.Println("closed")
	return 1
}
