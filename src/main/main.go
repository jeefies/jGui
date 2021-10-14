package main

import (
    _ "log"
)

import "jgui"

func main() {
	// jgui.Init() Move into init() function in private package

	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_SHOWN)

	win.DrawRectCentered(15, 15, 10, 0xff00)

	jgui.UpdateMethod(jgui.UPDATE_BY_SURFACE)

	win.DrawRect(0, 0, 1, 1, 0xffffff)
	btn := jgui.NewButton(15, 50, 100, 30, "Button")
	btn.Pack(win)
    btn.RegistEvent("mouse up", func(wg jgui.Widgets) {
       jgui.Print("mouse up\n")
    })

    lb := jgui.NewLabel(100, 100, 10, 10, "My Label")
    lb.Pack(win)
    
	/*
	r := jgui.NewRect(1, 1, 5, 5)
	jgui.FillRect(win, r, 0xff00ff)
	*/

	// win.DrawCircle(100, 100, 100, 0xff0000)

	print("Mainloop start!\n")
    win.Update()

	jgui.Mainloop()

	win.Close()

	jgui.Quit()
}
