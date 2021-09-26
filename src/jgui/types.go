package jgui

import "sdl"

type Screen = sdl.Surface

type Window struct {
	win * sdl.Window
	scr * sdl.Surface
	ren * sdl.Renderer

	id uint32
}
