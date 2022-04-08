package jgui

import (
	"time"

	"jGui/sdl"
)

var FPS time.Duration = 30

func MAX(x, y int) int {
	if (x < y) {
		return y
	}
	return x
}

func MIN(x, y int) int {
	if (x < y) {
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
	for _, win := range wins {
		win.Update()
	}

	e := sdl.NewEvent()
	timer := time.NewTicker(time.Second / FPS)

	for {
		if !e.Poll() {
			time.Sleep(1e5) // 1e9 is a second, that mean sleep 1/1000s
			goto SELECT_UPDATE
		}

		switch e.Type() {
		case sdl.QUIT:
			goto MainEnd
		case sdl.MOUSE_MOTION:
			var move_out bool = true
			var win = GetWindowById(e.WinId())

			mx, my := e.MousePosition()
			var mousep = Point{mx, my}

			// logger.Printf("Current Child ID %u", win.current_child)
			if (win.current_child != ID_NULL) {
				area, err := win.GetWidgetArea(win.current_child)
				if (err != nil) {
					panic(err)
				}

				if mousep.IsIn(area) {
					goto MOTION_CLEAR_UP
				} else {
					win.SendEvent(win.current_child, WE_OUT)
				}
			}

			if move_out {
				for index := range win.areas {
					id := ID(index)
					if mousep.IsIn(win.areas[index]) {
						win.current_child = id
						win.SendEvent(id, WE_IN)
						goto MOTION_CLEAR_UP
					}
				} // for loop
			}
			// if not move_out
			win.current_child = ID_NULL

			MOTION_CLEAR_UP: 
			// logger.Printf("Change to %u", win.current_child)
			

		} // swicth match

		SELECT_UPDATE:
		// Refresh windows
		select {
		case <-timer.C:
			for _, win := range wins {
				win.Show()
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
