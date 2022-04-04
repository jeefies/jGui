package jgui

import (
	"jGui/sdl"
)

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
	return
}

func Quit() {
	sdl.Quit()
	return
}
