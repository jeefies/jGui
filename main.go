package main

import (
	"jGui/jgui"
)

func main() {
	// jgui.Init() Move into init() function in private package
	defer jgui.Quit()

	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_DEFAULT)
	defer win.Close()
	win.SetUpdateMode(jgui.WIN_UPDATE_RENDER)

	area := jgui.NewRect(10, 10, 50, 50)
	lb := jgui.NewLabel("Fuck", 15)
	lb.Pack(win, area)

	jgui.Mainloop()
}
