package jgui

import (
	"time"

	"jGui/sdl"
)

var FPS time.Duration = 30

func MAX (x, y int) int {
	if (x > y) {
		return x
	}
	return y
}

func ABS(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

func init() {
	sdl.Init(sdl.INIT_DEFAULT)
}

func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func Mainloop() {
	e := sdl.NewEvent()
	for _, win := range wins {
		win.Update()
		win.Event = e
		logger.Printf("win %d at %p", win.id, win)
	}

	timer := time.NewTicker(time.Second / FPS)

	var eTime MouseClickEvent

	for {
		if !e.Poll() {
			time.Sleep(1e5) // 1e9 is a second, that mean sleep 1/1000s
			goto SELECT_UPDATE
		}

		switch e.Type() {
		case sdl.QUIT:
			goto MainEnd
		case sdl.MOUSE_MOTION:
			win := GetWindowById(e.WinId())
			var move_out bool = true

			mx, my := e.MousePosition()
			var mousep = Point{mx, my}

			// logger.Printf("Current Child ID %u", win.current_child)
			if (win.current_child != ID_NULL) {
				area, err := win.GetWidgetArea(win.current_child)
				if (err != nil) {
					panic(err)
				}

				if mousep.IsIn(area.MapVH(win.Size())) {
					goto MOTION_CLEAR_UP
				} else {
					win.SendEvent(win.current_child, WE_OUT)
				}
			}

			if move_out {
				for index := range win.areas {
					id := ID(index)
					if mousep.IsIn(win.areas[index].MapVH(win.Size())) {
						win.current_child = id
						win.SendEvent(id, WE_IN)
						goto MOTION_CLEAR_UP
					}
				} // for loop
			}
			// if not move_out
			win.current_child = ID_NULL

			MOTION_CLEAR_UP: 
			win.Show()
		case sdl.WINDOWS_EVENT:
			win := GetWindowById(e.WinId())
			switch e.WinEvent() {
			case sdl.WINDOW_CLOSE:
				win.Close()
			case sdl.WINDOW_RESIZED:
				fallthrough
			case sdl.WINDOW_SIZE_CHANGED:
				win.Clear()
				win.SendEventALL(WE_RESIZE)
			}

		case sdl.MOUSE_DOWN:
			// win := GetWindowById(e.WinId())
			eTime.Down()
			eTime.lastWinId = e.WinId()
			eTime.lastButton = e.GetButton()

		case sdl.MOUSE_UP:
			if eTime.lastWinId != e.WinId() {
				break
			}

			if eTime.lastButton != e.GetButton() {
				break
			}
			
			win := GetWindowById(e.WinId())
			ct := eTime.Up()

			var CLICK, DCLICK WidgetEvent
			switch e.GetButton() {
			case sdl.BUTTON_LEFT:
				CLICK = WE_CLICKL
				DCLICK = WE_CLICK_DOUBLEL
			case sdl.BUTTON_MIDDLE:
				CLICK = WE_CLICKM
				DCLICK = WE_CLICK_DOUBLEM
			case sdl.BUTTON_RIGHT:
				CLICK = WE_CLICKR
				DCLICK = WE_CLICK_DOUBLER
			}

			if ct != 0 { // Click
				win.SendEvent(win.current_child, CLICK)

				if win.focus_child != win.current_child {
					win.SendEvent(win.focus_child, WE_FOCUSOUT)
					win.focus_child = win.current_child
					win.SendEvent(win.current_child, WE_FOCUSIN)
				} else if ct > 0 {
					win.SendEvent(win.current_child, DCLICK)
				}
			}

		case sdl.KEYUP:
			win := GetWindowById(e.WinId())
			if win == nil { break }

			if win.focus_child != ID_NULL {
				win.SendEvent(win.focus_child, WE_KEY)
			} else {
				win.SendEvent(win.current_child, WE_KEY)
			}
		case sdl.TEXT_EDITING:
			win := GetWindowById(e.WinId())
			if win == nil { break }

			if (win.focus_child != ID_NULL) {
				win.SendEvent(win.focus_child, WE_TEXT_EDITING)
			}
		case sdl.TEXT_INPUT:
			win := GetWindowById(e.WinId())
			win.SendEvent(win.focus_child, WE_TEXT_INPUT)
		} // swicth match

		SELECT_UPDATE:
		// Refresh windows
		select {
		case <-timer.C:
			for _, cw := range wins {
				cw.Show()
			}
		default:
		}
	} // For loop

	MainEnd:
	return
}

func Quit() {
	sdl.Quit()
	return
}
