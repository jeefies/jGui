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


// CreateWindow returns a Window instance
// give title string, width, heigth int, flags uint32
func CreateWindow(title string, w, h int, flags uint32) (win * Window) {
    var err error
    win = new(Window)
    win.id = uint32(len(winlist))

    winlist = append(winlist, win)

    win.win, err = sdl.CreateWindowWithFlags(title, w, h, sdl.WINPOS_CENTERED, sdl.WINPOS_CENTERED, flags | sdl.WIN_RESIZABLE)
    check(err)

    win.ren, err = sdl.CreateRenderer(win.win, sdl.RENDERER_PRESENTVSYNC)
    check(err)


    mw, mh := win.win.GetMaxSize()
    logger.Printf("Max w h: %d %d\n", mw, mh)

    // win._scr = win.win.GetSurface()
    win._scr, err = sdl.CreateSurface(mw, mh)
    check(err)

    win.sid = win.win.Id()

    win.childs = make([]Widgets, 0, 5)
    win.current_widget = nil

    return
}

// GetWinById returns a Window instance by Id
func GetWinById(id uint32) (*Window) {
	if id < uint32(len(winlist)) {
		return winlist[int(id)]
	}
	panic(NewError("Could Not Get Window By these Id"))
	return nil
}

// scr returns a Screen instance for the window instance.
func (win * Window) scr() (*Screen) {
    s := win.GetScreen()
    s.Clear()

    ww, wh := win.Size()
    r := sdl.NewRect(0, 0, ww, wh)
    s.BlitPart(win._scr, r, nil)

    return s
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
    return win.win.GetSurface()
}

// Get the id of the window (in jgui).
// Id might not use, but to get the window instance by id
func (win * Window) GetId() uint32 {
    return win.id
}

// GetSDLId returns the id of the window in sdl
// It's not equal to GetId method
func (win * Window) GetSDLId() uint32 {
    return win.win.Id()
}

// handleWindowEvent handles the windows event like resizing the size
func (win * Window) handleWindowEvent(etype uint32) {
    // TODO: Window handleWindowEvent NOT COMPLETE NOW
}

// handleEvent handles all event of the window
func (win * Window) handleEvent(e * sdl.Event) {
    switch e.Type() {
        case sdl.WINDOWS_EVENT:
            win.handleWindowEvent(e.Type())
        case sdl.MOUSE_MOTION:
            x, y := e.MousePosition()
            if (win.current_widget == nil) {
                for _, w := range win.childs {
                    if w.IsIn(x, y) {
                        win.current_widget = w
                        w.Call("active")
                        logger.Printf("Active at %d %d\n", x, y)
                        break
                    }
                }
            } else {
                cw := win.current_widget
                if cw.IsIn(x, y) {
                    cw.Call("mouse move")
                } else {
                    logger.Printf("Deactive at %d %d\n", x, y)
                    cw.Call("deactive")

                    win.current_widget = nil
                    for _, w := range win.childs {
                        if w.IsIn(x, y) {
                            win.current_widget = w
                            w.Call("active")
                            break
                        }
                    }
                }
            }
        case sdl.MOUSE_DOWN:
            if (win.current_widget == nil) {
                // TODO: win on click event
            } else {
                win.current_widget.Call("mouse down")
            }
        case sdl.MOUSE_UP:
            if (win.current_widget == nil) {
                // TODO: win click event
            } else {
                win.current_widget.Call("mouse up")
                logger.Print("Call mouse up ok")
            }
        case sdl.KEYDOWN:
            k := e.Key()
            Printf("Get Key %c\n", byte(k))
            if (win.current_widget == nil) {
                // TODO: win keydown Events
            } else {
                // win.current_widget.Call()
            }
        default:
            // TODO: Other SDL Events here
    }
}

// Use native interface of SDL_UpdateWindowSurface to update the window
// Not recommend, use `Update` instead
func (win * Window) UpdateSurface() {
    win.scr()
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
                wg.Done()
            }(w, &wg)
        }
    } else if updateWindowMethod == 1 {
        for _, w := range winlist {
            go func(win *Window, wg *sync.WaitGroup) {
                win.Update()
                wg.Done()
            }(w, &wg)
        }
    }

    wg.Wait()
}

func updateWindows_timerfunc(interval uint32, param interface {}) uint32 {
    updateWindows()
    return interval
}
