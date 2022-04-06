package jgui

import (
    "jGui/sdl"
)

type Screen = sdl.Surface
type Color = sdl.Color

type Point struct {
	X, Y int
}

type Rect struct {
	x, y, w, h int
}

type Widget interface {
	Size() (w, h int)
	Id() int
	// Replace(x, y int)
	Call(e WidgetEvent)
	Draw(sur *Screen, area * Rect)

	Width() int
	Height() int
	SetWidth(w int) int
	SetHeight(h int) int
}

type Window struct {
	win * sdl.Window
	ren * sdl.Renderer
	_scr * Screen // this is not from win.GetSurface

	id uint32
	update_mode uint8

	bgColor Color

	current_child uint32  // Id
	childs []Widget
	areas [](*Rect)
}

func (r Rect) ToSDL() *sdl.Rect {
	return sdl.NewRect(r.x, r.y, r.w, r.h)
}

func (r Rect) Pout() {
	logger.Printf("R x, y, w, h = %d %d %d %d", r.x, r.y, r.w, r.h)
}

func NewRect(x, y, w, h int) *Rect {
	return &Rect{x, y, w, h}
}

func (r Rect) Copy() (*Rect) {
	return NewRect(r.x, r.y, r.w, r.h)
}

func (p Point) IsIn(r *Rect) bool {
	if p.X >= r.x && p.X <= r.x + r.w && p.Y >= r.y && p.Y <= r.y + r.h {
		return true
	}
	return false
}
