package jgui

import sdl "jGui/sdl"

const ID_NULL ID = ^uint32(0)

const (
	WIN_FULLSCREEN   = sdl.WIN_FULLSCREEN
	WIN_OPENGL       = sdl.WIN_OPENGL
	WIN_SHOWN        = sdl.WIN_SHOWN
	WIN_HIDDEN       = sdl.WIN_HIDDEN
	WIN_BORDERLESS   = sdl.WIN_BORDERLESS
	WIN_RESIZABLE    = sdl.WIN_RESIZABLE
	WIN_DEFAULT      = (WIN_SHOWN | WIN_RESIZABLE)
	// No complete

	WINPOS_CENTERED  = sdl.WINPOS_CENTERED
	WINPOS_UNDIFINED = sdl.WINPOS_UNDIFINED
)

const (
	WIN_UPDATE_SURFACE = iota
	WIN_UPDATE_RENDER
)

const (
	RENDERER_SOFTWARE      = sdl.RENDERER_SOFTWARE
	RENDERER_TARGETTEXTURE = sdl.RENDERER_TARGETTEXTURE
	RENDERER_PRESENTVSYNC  = sdl.RENDERER_PRESENTVSYNC
	RENDERER_ACCELERATED   = sdl.RENDERER_ACCELERATED
)

type WidgetEvent = uint16
const (
	WE_IN WidgetEvent = iota
	WE_OUT
	WE_RESIZE
	WE_FOCUSIN
	WE_FOCUSOUT
	WE_CLICKL
	WE_CLICKM
	WE_CLICKR
	WE_CLICK_DOUBLEL
	WE_CLICK_DOUBLEM
	WE_CLICK_DOUBLER
)

const (
	ALIGN_CENTER = 1 << iota
	ALIGN_LEFT
	ALIGN_RIGHT
	ALIGN_TOP
	ALIGN_BOTTOM
)

const (
	REL_X = 1 << iota
	REL_Y
	REL_W
	REL_H
	REL_FLAGS = 0b1111
)
