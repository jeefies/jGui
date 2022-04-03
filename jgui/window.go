/*
 * Define window actions here
 */
package jgui

import "jGui/sdl"

func CreateWindow(title string, w, h int, flags uint32) *Window {
	var err error
	win, err := sdl.CreateWindowWithFlags(title, w, h, WINPOS_CENTERED, WINPOS_CENTERED, flags)
	check(err)
	ren, err := sdl.CreateRenderer(win, RENDERER_PRESENTVSYNC)
	check(err)

	_scr := win.GetSurface()

	return &Window{win, ren, _scr, win.Id()}
}

func (w *Window) Close() {
	w.ren.Close()
	w.win.Close()
}
