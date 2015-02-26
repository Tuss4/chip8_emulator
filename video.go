// Thanks in part to
// https://github.com/go-gl/examples/blob/master/glfw/simplewindow/main.go
package main

import (
	// "fmt"
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

	window, err := sdl.CreateWindow(v.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		v.width, v.height, sdl.WINDOW_SHOWN)

	v.the_screen = window

	if err != nil {
		log.Fatal(sdl.GetError())
	}
	defer v.the_screen.Destroy()

	renderer, r_err := sdl.CreateRenderer(v.the_screen, -1, sdl.RENDERER_ACCELERATED)
	v.the_renderer = renderer
	if r_err != nil {
		log.Fatal(sdl.GetError())
	}
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	v.the_renderer.Present()
	defer v.the_renderer.Destroy()
	for {
		msg := <-sig
		switch {
		case msg.Msg == "draw":
			v.Draw(msg.Xcoord, msg.Ycoord, msg.Bytes)
		case msg.Msg == "clear":
			v.Clear()
		}
	}
}

func (v *Video) SetWidthHeight(w, h int) {
	v.height = h
	v.width = w
}

func (v *Video) SetTitle(title string) {
	v.title = title
}

func (v *Video) Draw(x, y uint8, sprite []uint8) {
	height := len(sprite)
	all_points := make([]Row, height)
	for yline := uint8(0); yline < uint8(height); yline++ {
		sprite_row := sprite[yline]
		row := make([]sdl.Rect, 8)
		for xline := uint8(0); xline < 8; xline++ {
			if sprite_row&(0x80>>xline) != 0 {
				rect := sdl.Rect{int32(x + xline), int32(y + yline), 1, 1}
				row[xline] = rect
			}
		}
		all_points[yline].pixels = row
	}

	v.the_renderer.SetDrawColor(255, 255, 255, 0)
	for _, row := range all_points {
		v.the_renderer.DrawRects(row.pixels)
		v.the_renderer.FillRects(row.pixels)
		v.the_renderer.Present()
	}
}

func (v *Video) Clear() {
	v.the_renderer.Clear()
	v.the_renderer.SetDrawColor(0, 0, 0, 0)
	v.the_renderer.Present()
}
