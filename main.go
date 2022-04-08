package main

import (
	"time"

	"jGui/jgui"
)

func TestWindow1() *jgui.Window {
	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_DEFAULT)
	// win.SetUpdateMode(jgui.WIN_UPDATE_RENDER)

	fontSize := 15
	area := jgui.NewRect(5, 5, 45, 45)
	lb := jgui.NewLabel("L", fontSize)
	lb.Configure("align", jgui.ALIGN_LEFT)
	lb.Pack(win, area)

	jgui.NewLabel("R", fontSize).Configure("align", jgui.ALIGN_RIGHT).Pack(win, jgui.NewRect(55, 5, 45, 45))
	jgui.NewLabel("T", fontSize).Configure("align", jgui.ALIGN_TOP).Pack(win, jgui.NewRect(105, 5, 45, 45))
	jgui.NewLabel("D", fontSize).Configure("align", jgui.ALIGN_BOTTOM).Pack(win, jgui.NewRect(155, 5, 40, 45))

	return win
}

func main() {
	// jgui.Init() Move into init() function in private package
	defer jgui.Quit()

	win := TestWindow1()
	defer win.Close()

	go func() {
		time.Sleep(time.Second * 2)
		lb1, _ := win.GetWidget(0)
		lb1.(*jgui.Label).Configure("align", jgui.ALIGN_LEFT | jgui.ALIGN_TOP)
	}()

	jgui.Mainloop()
}
