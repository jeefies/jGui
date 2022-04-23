/*
A Detailed Shape drawing module.
*/
package jgui

/*
#include <SDL2/SDL.h>

void dBorder(SDL_Surface * sur, SDL_Rect * area, int width, Uint32 color);
void dCircleOBorder(SDL_Surface * sur, int cx, int cy, int r, Uint32 color);
void dCircleFilled(SDL_Surface * sur, int cx, int cy, int r, Uint32 color);
void dOLine(SDL_Surface * sur, int x1, int y1, int x2, int y2, Uint32 color);
void dLine(SDL_Surface * sur, int x1, int y1, int x2, int y2, int width, Uint32 color);
*/
import "C"

import (
	"unsafe"

	"jGui/sdl"
)

func DrawBorder(scr * Screen, area *sdl.Rect, width int, color Color) {
	C.dBorder(SCREEN(scr), RECT(area), C.int(width), C.Uint32(color.Map(scr)))
}

func DrawOLine(scr * Screen, x1, y1, x2, y2 int, color Color) {
	C.dOLine(SCREEN(scr), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.Uint32(color.Map(scr)))
}

func DrawLine(scr * Screen, x1, y1, x2, y2, width int, color Color) {
	C.dLine(SCREEN(scr), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(width), C.Uint32(color.Map(scr)))
}

func DrawCircleOBorder(scr *Screen, cx, cy, r int, color Color) {
	C.dCircleOBorder(SCREEN(scr), C.int(cx), C.int(cy), C.int(r), C.Uint32(color.Map(scr)))
}

func DrawCircle(scr *Screen, cx, cy, r int, color Color) {
	C.dCircleFilled(SCREEN(scr), C.int(cx), C.int(cy), C.int(r), C.Uint32(color.Map(scr)))
}

func SCREEN(scr * Screen) *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(scr))
}

func RECT(r *sdl.Rect) *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
}
