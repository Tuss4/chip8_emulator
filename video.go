// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	"github.com/tuss4/chip8_emulator/chip_8"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Video struct {
	width, height int
	title         string
	the_screen    *sdl.Window
	the_renderer  *sdl.Renderer
}

func (v *Video) Initialize(sig chan chip_8.Signal) {
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
	v.the_renderer.Clear()
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	v.the_renderer.Present()
	defer v.the_renderer.Destroy()
	for {
		msg := <-sig
		switch {
		case msg.Msg == "draw":
			v.Draw(msg.Xcoord, msg.Ycoord)
			sdl.Delay(3000)
		case msg.Msg == "clear":
			v.Clear()
		}
	}
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

func (v *Video) Draw(x, y uint8) {
	v.the_renderer.Clear()
	rect := sdl.Rect{int32(x), int32(y), 1, 1}
	v.the_renderer.DrawRect(&rect)
	v.the_renderer.SetDrawColor(255, 255, 255, 0)
	v.the_renderer.FillRect(&rect)
	v.the_renderer.Present()
}

func (v *Video) Clear() {
	v.the_renderer.Clear()
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	v.the_renderer.Present()
}
