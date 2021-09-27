package jgui

import "sdl"

type Screen = sdl.Surface
type Rect = sdl.Rect

type Window struct {
	win * sdl.Window
	ren * sdl.Renderer
	_scr * sdl.Surface // scr would change to private function

	id uint32
}
