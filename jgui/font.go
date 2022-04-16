package jgui

import (
	"os"

	"jGui/sdl/ttf"
)

var LoadedFonts [](*Font)

var DefaultFontFile = "fonts/Source-Han-Sans-Regular.ttf"

type Font struct {
	id ID
	fontFile string
	fontCaches map[int] *ttf.Font
}

var defaultFont *Font

func init() {
	LoadedFonts = make([]*Font, 0, 5)
	defaultFont = NewFont(DefaultFontFile)
}

func NewFont(fontFile string) *Font {
	font := new(Font)

	font.fontFile = fontFile
	_, err := os.Stat(fontFile)
	if os.IsNotExist(err) { panic(err) }

	font.fontCaches = make(map[int] *ttf.Font, 5)

	font.id = uint32(len(LoadedFonts))
	LoadedFonts = append(LoadedFonts, font)

	return font
}

func GetFontById(id ID) *Font {
	if int(id) >= len(LoadedFonts) {
		return nil
	}

	return LoadedFonts[id]
}

func (f *Font) Render(size int, text string, color Color) *Screen {
	sur, err := f.CheckSize(size).RenderText(text, color)
	check(err)

	return sur
}

func (f *Font) RenderShaded(size int, text string, colorfg, colorbg Color) *Screen {
	sur, err := f.CheckSize(size).RenderShaded(text, colorfg, colorbg)
	check(err)

	return sur
}

func (f *Font) RenderQuick(size int, text string, color Color) *Screen {
	sur, err := f.CheckSize(size).RenderQuick(text, color)
	check(err)

	return sur
}

func (f *Font) CheckSize(size int) *ttf.Font {
	if _, ok := f.fontCaches[size]; !ok {
		ft, err := ttf.OpenFont(f.fontFile, size)
		check(err)

		f.fontCaches[size] = ft
	}
	return f.fontCaches[size]
}

func (f *Font) GuessSize(size int, text string) (int, int) {
	ft := f.CheckSize(size)
	logger.Printf("%p", ft)
	w, h, err := ft.GuessSize(text)
	check(err)

	return w, h
}
