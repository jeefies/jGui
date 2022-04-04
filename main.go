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
	win.ChangeUpdateMode(jgui.WIN_UPDATE_RENDER)

	area := jgui.NewRect(10, 10, 50, 50)
	lb := jgui.NewLabel("Fuck", 14)
	lb.Pack(win, area)

	area2 := jgui.NewRect(40, 50, 200, 200)
	win.GetScreen().Fill(area2.ToSDL(), 0xff00)

	win.Update()
	win.RenderScreen()
	win.Show()

	time.Sleep(2e9)

	lb.Configure("align", jgui.ALIGN_LEFT)
	win.Update()
	win.RenderScreen()
	win.Show()
	time.Sleep(2e9)
}
