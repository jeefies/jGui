package jgui

import "sdl"

func init() {
	Init()
}

func Init() {
	sdl.Init(sdl.INIT_DEFAULT)
}

func Quit() {
	sdl.Quit()
}

func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func Delay(ms uint32) {
	sdl.Delay(ms)
}

func Mainloop() {
	_update_func_id := new(sdl.TimerID)
	go sdl.AddTimer(1000/24, updateWindows_timerfunc, nil, _update_func_id)
	defer sdl.RemoveTimer(_update_func_id)

	e := sdl.NewEvent()
	for {
		for e.Poll() {
			switch e.Type() {
			case sdl.QUIT:
				return
			case sdl.KEYDOWN:
				switch e.Key() {
				case 'w':
					winlist[0].DrawRectCentered(30, 30, 5, 0xff)
				}
			case sdl.WINDOWS_EVENT:
				for _, w := range winlist {
					if w.sid == e.WinId() {
						w.handleWindowEvent(e.WinEvent())
						break
					}
				}
			}
		}
	}
}
