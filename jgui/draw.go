/*
About The Code in C. See draw.c
Provide the Interface of drawing shapes
*/
package jgui

/*
#cgo pkg-config: sdl2
#include <SDL.h>
int j_drawrect_centered(SDL_Surface * sur, int x, int y, int r, Uint32 c);
int j_drawrect_centered2(SDL_Surface * sur, int x, int y, int r, Uint32 c);
int j_draw_circle(SDL_Surface * sur, int x, int y, int r, Uint32 c);
int j_drawrect_border(SDL_Surface * sur, int x, int y, int xw, int yh, int width, Uint32 color);
*/
import "C"
import "jGui/sdl"

import "unsafe"

func (win * Window) DrawRect(x1, y1, x2, y2 int, color uint32) {
	rect := sdl.NewRect(x1, y1, x2 - x1, y2 - y1)
	win._scr.FillRect(rect, color)
}

func (win * Window) DrawRectCentered(x, y, r int, color uint32) {
	var cx, cy, cr C.int = C.int(x), C.int(y), C.int(r)
	var ccolor C.Uint32 = C.Uint32(color)

	C.j_drawrect_centered((*C.SDL_Surface)(unsafe.Pointer(win._scr)), cx, cy, cr, ccolor)
}

func (win * Window) DrawRectCentered2(x, y, r int, color uint32) {
	var cx, cy, cr C.int = C.int(x), C.int(y), C.int(r)
	var ccolor C.Uint32 = C.Uint32(color)

	C.j_drawrect_centered2((*C.SDL_Surface)(unsafe.Pointer(win._scr)), cx, cy, cr, ccolor)
}

func (win * Window) DrawRectBorder(x1, y1, x2, y2, width int, color uint32) {
	var a, b, c, d, w = C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(width)
	var ccolor = C.Uint32(color)

	C.j_drawrect_border((*C.SDL_Surface)(unsafe.Pointer(win._scr)), a, b, c, d, w, ccolor)
}

func (win * Window) DrawCircle(x, y, r int, color uint32) {
	cx, cy, cr := C.int(x), C.int(y), C.int(r)
	cc := C.Uint32(color)

	C.j_draw_circle((*C.SDL_Surface)(unsafe.Pointer(win._scr)), cx, cy, cr, cc)
}
