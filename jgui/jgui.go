package jgui

import (
	"jGui/sdl"
)

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
