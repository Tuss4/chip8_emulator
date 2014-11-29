// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	"fmt"
	"github.com/go-gl/glfw"
	"os"
)

var running bool

type Video struct {
	width, height int
	title         string
}

func (v *Video) Initialize() {
	var err error

	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.Terminate()

	if err = glfw.OpenWindow(v.width, v.height, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)

	glfw.SetWindowTitle(v.title)

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
