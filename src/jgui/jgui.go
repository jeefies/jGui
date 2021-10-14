package jgui

import (
    "log"
    _ "sync/atomic"

    "sdl"
)

// The frequence to update the window (flash times per second)
// 图形化界面的帧率，每秒刷新的次数
var FPS uint32 = 35

// Init sdl method automatically when import the package
// 在引入jgui模块时调用sdl的初始化(非编译时进行)
func init() {
	Init()
}

func Init() {
	sdl.Init(sdl.INIT_DEFAULT)
}

// Close sdl unit  when end the program
// 最后关闭sdl界面，调用c sdl的SDL_Quit结构，定义在package sdl中
func Quit() {
    print("jgui Quit\n")
	sdl.Quit()
}

// if err is not nil, panic the error (as a unexpected mistake)
// 把错误当作异常抛出，在jgui中，几乎每一个错误都会导致界面运行失败
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

// like time.Sleep, but use interface of sdl
// 想time标准库中的Sleep方法，只是调用的时sdl的接口(SDL_Delay)，注意单位是毫秒
func Delay(ms uint32) {
	sdl.Delay(ms)
}

// The main body to run the program.
// 整个gui模块的主体部分。类似于c sdl中的大循环
func Mainloop() {
    // log.Println("Mainloop start")
	_update_func_id := new(sdl.TimerID)
    for i, win := range winlist {
        for j, wg := range win.childs {
            // FIXME: Would not draw widgets for the first tim
            // wg.Draw(0)
            // Solution, use event "deactive to draw the default statement"
            wg.Call("deactive")
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
