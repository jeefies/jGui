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
var DefaultFg = WHITE
var DefaultBg = WHITE
var DefaultAFg = RED
var DefaultABg = WHITE

// type Color = sdl.Color

type BaseWidget struct {
	Text string
	FontSize int
	Align int
	w, h int
	Fg, Bg Color
	Afg, Abg Color // Active fg, Active bg
	BorderWidth int
	Padding int

	id ID
	Parent *Window
}

type Label struct {
	actived bool
	sur *sdl.Surface

	min_w, min_h int

	BaseWidget
}

func init() {
	fontCache = make(map[int] (*ttf.Font))
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

	lb.Fg = DefaultFg
	lb.Bg = DefaultBg
	lb.Afg = DefaultFg
	lb.Abg = DefaultBg
	switch len(colors) {
	case 4:
		lb.Afg = colors[3]
		lb.Abg = colors[4]
		fallthrough
	case 2:
		lb.Fg = colors[0]
		lb.Bg = colors[1]
	}

	lb.sur = nil
	lb.id = ID_NULL

	lb.BorderWidth = 2
	lb.Padding = 2

	lb.SetWidth(0)
	lb.SetHeight(0)

	lb.Align = ALIGN_CENTER

	return lb
}

func (lb *Label) Size() (w, h int) {
	w, h = lb.w, lb.h
	return
}

func (lb *Label) Width() int {
	return lb.w
}

func (lb *Label) Height() int {
	return lb.h
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
	fg, bg := l.Fg, l.Bg
	if l.actived {
		fg = l.Afg
		bg = l.Abg
	}

	// Get Text Surface to fill
	if (l.sur == nil) {
		l.sur, err = fontCache[l.FontSize].RenderText(l.Text, fg)
		check(err)
	}
	l.w, l.h = l.sur.Size()
		
	area.w = MAX(area.w, l.w + l.BorderWidth * 2 )
	area.h = MAX(area.h, l.h + l.BorderWidth * 2 )

	{ // Clear Origin
		sur.Fill(area.ToSDL(), l.Parent.bgColor.Map(sur))
	}

	{ // Draw Border
		sur.DrawBorder(area.ToSDL(), l.BorderWidth, bg)
	}

	{ // Draw Text
		bdw := l.BorderWidth
		pdw := l.Padding

		if l.Align & ALIGN_LEFT == ALIGN_LEFT {
			area.x += bdw + pdw
		} else if l.Align & ALIGN_RIGHT == ALIGN_RIGHT {
			area.x = area.x + area.w - l.w - bdw - pdw
		} else {
			area.x = area.x + area.w / 2 - l.w / 2
		}

		if l.Align & ALIGN_TOP == ALIGN_TOP {
			area.y += bdw + pdw
		} else if l.Align & ALIGN_BOTTOM == ALIGN_BOTTOM {
			area.y = area.y + area.h - l.h - bdw - pdw
		} else {
			area.y = area.y + area.h / 2 - l.h / 2
		}

		err = sur.Blit(l.sur, area.ToSDL())
		check(err)
	}
}

func (l *Label) Call(e WidgetEvent) {
	switch e  {
	case WE_IN:
		if l.actived == false {
			l.Clear()
		}
		l.actived = true
	case WE_OUT:
		if l.actived == true {
			l.Clear()
		}
		l.actived = false
	default:
	}
	logger.Printf("Call, now active: %v", l.actived)
	l.Parent.UpdateWidget(l.Id())
}

func (l Label) Id() ID {
	return l.id
}

func (l *Label) Pack(w *Window, area * Rect) *Label {
	l.id = w.Register(l, area.x, area.y, MAX(area.w, l.w), MAX(area.h, l.h))
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
	case "bg":
		if val, ok := value.(Color); ok {
			l.Bg = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "abg":
		if val, ok := value.(Color); ok {
			l.Abg = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "fg":
		if val, ok := value.(Color); ok {
			l.Fg = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "afg":
		if val, ok := value.(Color); ok {
			l.Afg = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "text":
		if val, ok := value.(string); ok {
			l.Text = val
			w, h, err := fontCache[l.FontSize].GuessSize(l.Text)
			check(err)
			l.w = w + 6
			l.h = h + 6
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "font size":
		if val, ok := value.(int); ok {
			l.FontSize = val
			w, h, err := fontCache[l.FontSize].GuessSize(l.Text)
			check(err)
			l.w = w + 6
			l.h = h + 6
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "border width":
		if val, ok := value.(int); ok {
			l.BorderWidth = val
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	}

	if l.id != ID_NULL { 
		l.Clear()
		l.Parent.UpdateWidget(l.id)
	}
	return l
}

func (l *Label) Clear() {
	l.sur.Close()
	l.sur = nil
}
