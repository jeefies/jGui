package jgui

import "sdl"

type Screen = sdl.Surface
type Rect = sdl.Rect

type Window struct {
	win * sdl.Window
	ren * sdl.Renderer
	_scr * sdl.Surface // scr would change to private function

	childs []Widgets
    current_widget Widgets

	id uint32
	sid uint32
}

type JguiError struct {
	msg string
}

func (je *JguiError) Error() string {
	return je.msg
}

func NewError(msg string) (*JguiError) {
	return &JguiError{msg}
}
