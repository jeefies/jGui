package jgui

import (
    "log"
    _ "sync/atomic"

    "sdl"
)

var FPS uint32 = 35

func init() {
	Init()
}

func Init() {
	sdl.Init(sdl.INIT_DEFAULT)
}

func Quit() {
    print("jgui Quit\n")
	sdl.Quit()
}

func check(err error) {
	if (err != nil) {
		panic(err)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Delay(ms uint32) {
	sdl.Delay(ms)
}

func Mainloop() {
    // log.Println("Mainloop start")
	_update_func_id := new(sdl.TimerID)
    for i, win := range winlist {
        for j, wg := range win.childs {
            // FIXME: Would not draw widgets for the first time
            wg.Draw(0)
            log.Printf("win %d; wg %d\n", i, j)
        }
    }
    
	go sdl.AddTimer(1000/FPS, updateWindows_timerfunc, nil, _update_func_id)
    defer sdl.RemoveTimer(_update_func_id)

    done := make(chan bool)

    go func() {
        var ev *sdl.Event
        var c int
        var avaliable bool = true
        for {
            ev, c, avaliable = PollEvent()
            if avaliable {
                // Print("Go func\n")
                go func(e *sdl.Event, i int) {
                    defer CloseEvent(i)
                    switch e.Type() {
                        case sdl.QUIT:
                            done <- true
                            return
                        case sdl.WINDOWS_EVENT:
                            for _, w := range winlist {
                                if w.sid == e.WinId() {
                                    w.handleWindowEvent(e.WinEvent())
                                    break
                                }
                            }
                        default:
                            for _, w := range winlist {
                                if w.sid == e.WinId() {
                                    w.handleEvent(e)
                                    // logger.Print("still ok\n")
                                    break
                                }
                            }
                    }
                }(ev, c)
            } else {
                CloseEvent(c)
                Delay(1000/FPS)
            }
        }
    }()

    for {
        select {
            case <-done:
                return
            default:
                // logger.Println("flushing")
                flushOut()
                // logger.Println("flushing ok")
        }
    }
}
