// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	"github.com/banthar/Go-SDL/sdl"
	"log"
)

type Video struct {
	width, height int
	title         string
}

func (v *Video) Initialize() {
	sdl.Init(sdl.INIT_VIDEO)

	defer sdl.Quit()

	window := sdl.SetVideoMode(v.width, v.height, 32, sdl.OPENGL)
	if window == nil {
		log.Fatal(sdl.GetError())
	}

	sdl.WM_SetCaption(v.title, "")
	sdl.Delay(3000)
	return

}
