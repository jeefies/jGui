package jgui

/*
Package jgui.Widgets
A file for all widgets like Button or Label
*/

import (
	"sdl"
	"sdl/ttf"
)

type Color = uint32

type Widget struct {
	childs [](*Widget)
	parent * Widget

	x, y, w, h int

	bg, fg Color
	win_id uint32
}

type Button struct {
	Widget
	text string
	text_size int
	text_color Color
	active_bg, active_fg Color
}

var default_font * ttf.Font

func init() {
    var err error
	default_font, err = ttf.OpenFont("font.ttf", 14)
	if (err != nil) {
		panic(err)
	}
}

func (btn * Button) Draw() {
	var err error
	textsur, err := default_font.RenderText(btn.text, ToSDLColor(btn.text_color))
	check(err)

	tw, th := textsur.Size()

	win := GetWinById(btn.win_id)

	sw, sh := min(tw, btn.w), min(th, btn.h)
	win.DrawRect(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, btn.fg)
	win.DrawRectBorder(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, 2, btn.bg)
	f := func(n, a, b int) int {
		return n + (a - b) / 2
	}
	r := sdl.NewRect(f(btn.x, tw, sw), f(btn.y, th, sh), 0, 0)
	win._scr.Blit(textsur, r)
}

func (btn * Button) Pack(win * Window) {
	btn.win_id = win.id
	btn.Draw()
}

func NewButton(x, y, w, h int, text string) (*Button) {
	btn := new(Button)
	btn.x, btn.y, btn.w, btn.h = x, y, w, h
	btn.text = text
	btn.bg = 0xffffff
	btn.fg = 0x000000
	btn.active_bg = 0xffffff
	btn.active_fg = 0xff0000
	btn.text_color = 0xffffff
	return btn
}

func ToRGB(c uint32) (r, g, b uint8) {
    r = uint8((c | 0xff0000) >> 4)
    g = uint8((c | 0x00ff00) >> 2)
    b = uint8((c | 0x0000ff) >> 0)
    return
}

func ToSDLColor(c uint32) sdl.Color {
    r, g, b := ToRGB(c)
    return *sdl.NewColor(r, g, b)
}
