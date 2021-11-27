package main

import (
    _ "log"
)

import "jgui"

func main() {
	// jgui.Init() Move into init() function in private package

	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_SHOWN)

	win2 := jgui.CreateWindow("Hello Gui2", 200,  200, jgui.WIN_SHOWN)
	win2.DrawRectCentered(100, 100, 90, 0x55ffffff)

	win.DrawRectCentered(15, 15, 10, 0xff00)

	jgui.UpdateMethod(jgui.UPDATE_BY_SURFACE)

	win.DrawRect(0, 0, 1, 1, 0xffffff)
	btn := jgui.NewButton(15, 50, 100, 30, "Button")
	btn.Pack(win)
	btn.RegistEvent("mouse up", func(wg jgui.Widgets) {
		jgui.Print("mouse up\n")
	})

	lb := jgui.NewLabel(40,  15, 0, 0, "My Label")
	lb.Pack(win)

	print("Mainloop start!\n")
	win.Update()

	jgui.Mainloop()

	win.Close()

	jgui.Quit()
}
