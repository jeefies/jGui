package jgui

import (
	"sync"

	"jGui/sdl"
)

var (
	WHITE = Color{255, 255, 255, 0}
	BLACK = Color{0, 0, 0, 0}
	RED = Color{255, 0, 0, 0}
	GREEN = Color{0, 255, 0, 0}
	BLUE = Color{0, 0, 255, 0}
	GREY = Color{128, 128, 128, 0}
	DARKGREY = Color{169, 169, 169, 0}
	LIGHTGREY = Color{211, 211, 211, 0}
	SILVER = Color{192, 192, 192, 0}
)

var (
	DefaultTextColor = WHITE
	DefaultBackgroundColor = BLACK
	DefaultBorderColor = WHITE
)

var (
	DefaultActiveTextColor = WHITE
	DefaultActiveBackgroundColor = BLACK
	DefaultActiveBorderColor = RED
)

// type Color = sdl.Color

type BaseWidget struct {
	Text string
	FontSize int
	Font *Font
	Align int
	w, h int
	BorderWidth int
	Padding int

	TextColor, BackgroundColor, BorderColor Color
	actived bool
	ActiveTextColor, ActiveBackgroundColor, ActiveBorderColor Color

	id ID
	Parent *Window
	ActualArea *Rect
	TextSur *Screen

	AutoFocus bool
	EventList map[WidgetEvent]func(we WidgetEvent, wg Widget)

	sync.Mutex
}

type Label struct {
	min_w, min_h int

	BaseWidget
}

type Input struct {
	edisur *sdl.Surface

	currentRune []rune
	editingText string
	afterEdit bool
	cursorPlace int

	BaseWidget
}

func init() {
}

func AlignSet(w, h int, bdw, pdw int, area *Rect, alignFlags int) {
	if alignFlags & ALIGN_LEFT == ALIGN_LEFT {
		area.SetX(area.X() + bdw + pdw)
	} else if alignFlags & ALIGN_RIGHT == ALIGN_RIGHT {
		area.SetX(area.X() + area.W() - w - bdw - pdw)
	} else {
		area.SetX(area.X() + (area.W() - w) / 2)
	}

	if alignFlags & ALIGN_TOP == ALIGN_TOP {
		area.SetY(area.Y() + bdw + pdw)
	} else if alignFlags & ALIGN_BOTTOM == ALIGN_BOTTOM {
		area.SetY(area.Y() + area.H() - h - bdw - pdw)
	} else {
		area.SetY(area.Y() + (area.H() - h) / 2)
	}
}

func InitBaseWidget(bw *BaseWidget, text string, font_size int, colors ...Color) {
	bw.Text = text
	bw.FontSize = font_size

	bw.TextColor = DefaultTextColor
	bw.BackgroundColor = DefaultBackgroundColor
	bw.BorderColor = DefaultBorderColor

	bw.ActiveTextColor = DefaultTextColor
	bw.ActiveBorderColor = DefaultActiveBorderColor
	bw.ActiveBackgroundColor = DefaultBackgroundColor

	switch len(colors) {
	case 6:
		bw.ActiveBorderColor = colors[5]
		fallthrough
	case 5:
		bw.ActiveBackgroundColor = colors[4]
		fallthrough
	case 4:
		bw.ActiveTextColor = colors[3]
		fallthrough
	case 3:
		bw.BorderColor = colors[2]
		fallthrough
	case 2:
		bw.BackgroundColor = colors[1]
		fallthrough
	case 1:
		bw.TextColor = colors[0]
	}

	bw.id = ID_NULL

	bw.BorderWidth = 2
	bw.Padding = 2

	bw.Align = ALIGN_LEFT
	bw.Font = defaultFont

	bw.EventList = make(map[WidgetEvent]func(WidgetEvent, Widget), 5)
}

func (bw *BaseWidget) Size() (w, h int) {
	w, h = bw.w, bw.h
	return
}

func (bw *BaseWidget) Update() {
	bw.CleanUp()
	bw.Parent.UpdateWidget(bw.id)
}

func (bw *BaseWidget) CleanUp() {
	bw.Lock()
	defer bw.Unlock()

	if bw.TextSur != nil {
		bw.TextSur.Close()
		bw.TextSur = nil
	}

	pt := bw.Parent

	pt.Lock()
	defer pt.Unlock()

	// Once there's a segmentation fault because of the null pointer after window closed
	if pt._scr == nil { return }

	pt.GetOriginScreen().Fill(bw.ActualArea.ToSDL(), pt.BackgroundColor.Map(pt.GetScreen()))
}

func (bw *BaseWidget) Width() int {
	return bw.w
}

func (bw *BaseWidget) Height() int {
	return bw.h
}

func (bw *BaseWidget) SetWidth(w int) int {
	bw.w = w
	return w
}

func (bw *BaseWidget) SetHeight(h int) int {
	bw.h = h
	return h
}

func (bw *BaseWidget) RegisterEvent(we WidgetEvent, f func(we WidgetEvent, wg Widget)) {
	bw.EventList[we] = f
}

func (bw *BaseWidget) OnEdit() {
	bw.Lock()
}

func (bw *BaseWidget) EndEdit() {
	bw.Unlock()
}

func (bw BaseWidget) Id() ID {
	return bw.id
}

func NewLabel(text string, font_size int, colors ...Color) (*Label) {
	lb := new(Label)
	InitBaseWidget(&lb.BaseWidget, text, font_size, colors...)

	w, h := lb.Font.GuessSize(font_size, text)
	lb.min_w = w 
	lb.min_h = h

	lb.SetWidth(0)
	lb.SetHeight(0)


	return lb
}

func (lb *Label) SetWidth(w int) int {
	lb.w = MAX(ABS(w), lb.min_w + lb.BorderWidth * 2 + lb.Padding * 2)
	return lb.w
}

func (lb *Label) SetHeight(h int) int {
	lb.h = MAX(ABS(h), lb.min_h + lb.BorderWidth * 2 + lb.Padding * 2)
	return lb.h
}

func (l *Label) Draw(sur *Screen, area * Rect) {
	var err error
	l.Lock()
	defer l.Unlock()

	// Check for label color
	// Label does not have active state
	var (
		textColor = l.TextColor
		backgroundColor = l.BackgroundColor
		borderColor = l.BorderColor
	)

	if l.ActualArea != nil { // Clear Origin
		sur.Fill(l.ActualArea.ToSDL(), l.Parent.BackgroundColor.Map(sur))
	}

	var content = true
	if l.Text == "" { content = false }

	// Get Text Surface to fill
	if (content && l.TextSur == nil) {
		logger.Printf("Render Text %s", l.Text)
		l.TextSur = l.Font.Render(l.FontSize, l.Text, textColor)
	}
	var mw, mh int
	if content {
		mw, mh = l.TextSur.Size()
	}
	// Check for best widget size	
	area.SetW(MAX(area.W(), mw + l.BorderWidth * 2 + l.Padding * 2))
	area.SetH(MAX(area.H(), mh + l.BorderWidth * 2 + l.Padding * 2))

	l.ActualArea = area.Copy()

	{ // Draw Background
		sur.Fill(area.ToSDL(), backgroundColor.Map(sur))
	}

	{ // Draw Border
		DrawBorder(sur, area.ToSDL(), l.BorderWidth, borderColor)
	}

	if content { // Draw Text
		AlignSet(mw, mh, l.BorderWidth, l.Padding, area, l.Align)

		err = sur.Blit(l.TextSur, area.ToSDL())
		check(err)
	}
}

func (l *Label) Call(e WidgetEvent) {
	if f, ok := l.EventList[e]; ok {
		f(e, l)
		l.CleanUp()
		l.Parent.UpdateWidget(l.Id())
	}
}

func (l *Label) Pack(w *Window, area * Rect) *Label {
	l.id = w.Register(l, area)
	logger.Printf("Regist id %d", l.id)
	l.Parent = w
	return l
}

func NewInput(font_size int, colors ...Color) (*Input) {
	ip := new(Input)
	InitBaseWidget(&ip.BaseWidget, "", font_size, colors...)
	ip.edisur = nil
	return ip
}

func (ip *Input) Draw(sur *Screen, area * Rect) {
	ip.Lock()
	defer ip.Unlock()

	var err error
	var editing = true
	var content = true

	var (
		textColor = ip.TextColor
		backgroundColor = ip.BackgroundColor
		borderColor = ip.BorderColor
	)

	if ip.ActualArea != nil { // Clear Origin
		sur.Fill(ip.ActualArea.ToSDL(), backgroundColor.Map(sur))
	}


	if ip.actived {
		borderColor = ip.ActiveBorderColor
	}

	ip.Text = string(ip.currentRune)

	if ip.Text == "" { content = false }
	if ip.editingText == "" { editing = false }

	// Get Text Surface to fill
	if content && ip.TextSur == nil {
		ip.TextSur = ip.Font.Render(ip.FontSize, ip.Text, textColor)
	}

	if editing && ip.edisur == nil {
		ip.edisur = ip.Font.Render(ip.FontSize, ip.editingText, SILVER)
	}

	var ew, eh, mw, mh int
	if content {
		mw, mh = ip.TextSur.Size()
	}
	if editing {
		ew, eh = ip.edisur.Size()
	}

	// Check for best widget size	
	area.SetW(MAX(area.W(), mw + ew + ip.BorderWidth * 2 + ip.Padding * 2))
	area.SetH(MAX(area.H(), mh + eh + ip.BorderWidth * 2 + ip.Padding * 2))
	ip.ActualArea = area.Copy()

	{ // Draw Background
		sur.Fill(area.ToSDL(), backgroundColor.Map(sur))
	}

	{ // Draw Border
		DrawBorder(sur, area.ToSDL(), ip.BorderWidth, borderColor)
	}

	{ // Draw Text
		AlignSet(mw, mh, ip.BorderWidth, ip.Padding, area, ip.Align)

		if content {
			err = sur.Blit(ip.TextSur, area.ToSDL())
			check(err)
		}

		if editing {
			// Draw editing text
			area.SetX(area.X() + mw)
			err = sur.Blit(ip.edisur, area.ToSDL())
			check(err)
		}
	}

	DrawCircle(sur, 50, 140, 40, WHITE)
	DrawLine(sur, 10, 10, 40, 160, 4, RED)
}

func (ip *Input) Call(we WidgetEvent) {
	switch we {
	case WE_FOCUSIN:
		ip.actived = true
		sdl.StartTextInput()
		area, _ := ip.Parent.GetWidgetArea(ip.id)
		if false {
			area = area.MapVH(ip.Parent.Size())
			area.y += area.h
			area.Pout()
		}
		sdl.SetTextInputRect(area.ToSDL())
	case WE_FOCUSOUT:
		ip.actived = false
		sdl.StopTextInput()
	case WE_INIT:
		if ip.AutoFocus {
			ip.Parent.Focus(ip.id)
		}
	case WE_TEXT_INPUT:
		addText := []rune(ip.Parent.Event.InputText())
		ip.currentRune = append(ip.currentRune, addText...)
		logger.Printf("Input add text : %s, now text is %s", string(addText), string(ip.currentRune))
	case WE_TEXT_EDITING:
		ip.editingText = ip.Parent.Event.EditingText()

		ip.edisur.Close()
		ip.edisur = nil

		ip.afterEdit = true
	case WE_KEY:
		switch ip.Key() {
		case sdl.KBACKSPACE:
			l := len(ip.currentRune) - 1
			if (l >= 0) {
				ip.currentRune = ip.currentRune[:l]
			} else {
				ip.currentRune = ip.currentRune[:0]
			}
		case sdl.KESC:
			ip.Parent.FocusOut()	
		}
	}
	if f, ok := ip.EventList[we]; ok {
		f(we, ip)
	}
	ip.CleanUp()
	ip.Parent.UpdateWidget(ip.Id())
}

func (ip *Input) Pack(w *Window, area * Rect) *Input {
	ip.id = w.Register(ip, area)
	logger.Printf("Regist id %d", ip.id)
	ip.Parent = w
	return ip
}

func (ip *Input) CleanUp() {
	ip.BaseWidget.CleanUp()

	ip.Lock()
	defer ip.Unlock()

	if ip.edisur != nil {
		ip.edisur.Close()
		ip.edisur = nil
	}
}

func (ip *Input) Key() uint32 {
	return ip.Parent.Event.Key()
}
