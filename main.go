package main

import (
	"time"

	"jGui/jgui"
)

func main() {
	// jgui.Init() Move into init() function in private package
	defer jgui.Quit()

	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_DEFAULT)
	defer win.Close()

	// lb := win.NewLabel("Fuck you")

	jgui.Mainloop()
	time.Sleep(2e9);
}
