/*
Package jGui/jgui implements a higher level api for SDL framework

Author: Jeefy
Email: jeefy163@163.com (in China) or jeefyol@outlook.com (Not much used)

There's no direct use of sdl by `C` module but import jGui/sdl, which defines simple apis for sdl.
In this module, main structures are:
	Window: The pointer to a window's instance to control the events of it
	BaseWidget: Defines basic properties and simplest methods of a widget (Can not be directly used)
	Rect: Can also be called Area, which mainly represents an area or a relative area
and interfaces are:
	Widget
type alias:
	WidgetEvent: It's an alias of `uint16` which is the type of widget events
	ID: It's an alias of `uint32` which is the type of ALL ID field (ID_NULL is ^uint32(0))
*/
package jgui

import (
	"time"

	"jGui/sdl"
)

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
		win.SendEventALL(WE_INIT)
		logger.Printf("win %d at %p", win.id, win)
	}

	var eTime MouseClickEvent
	var kPress = NewKeyPressEvent()

	for {
		for !e.Poll() {
			time.Sleep(5e6) // 1e9 is a second, that mean sleep 1/20s
		}

		win := GetWindowById(e.WinId())
		if win == nil {
			if e.Type() == QUIT {
				goto MainEnd
			}
			continue
		}

		switch e.Type() {
		case MOUSE_MOTION:
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
		case WINDOWS_EVENT:
			switch e.WinEvent() {
			case sdl.WINDOW_CLOSE:
				win.Close()
			case sdl.WINDOW_RESIZED:
				fallthrough
			case sdl.WINDOW_SIZE_CHANGED:
				win.SendEventALL(WE_RESIZE)
				win.Update()
			}

		case MOUSE_DOWN:
			eTime.Down()
			eTime.lastWinId = e.WinId()
			eTime.lastButton = e.GetButton()

		case MOUSE_UP:
			if eTime.lastWinId != e.WinId() {
				break
			}

			if eTime.lastButton != e.GetButton() {
				break
			}
			
			ct := eTime.Up()

			if ct != 0 { // Click
				win.SendEvent(win.current_child, WE_CLICK)

				if win.focus_child != win.current_child {
					win.SendEvent(win.focus_child, WE_FOCUSOUT)
					win.focus_child = win.current_child
					win.SendEvent(win.current_child, WE_FOCUSIN)
				} else if ct > 0 {
					win.SendEvent(win.current_child, WE_CLICK_DOUBLE)
				}
			}

		case KEYDOWN:
			kPress.Down(e.WinId(), e.Key())
		case KEYUP:
			kPress.Up(e.WinId(), e.Key())
			if win.focus_child != ID_NULL {
				win.SendEvent(win.focus_child, WE_KEY)
			} else {
				win.SendEvent(win.current_child, WE_KEY)
			}
		case TEXT_EDITING:
			if (win.focus_child != ID_NULL) {
				win.SendEvent(win.focus_child, WE_TEXT_EDITING)
			}
		case TEXT_INPUT:
			win.SendEvent(win.focus_child, WE_TEXT_INPUT)
		} // swicth match
		win.Show()

	} // For loop

	MainEnd:
	return
}

func Quit() {
	sdl.Quit()
	return
}
