package main

import (
	"time"

	jgui "jGui/jgui"
)

func TestWindow1() *jgui.Window {
	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_DEFAULT)
	// win.SetUpdateMode(jgui.WIN_UPDATE_RENDER)

	fontSize := 15
	area := jgui.NewRect(5, 5, 45, 45)
	lb := jgui.NewLabel("L", fontSize)
	lb.Align = jgui.ALIGN_LEFT
	lb.Pack(win, area)

	lb1 := jgui.NewLabel("R", fontSize).Pack(win, jgui.NewRect(55, 5, 45, 45))
	lb1.Align = jgui.ALIGN_RIGHT
	jgui.NewLabel("T", fontSize).Pack(win, jgui.NewRect(105, 5, 45, 45)).Align = jgui.ALIGN_TOP
	jgui.NewLabel("D", fontSize).Pack(win, jgui.NewRect(155, 5, 40, 45)).Align = jgui.ALIGN_BOTTOM

	lb1.BackgroundColor = jgui.Color{0, 255, 0, 0}
	lb1.RegisterEvent(jgui.WE_FOCUSIN, func(we jgui.WidgetEvent, wg jgui.Widget) {
		lb := wg.(*jgui.Label)
		lb.TextColor = jgui.BLUE
	})
	lb1.RegisterEvent(jgui.WE_FOCUSOUT, func(we jgui.WidgetEvent, wg jgui.Widget) {
		lb := wg.(*jgui.Label)
		lb.TextColor = jgui.WHITE
	})

	return win
}

func TestWindow2() *jgui.Window {
	win := jgui.CreateWindow("Hello J_Gui", 200, 200, jgui.WIN_DEFAULT)

	area := jgui.NewRelRect(5, 5, 0.2, 0.2, jgui.REL_W | jgui.REL_H)

	lb := jgui.NewLabel("Rel", 15).Pack(win, area)

	ip := jgui.NewInput(20).Pack(win, jgui.NewRelRect(5, 0.3, 0.8, 35, jgui.REL_Y | jgui.REL_W))
	ip.AutoFocus = true

	ip.Font = jgui.NewFont("fonts/Retro Rescued.ttf")

	ip.RegisterEvent(jgui.WE_KEY, func (we jgui.WidgetEvent, wg jgui.Widget) {
		input := wg.(*jgui.Input)
		switch input.Key() {
		case jgui.KRETURN:
			lb.OnEdit()
			lb.Text  = input.Text
			lb.EndEdit()
		}
	})


	return win
}

func main() {
	// jgui.Init() Move into init() function in private package
	defer jgui.Quit()

	win := TestWindow1()

	TestWindow2()

	go func() {
		time.Sleep(time.Second * 2)
		lb1, _ := win.GetWidget(0)
		lb := lb1.(*jgui.Label)

		lb.OnEdit()
		lb.Align = jgui.ALIGN_LEFT | jgui.ALIGN_TOP
		lb.FontSize = 16
		lb.EndEdit()

		lb.Update()

		time.Sleep(time.Second * 2)
		win.MoveWidgetTo(lb.Id(), jgui.NewRelRect(5, 0.5, 45, 45, jgui.REL_Y))
	}()

	jgui.Mainloop()
}
