package jgui

import (
	"sync"
)

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

	win._scr = win.win.GetSurface()

	return
}

func (win * Window) scr() (*Screen) {
	now_scr := win.win.GetSurface()
	if now_scr == win._scr {
		return now_scr
	}

	/*
	Notice that the surface pointer would be changed
	when resizing the Window.
	This is the function to make sure the old content
	would be copied to new screen.
	Depends on the smaller screen size (big -> small, small not change)
	*/

	min := func(x, y int) int {
		if x <= y {
			return x
		} else {
			return y
		}
	}


	ow, oh := win._scr.Size() // Old Screen Size
	nw, nh := now_scr.Size() // New Screen Size

	sw, sh := min(ow, nw), min(oh, nh)
	copy_size := sdl.NewRect(0, 0, sw, sh)

	now_scr.Blit(win._scr, copy_size)

	win._scr = now_scr

	return now_scr
}

// Window.Size returns two integer (int, int) means the width and the height of the window
func (win * Window) Size() (int, int) {
	return win.win.GetSize()
}

// Destroy the window
func (win * Window) Close() {
	// First should destroy the renderer
	win.ren.Close()
	win.win.Close()
}

// Returns Screen (sdl.Surface) of the window
// No much use with type Screen
func (win * Window) GetScreen() (* Screen) {
	return win.scr()
}

// Get the id of the window.
// Id might not use, but to get the window instance by id
func (win * Window) GetId() uint32 {
	return win.id
}

// Use native interface of SDL_UpdateWindowSurface to update the window
// Not recommend, use `Update` instead
func (win * Window) UpdateSurface() {
	win.win.UpdateSurface()
}

// Update the window by creating texture from the surface and posting in renderer
func (win * Window) Update() {
	var err error

	text, err := win.ren.CreateTextureFromSurface(win.scr())
	check(err)

	win.ren.Clear()
	err = win.ren.FillTexture(text, nil)
	check(err)
}

// Decide the way of updating windows is to use surface or renderer when mainloop
// flag could be one of UPDATE_BY_RENDERER and UPDATE_BY_SURFACE
func UpdateMethod(flag int) error {
	if 0 > flag || flag > 1 {
		return sdl.NewSDLError("No such Update Method")
	}
	updateWindowMethod = flag

	return nil
}

// The function to update all of the windows in one time
func updateWindows() {
	// Use different thread to update all the window
	// By sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(len(winlist))

	if updateWindowMethod == 0 {
		for _, w := range winlist {
			go func(win *Window, wg *sync.WaitGroup) {
				win.UpdateSurface()
			}(w, &wg)
		}
	} else if updateWindowMethod == 1 {
		for _, w := range winlist {
			go func(win *Window, wg *sync.WaitGroup) {
				win.Update()
			}(w, &wg)
		}
	}

	wg.Wait()
	print("Update +1  ")
}

func updateWindows_timerfunc(interval uint32, param interface {}) uint32 {
	updateWindows()
	return interval
}
