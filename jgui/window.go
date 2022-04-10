/*
 * Define window actions here
 */
package jgui

import (
	"strconv" 

	"jGui/sdl"
)

var wins map[ID](*Window)

func init() {
	wins = make(map[ID] (*Window))
}

func CreateWindow(title string, w, h int, flags ID) (jw *Window) {
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
	jw.current_child = ID_NULL
	jw.focus_child = ID_NULL
	jw.childs = make([]Widget, 0, 10)
	jw.areas = make([](*Rect), 0, 10)

	jw.BackgroundColor = BLACK

	wins[jw.id] = jw

	return
}

func (w *Window) GetOriginScreen() (*Screen) {
	return w._scr
}

func (w *Window) SetUpdateMode(flag uint8) error {
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

func (w *Window) GetUpdateMode() uint8 {
	return w.update_mode
}

func (w *Window) CopyToOrigin() {
	s := w.GetScreen()

	ww, wh := w.win.GetSize()
	r := sdl.NewRect(0, 0, ww, wh)
	s.BlitPart(w._scr, r, r)
}

func (w *Window) Size() (int, int) {
	return w.win.GetSize()
}

func (w *Window) Width() int {
	width, _ := w.Size()
	return width
}

func (w *Window) Height() int {
	_, h := w.Size()
	return h
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
    text, err := w.ren.CreateTextureFromSurface(w.GetScreen())
    check(err)

    w.ren.Clear()
    err = w.ren.FillTexture(text, nil)
    check(err)

}

func (w *Window) Close() {
	logger.Printf("Window %d closed", w.id)
	w.ren.Close()
	w.win.Close()
	w._scr.Close()

	delete(wins, w.id)
}

func (w *Window) Clear() {
	w.GetOriginScreen().Fill(nil, w.BackgroundColor.Map(w.GetOriginScreen()))
}

func (w Window) GetScreen() *sdl.Surface {
	return w.win.GetSurface()
}

func (win *Window) Register(wg Widget, area *Rect) ID {
	win.childs = append(win.childs, wg)
	win.areas = append(win.areas, area)

	// No 0 id
	return ID(len(win.childs) - 1)
}

func (w *Window) Update() {
	scr := w.GetOriginScreen()
	scr.ClearWith(w.BackgroundColor)
	for i, wg := range w.childs {
		wg.Draw(scr, w.areas[i].MapVH(w.Width(), w.Height()))
	}
}

func (w *Window) UpdateWidget(id ID) error {
	if int(id) >= len(w.childs) {
		return sdl.NewSDLError("Not Valid Id: " + strconv.Itoa(int(id)))
	}
	w.childs[id].Draw(w.GetOriginScreen(), w.areas[id].MapVH(w.Width(), w.Height()))

	return nil
}

func (w *Window) GetWidget(id ID) (wg Widget, err error) {
	if int(id) >= len(w.childs) {
		err = sdl.NewSDLError("Not Valid Id: " + strconv.Itoa(int(id)))
		return
	}
	wg = w.childs[id]
	return
}

func (w *Window) GetWidgetArea(id ID) (area *Rect, err error) {
	if int(id) >= len(w.childs) {
		err = sdl.NewSDLError("Not Valid Id: " + strconv.Itoa(int(id)))
		return
	}
	area = w.areas[id].MapVH(w.Width(), w.Height())
	return
}

func (w *Window) SendEvent(id ID, we WidgetEvent) {
	// if id == ID_NULL { return } this can be included in following line
	if int(id) >= len(w.childs) {
		return
	}
	logger.Printf("Win %d Widget %d Call Event %d", w.id, id, we)
	w.childs[id].Call(we)
}

func (w *Window) SendEventALL(we WidgetEvent) {
	logger.Printf("Win %d All widget Call Event", w.id)

	for i := range w.childs {
		w.childs[i].Call(we)
	}
}

func GetWindowById(id ID) (w *Window) {
	var ok bool
	if w, ok = wins[id]; ok {
		return
	}
	w = nil
	return
}

func (w *Window) Move(id ID, area *Rect) {
	if int(id) >= len(w.childs) {
		return
	}

	wg, _ := w.GetWidget(id)
	wg.Clear()
	w.areas[id] = area
	wg.Draw(w.GetOriginScreen(), area.MapVH(w.Width(), w.Height()))
}

func (w *Window) Focus(id ID) {
	if (w.focus_child != ID_NULL) {
		w.SendEvent(w.focus_child, WE_FOCUSOUT)
	}

	w.focus_child = id
	w.SendEvent(w.focus_child, WE_FOCUSIN)
}

func (w *Window) FocusOut() {
	w.Focus(ID_NULL)
}
