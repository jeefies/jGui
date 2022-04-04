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

type Label struct {
	text string
	font_size int
	align int
	w, h int
	Fg, Bg Color
	Afg, Abg Color // Active fg, Active bg

	actived bool
	sur *sdl.Surface

	parent *Window
	id int
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
	lb.text = text
	lb.font_size = font_size
	w, h, err := font.GuessSize(text)
	check(err)
	lb.w = w + 6
	lb.h = h + 6

	lb.Fg = DefaultFg
	lb.Bg = DefaultBg
	lb.Afg = DefaultAFg
	lb.Abg = DefaultABg
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
	lb.id = -1

	lb.align = ALIGN_CENTER

	return lb
}

func (lb *Label) Size() (w, h int) {
	w, h = lb.w, lb.h
	return
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
		l.sur, err = fontCache[l.font_size].RenderText(l.text, fg)
		check(err)
	}
	l.w, l.h = l.sur.Size()
	logger.Printf("Font sur w, h = %d %d\n", l.w, l.h)
		
	area.w = MAX(area.w, l.w + 6)
	area.h = MAX(area.h, l.h + 6)

	{ // Clear Origin
		sur.Fill(area.ToSDL(), BLACK.Map(sur))
	}

	{ // Draw Border
		sur.DrawBorder(area.ToSDL(), 2, bg)
	}

	{ // Draw Text
		switch l.align {
		case ALIGN_LEFT:
			area.x += 3
		case ALIGN_RIGHT:
			area.x = area.x + area.w - 3 - l.w
		case ALIGN_CENTER:
			area.x = area.x + area.w / 2 - l.w / 2
		}
		area.y = area.y + area.h / 2 - l.h / 2
		err = sur.Blit(l.sur, area.ToSDL())
		check(err)
	}
}

func (l *Label) Call(e WidgetEvent) {
	switch e  {
	case WE_IN:
		l.actived = true
	case WE_OUT:
		l.actived = false
	default:
	}
}

func (l Label) Id() int {
	return l.id
}

func (l *Label) Pack(w *Window, area * Rect) {
	l.id = w.Register(l, area.x, area.y, MAX(area.w, l.w), MAX(area.h, l.h))
	l.parent = w
}

func (l *Label) Configure(method string, value interface{}) {
	switch method {
	case "align":
		if val, ok := value.(int); ok {
			l.align = val
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
			l.text = val
			w, h, err := fontCache[l.font_size].GuessSize(l.text)
			check(err)
			l.w = w + 6
			l.h = h + 6
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	case "font_size":
		if val, ok := value.(int); ok {
			l.font_size = val
			w, h, err := fontCache[l.font_size].GuessSize(l.text)
			check(err)
			l.w = w + 6
			l.h = h + 6
		} else { panic(sdl.NewSDLError("Not Valid config value")) }
	}
}
