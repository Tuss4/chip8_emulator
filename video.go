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

	window := sdl.SetVideoMode(v.width, v.height, 32, sdl.SWSURFACE)
	if window == nil {
		log.Fatal(sdl.GetError())
	}

	image := sdl.Load("spidey.jpg")

	window.Blit(nil, image, nil)

	window.Flip()

	sdl.WM_SetCaption(v.title, "")

	rect := sdl.Rect{40, 40, 200, 200}

	window.FillRect(&rect, sdl.MapRGB(window.Format, 0xff, 0xff, 0xff))

	window.Flip()

	sdl.Delay(1000)

	image.Free()

	return

}
