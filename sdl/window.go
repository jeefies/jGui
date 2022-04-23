package sdl

import (
    "unsafe"
)

/*
#include<SDL2/SDL.h>

int isNull(void * p) {
    if (p == NULL) {
        return 1;
    }
    return 0;
}
 */
import "C"

import "fmt"

// const for window's postion (when creating)
const (
    WINPOS_CENTERED  = int(C.SDL_WINDOWPOS_CENTERED_MASK)
    WINPOS_UNDIFINED = int(C.SDL_WINDOWPOS_UNDEFINED_MASK)
)

// const for window's flags
const (
    WIN_FULLSCREEN = uint32(C.SDL_WINDOW_FULLSCREEN)
    WIN_OPENGL     = uint32(C.SDL_WINDOW_OPENGL)
    WIN_SHOWN      = uint32(C.SDL_WINDOW_SHOWN)
    WIN_HIDDEN     = uint32(C.SDL_WINDOW_HIDDEN)
    WIN_BORDERLESS = uint32(C.SDL_WINDOW_BORDERLESS)
    WIN_RESIZABLE  = uint32(C.SDL_WINDOW_RESIZABLE)
    // No complete
)

// Window's Part
func CreateWindow(title string, width, height, x_pos, y_pos int) (*Window, error) {
    return CreateWindowWithFlags(title, width, height, x_pos, y_pos, WIN_SHOWN | WIN_RESIZABLE)
}

func CreateWindowWithFlags(title string, w, h, x, y int, flags uint32) (win *Window, err error) {
    ctitle := C.CString(title)
    defer C.free(unsafe.Pointer(ctitle))

    win = C.SDL_CreateWindow(ctitle, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
    if (C.isNull(unsafe.Pointer(win)) == 1) {
        err = NewSDLError("Could Not Create Window")
    }

    // fmt.Printf("max w, h: %d %d\n", mw, mh)

    return
}

func (win * Window) Close() {
    C.SDL_DestroyWindow(win)
}

func (win * Window) GetSurface() (*Surface) {
    return C.SDL_GetWindowSurface(win)
}

// GetMaxSize returns two integer means the Screen size (width, height)
func (win * Window) GetMaxSize() (w, h int) {
    // There's no use when using this: C.SDL_GetWindowMaximumSize(win, &cw, &ch)
    // I don't know why it always returns 0 (maybe must be in fullscreen mode?)

    current := new(C.SDL_DisplayMode)
    i := C.SDL_GetWindowDisplayIndex(win)
    err := int(C.SDL_GetCurrentDisplayMode(i, current))
    if (err != 0) {
        panic(NewSDLError("Could not get window's current display mode"))
    }
    w, h = int(current.w), int(current.h)

    return
}

// GetDesktopSize returns nil but print out all the info
// That means this is just a development function should not be used
func (win * Window) GetDesktopSize() {
    current := new(C.SDL_DisplayMode)
    for i := 0; i < int(C.SDL_GetNumVideoDisplays()); i++ {
        err := int(C.SDL_GetCurrentDisplayMode(C.int(i), current))
        if (err != 0) {
            panic("Error")
        }
        fmt.Printf("Display W, H: %d %d\n", current.w, current.h)
    }
}

func (win * Window) GetSize() (w, h int) {
    var cw, ch C.int = 0, 0
    C.SDL_GetWindowSize(win, &cw, &ch)
    w, h = int(cw), int(ch)
    return
}

func (win * Window) UpdateSurface() {
    C.SDL_UpdateWindowSurface(win)
}

// Id returns the id of the window.
// Can use the id to get the instance of the window
func (win * Window) Id() uint32 {
    return uint32(C.SDL_GetWindowID(win))
}
