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
	timer := time.NewTimer(time.Second / FPS)

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

			if (win.current_child != 0) {
				area, _ := win.GetWidgetArea(win.current_child)

				if !mousep.IsIn(area) {
					move_out = false
				}
			}

			if move_out {
				win.SendEvent(win.current_child, WE_OUT)
				for idi, area := range win.areas {
					id := uint32(idi)
					if mousep.IsIn(area) {
						win.current_child = id
						win.SendEvent(id, WE_IN)
						break
					}
				} // for loop
				win.current_child = 0
			}
			// if not move_out
			

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
