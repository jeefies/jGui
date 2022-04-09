package jgui

import (
	"jGui/sdl"
	"jGui/sdl/ttf"
)

var DefaultFontFile = "NotoSansMono-Regular.ttf"
var fontCache map[int] (*ttf.Font)

var (
	WHITE = Color{255, 255, 255, 0}
	BLACK = Color{0, 0, 0, 0}
	RED = Color{255, 0, 0, 0}
	GREEN = Color{0, 255, 0, 0}
	BLUE = Color{0, 0, 255, 0}
)
var (
	DefaultTextColor = WHITE
	DefaultBackgroundColor = BLACK
	DefaultBorderColor = WHITE
)

var (
	DefaultActiveTextColor = WHITE
	DefaultActiveBackgroundColor = RED
	DefaultActiveBorderColor = WHITE
)

// type Color = sdl.Color

type BaseWidget struct {
	Text string
	FontSize int
	Align int
	w, h int
	BorderWidth int
	Padding int

	TextColor, BackgroundColor, BorderColor Color
	actived bool
	ActiveTextColor, ActiveBackgroundColor, ActiveBorderColor Color

	id ID
	Parent *Window

	EventList map[WidgetEvent]func(we WidgetEvent, wg Widget)
}

type Label struct {
	sur *sdl.Surface

	min_w, min_h int

	BaseWidget
}

func init() {
	fontCache = make(map[int] (*ttf.Font))
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

func (bw *BaseWidget) Size() (w, h int) {
	w, h = bw.w, bw.h
	return
}

func (bw *BaseWidget) Update() {
	bw.Clear()
	bw.Parent.UpdateWidget(bw.id)
}

func (bw *BaseWidget) Clear() {
	pt := bw.Parent
	area, _ := pt.GetWidgetArea(bw.Id())
	pt.GetOriginScreen().Fill(area.ToSDL(), pt.BackgroundColor.Map(pt.GetScreen()))
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

func (bw BaseWidget) Id() ID {
	return bw.id
}

func NewLabel(text string, font_size int, colors ...Color) (*Label) {
	var err error

	font, ok := fontCache[font_size]
	if !ok {
		font, err = ttf.OpenFont(DefaultFontFile, font_size)
		check(err)
		fontCache[font_size] = font
	}

	lb := new(Label)
	lb.Text = text
	lb.FontSize = font_size
	w, h, err := font.GuessSize(text)
	check(err)
	lb.min_w = w 
	lb.min_h = h

	lb.TextColor = DefaultTextColor
	lb.BackgroundColor = DefaultBackgroundColor
	lb.BorderColor = DefaultBorderColor

	lb.ActiveTextColor = DefaultTextColor
	lb.ActiveBorderColor = DefaultBorderColor
	lb.ActiveBackgroundColor = DefaultBackgroundColor

	switch len(colors) {
	case 6:
		lb.ActiveBorderColor = colors[5]
		fallthrough
	case 5:
		lb.ActiveBackgroundColor = colors[4]
		fallthrough
	case 4:
		lb.ActiveTextColor = colors[3]
		fallthrough
	case 3:
		lb.BorderColor = colors[2]
		fallthrough
	case 2:
		lb.BackgroundColor = colors[1]
		fallthrough
	case 1:
		lb.TextColor = colors[0]
	}

	lb.sur = nil
	lb.id = ID_NULL

	lb.BorderWidth = 2
	lb.Padding = 2

	lb.Align = ALIGN_CENTER

	lb.SetWidth(0)
	lb.SetHeight(0)

	lb.EventList = make(map[WidgetEvent]func(WidgetEvent, Widget), 5)

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

	// Check for label color
	// Label does not have active state
	var (
		textColor = l.TextColor
		backgroundColor = l.BackgroundColor
		borderColor = l.BorderColor
	)

	// Get Text Surface to fill
	if (l.sur == nil) {
		l.sur, err = fontCache[l.FontSize].RenderText(l.Text, textColor)
		check(err)
	}
	mw, mh := l.sur.Size()
	// Check for best widget size	
	area.SetW(MAX(area.W(), mw + l.BorderWidth * 2 + l.Padding * 2))
	area.SetH(MAX(area.H(), mh + l.BorderWidth * 2 + l.Padding * 2))

	{ // Clear Origin
		sur.Fill(area.ToSDL(), backgroundColor.Map(sur))
	}

	{ // Draw Border
		sur.DrawBorder(area.ToSDL(), l.BorderWidth, borderColor)
	}

	{ // Draw Text
		AlignSet(mw, mh, l.BorderWidth, l.Padding, area, l.Align)

		err = sur.Blit(l.sur, area.ToSDL())
		check(err)
	}
}

func (l *Label) Call(e WidgetEvent) {
	if f, ok := l.EventList[e]; ok {
		f(e, l)
		l.Parent.UpdateWidget(l.Id())
	}
}

func (l *Label) Pack(w *Window, area * Rect) *Label {
	l.id = w.Register(l, area)
	logger.Printf("Regist id %d", l.id)
	l.Parent = w
	return l
}

func (l *Label) Configure(method string, value interface{}) *Label {
	switch method {
	case "align":
		if val, ok := value.(int); ok {
			l.Align = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "text":
		if val, ok := value.(string); ok {
			l.Text = val
			w, h, err := fontCache[l.FontSize].GuessSize(l.Text)
			check(err)
			l.min_w = w
			l.min_h = h
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "font size":
		if val, ok := value.(int); ok {
			l.FontSize = val

			_, ok := fontCache[val]
			if !ok {
				font, err := ttf.OpenFont(DefaultFontFile, val)
				check(err)
				fontCache[val] = font
			}
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	}

	if l.id != ID_NULL { 
		l.Clear()
		l.Parent.UpdateWidget(l.id)
	}
	return l
}

func (l *Label) Clear() {
	l.BaseWidget.Clear()

	l.sur.Close()
	l.sur = nil
}
