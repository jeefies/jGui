package jgui

import (
    "jGui/sdl"
)

type Screen = sdl.Surface
type Color = sdl.Color
type ID = uint32

type Point struct {
	X, Y int
}

type Rect struct {
	x, y, w, h float64
	isRel uint8
}

type Widget interface {
	Size() (w, h int)
	Id() ID
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

	id ID
	update_mode uint8

	BackgroundColor Color

	current_child ID  // Id
	childs []Widget
	areas [](*Rect)
}


func (r Rect) ToSDL() *sdl.Rect {
	return sdl.NewRect(int(r.x), int(r.y), int(r.w), int(r.h))
}

func (r Rect) Pout() {
	logger.Printf("R x, y, w, h = %d %d %d %d", r.x, r.y, r.w, r.h)
}

func (p Point) Pout() {
	logger.Printf("P x, y = %d %d", p.X, p.Y)
}

func NewRect(x, y, w, h int) *Rect {
	return &Rect{float64(x), float64(y), float64(w), float64(h), 0}
}

func NewRelRect(x, y, w, h float64, relFlags uint8) *Rect {
	return &Rect{x, y, w, h, relFlags & REL_FLAGS}
}

func (r Rect) Copy() (*Rect) {
	return NewRelRect(r.x, r.y, r.w, r.h, 0)
}

func (r Rect) X() int {
	return int(r.x)
}

func (r Rect) Y() int {
	return int(r.y)
}

func (r Rect) W() int {
	return int(r.w)
}

func (r Rect) H() int {
	return int(r.h)
}

func (r *Rect) SetX(x int) {
	r.x = float64(x)
}
func (r *Rect) SetW(x int) {
	r.w = float64(x)
}
func (r *Rect) SetY(x int) {
	r.y = float64(x)
}
func (r *Rect) SetH(x int) {
	r.h = float64(x)
}

func (r Rect) MapVH(v_i, h_i int) (*Rect) {
	nr := r.Copy()
	v := float64(v_i)
	h := float64(h_i)

	if r.isRel & REL_X == REL_X { nr.x *= v }
	if r.isRel & REL_Y == REL_Y { nr.y *= h }
	if r.isRel & REL_W == REL_W { nr.w *= v }
	if r.isRel & REL_H == REL_H { nr.h *= h }
	return nr
}

func (p Point) IsIn(r *Rect) bool {
	x := float64(p.X)
	y := float64(p.Y)
	if x >= r.x && x <= r.x + r.w && y >= r.y && y <= r.y + r.h {
		return true
	}
	return false
}
