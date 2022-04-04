/*
 * Define window actions here
 */
package jgui

import (
	"jGui/sdl"
)

func CreateWindow(title string, w, h int, flags uint32) (jw *Window) {
	var err error
	jw = new(Window)
	win, err := sdl.CreateWindowWithFlags(title, w, h, WINPOS_CENTERED, WINPOS_CENTERED, flags)
	check(err)
	jw.win = win
	ren, err := sdl.CreateRenderer(win, RENDERER_PRESENTVSYNC)
	check(err)
	jw.ren = ren

	maxw, minh := win.GetMaxSize()
	_scr, err := sdl.CreateSurface(maxw, minh)
	check(err)
	jw._scr = _scr

	jw.id = win.Id()
	jw.update_mode = WIN_UPDATE_SURFACE
	jw.childs = make([]Widget, 0, 10)
	jw.areas = make([](*Rect), 0, 10)

	jw.bgColor = BLACK

	return
}

func (w *Window) scr() (*Screen) {
	return w.win.GetSurface()
}

func (w *Window) ChangeUpdateMode(flag uint8) error {
	switch flag {
	case WIN_UPDATE_SURFACE:
		w.update_mode = WIN_UPDATE_SURFACE
	case WIN_UPDATE_RENDER:
		w.update_mode = WIN_UPDATE_RENDER
	default:
		return sdl.NewSDLError("Could not determine update mode")
	}
	return nil
}

func (w *Window) CopyToOrigin() {
	s := w.scr()

	ww, wh := w.win.GetSize()
	r := sdl.NewRect(0, 0, ww, wh)
	s.BlitPart(w._scr, r, r)
}

func (w *Window) Size() (int, int) {
	return w.win.GetSize()
}

func (w *Window) Show() {
	if (w.update_mode == WIN_UPDATE_SURFACE) {
		w.CopyToOrigin()
		w.win.UpdateSurface()
	} else if (w.update_mode == WIN_UPDATE_RENDER) {
		w.RenderScreen()
		w.ren.Present()
	}
}

func (w *Window) RenderScreen() {
    var err error

	w.CopyToOrigin()
    text, err := w.ren.CreateTextureFromSurface(w.scr())
    check(err)

    w.ren.Clear()
    err = w.ren.FillTexture(text, nil)
    check(err)

}

func (w *Window) Close() {
	w.ren.Close()
	w.win.Close()
	w._scr.Close()
}

func (w Window) GetScreen() *sdl.Surface {
	return w._scr
}

func (win *Window) Register(wg Widget, x, y, w, h int) int {
	area := NewRect(x, y, w, h)

	win.childs = append(win.childs, wg)
	win.areas = append(win.areas, area)

	return len(win.childs)
}

func (w *Window) Update() {
	scr := w.GetScreen()
	for i, wg := range w.childs {
		wg.Draw(scr, w.areas[i].Copy())
	}
}

func (w *Window) UpdateWidget(id int) error {
	if (0 < id) || (id >= len(w.childs)) {
		return sdl.NewSDLError("Not valid Id")
	}
	w.childs[id].Draw(w.GetScreen(), w.areas[id].Copy())

	return nil
}
