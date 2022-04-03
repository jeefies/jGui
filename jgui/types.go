package jgui

import (
    "jGui/sdl"
)

type Screen = sdl.Surface
type Point = sdl.Point
type Rect struct {
	x, y, w, h int
}

type Window struct {
	win * sdl.Window
	ren * sdl.Renderer
	_scr * sdl.Surface // scr would change to private function

	id uint32
}

func (r Rect) ToSDL() *sdl.Rect {
	return sdl.NewRect(r.x, r.y, r.w, r.h)
}
