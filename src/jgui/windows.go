package jgui

import "sdl"

var winlist = make([](* Window), 0)
// 0 for surface, 1 for renderer
const (
	UPDATE_BY_SURFACE = 0
	UPDATE_BY_RENDERER = 1
)
var updateWindowMethod int = 0

func CreateWindow(title string, w, h int, flags uint32) (win * Window) {
	var err error
	win = new(Window)
	win.id = uint32(len(winlist))

	winlist = append(winlist, win)

	win.win, err = sdl.CreateWindowWithFlags(title, w, h, sdl.WINPOS_CENTERED, sdl.WINPOS_CENTERED, flags | sdl.WIN_RESIZABLE)
	check(err)

	win.ren, err = sdl.CreateRenderer(win.win, sdl.RENDERER_PRESENTVSYNC)
	check(err)

	win.scr = win.win.GetSurface()

	return
}

func (win * Window) Size() (int, int) {
	return win.win.GetSize()
}

func (win * Window) Close() {
	win.ren.Close()
	win.win.Close()
}

func (win * Window) GetScreen() (* Screen) {
	return win.scr
}

func (win * Window) GetId() uint32 {
	return win.id
}

func (win * Window) UpdateSurface() {
	win.win.UpdateSurface()
}

func (win * Window) Update() {
	var err error

	text, err := win.ren.CreateTextureFromSurface(win.scr)
	check(err)

	win.ren.Clear()
	err = win.ren.FillTexture(text, nil)
	check(err)
}

func UpdateMethod(flag int) error {
	if 0 > flag || flag > 1 {
		return sdl.NewSDLError("No such Update Method")
	}
	updateWindowMethod = flag

	return nil
}

func updateWindows() {
	if updateWindowMethod == 0 {
		for _, w := range winlist {
			w.UpdateSurface()
		}
	} else if updateWindowMethod == 1 {
		for _, w := range winlist {
			w.Update()
		}
	}
	print("Update +1  ")
}

func updateWindows_timerfunc(interval uint32, param interface {}) uint32 {
	updateWindows()
	return interval
}
