// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package chip_8

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Video struct {
	width, height int
	title         string
	the_screen    *sdl.Window
	the_renderer  *sdl.Renderer
}

func (v *Video) Initialize() {
	sdl.Init(sdl.INIT_EVERYTHING)

	defer sdl.Quit()
	window, err := sdl.CreateWindow(v.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		v.width, v.height, sdl.WINDOW_SHOWN)

	v.the_screen = window

	if err != nil {
		log.Fatal(sdl.GetError())
	}
	defer v.the_screen.Destroy()

	renderer, r_err := sdl.CreateRenderer(v.the_screen, 1, sdl.RENDERER_ACCELERATED)
	v.the_renderer = renderer
	if r_err != nil {
		log.Fatal(sdl.GetError())
	}
	defer v.the_renderer.Destroy()

	v.Draw()

	sdl.Delay(3000)
	return
}

func (v *Video) SetWidthHeight(w, h int) {
	v.height = h
	v.width = w
}

func (v *Video) SetTitle(title string) {
	v.title = title
}

func (v *Video) Draw() {
	v.the_renderer.Clear()
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	rect := sdl.Rect{50, 0, 50, 50}
	v.the_renderer.DrawRect(&rect)
	v.the_renderer.SetDrawColor(255, 255, 255, 0)
	v.the_renderer.FillRect(&rect)
	v.the_renderer.Present()
}
