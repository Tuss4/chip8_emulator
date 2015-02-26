// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	"fmt"
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

type Row struct {
	pixels []sdl.Rect
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
			v.Draw(msg.Xcoord, msg.Ycoord, msg.Bytes)
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

func (v *Video) Draw(x, y uint8, sprite []uint8) {
	fmt.Println(x, y)
	v.the_renderer.Clear()
	height := len(sprite)
	// rects := make([]Row, height)
	row := make([]sdl.Rect, 8)
	for yline := uint8(0); yline < uint8(height); yline++ {
		for xline := uint8(0); xline < 8; xline++ {
			// fmt.Println(xline, yline)
			rect := sdl.Rect{int32(x + xline), int32(y + yline), 1, 1}
			row[xline] = rect
		}
	}
	fmt.Println(row)
}

func (v *Video) Clear() {
	v.the_renderer.Clear()
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	v.the_renderer.Present()
}
