package ttf

import "jGui/sdl"
import "unsafe"

/*
#cgo pkg-config: sdl2 SDL2_ttf
#include <SDL2/SDL.h>
#include <SDL2/SDL_ttf.h>
*/
import "C"

type Font    = C.TTF_Font
type Color   = sdl.Color
type Surface = sdl.Surface

func init() {
	if C.TTF_Init() != C.int(0) {
		panic(sdl.NewSDLError("Could Not Init ttf"))
	}
}

func OpenFont(fontfile string, size int) (font *Font, err error) {
	font = C.TTF_OpenFont(C.CString(fontfile), C.int(size))
	if (font == nil) {
		err = sdl.NewSDLError("Could Not Load Font")
	}
	return
}

func (font *Font) Close() {
	C.TTF_CloseFont(font)
}

func (font *Font) RenderText(str string, color Color) (sur *Surface, err error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	sur = (*Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended(font, cstr, COLOR(color))))
	if (sur == nil) {
		err = sdl.NewSDLError("Could Not render text")
		return
	}
	return
}

func (font *Font) RenderQuick(str string, color Color) (sur *Surface, err error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	sur = (*Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Solid(font, cstr, COLOR(color))))
	if (sur == nil) {
		err = sdl.NewSDLError("Could Not render text")
		return
	}
	return
}

func (font *Font) RenderShaded(str string, colorfg Color, colorbg Color) (sur *Surface, err error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	sur = (*Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Shaded(font, cstr, COLOR(colorfg), COLOR(colorbg))))
	if (sur == nil) {
		err = sdl.NewSDLError("Could Not render text")
		return
	}
	return
}

func COLOR(color Color) (C.SDL_Color) {
	return *(*C.SDL_Color)(unsafe.Pointer(color.SDLColor()))
}

func (font *Font) GuessSize(str string) (w, h int, err error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	var wp = new(C.int)
	var hp = new(C.int)
	if (C.TTF_SizeUTF8(font, cstr, wp, hp) != 0) {
		err = sdl.NewSDLError("Could Not get text size")
		return
	}
	w = int(*wp)
	h = int(*hp)
	return
}
