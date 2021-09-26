package main

import (
	_ "os"
	_ "strconv"
	_ "math/rand"
	_ "time"
)

import "jgui"

const sqrL int = 12
const (
	LEFT = -2
	RIGHT = 2
	UP = -1
	DOWN = 1
)

func main() {
	jgui.Init()

	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_SHOWN)

	win.DrawRectCentered(15, 15, 10, 0xff00)

	win.Update()

	jgui.UpdateMethod(jgui.UPDATE_BY_SURFACE)

	/*
	r := jgui.NewRect(1, 1, 5, 5)
	jgui.FillRect(win, r, 0xff00ff)
	*/

	win.DrawCircle(100, 100, 100, 0xff0000)

	print("Mainloop start!\n")

	jgui.Mainloop()

	win.Close()

	jgui.Quit()
}
