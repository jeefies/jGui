/*
A Detailed Shape drawing module.
*/
package jgui

/*
#include <SDL2/SDL.h>
#define GI(x,y) (y)*ppr+(x)

// Get Index
#define fillRect(x,y,w,h) {\
	for (int i = 0; i < w; i++) {\
		for (int j = 0; j < h; j++) {\
			pixels[GI(x+i,y+j)] = color;\
		}\
	}\
}

// Get basic data
#define INIT Uint32 * pixels = (Uint32 *)sur->pixels;\
	int ppr = sur->pitch / sur->format->BytesPerPixel;

#define INIT_RECT int x = area->x;\
	int y = area->y;\
	int h = area->h;\
	int w = area->w

int neg(int x) {
	return x<0?1:0;
}

int abs(int x) {
	return x<0?-x:x;
}

void pBorder(SDL_Surface * sur, SDL_Rect * area, int width, Uint32 color) {
	INIT;
	INIT_RECT;

	// Top line
	fillRect(x + width, y, w - 2 * width, width);
	// Botton line
	fillRect(x + width, y + h - width, w - 2 * width, width);
	// Left linne
	fillRect(x, y + width, width, h - 2 * width);
	// Right Line
	fillRect(x + w - width, y + width, width, h - 2 * width);

	// Top left rect
	fillRect(x + width, y + width, width, width);
	// Top Right
	fillRect(x + w - 2 * width, y + width, width, width);
	// botton left
	fillRect(x + width, y + h - 2 * width, width, width);
	// botton right
	fillRect(x + w - 2 * width, y + h - 2 * width, width, width);
}

#define CIRCLE_POINT_DRAW pixels[GI(cx + x, cy + y)] = color;\
		pixels[GI(cx + x, cy - y)] = color;\
		pixels[GI(cx - x, cy + y)] = color;\
		pixels[GI(cx - x, cy - y)] = color

void pCircleOBorder(SDL_Surface *sur, int cx, int cy, int r, Uint32 color) {
	INIT;

	int x = 0, y = r;
	int pr = r * r;

	for (int i = 0; i <= r; i++) {
		CIRCLE_POINT_DRAW;

		while (x * x + y * y >= pr - r) {
			y--;
			CIRCLE_POINT_DRAW;
			if (y <= 0) break;
		}
		x++;
	}
}

void pCircleFilled(SDL_Surface *sur, int cx, int cy, int r, Uint32 color) {
	INIT;
	int x = 0, y = r;
	int pr = r * r;
	for (int i = 0; i <= r; i++) {
		fillRect(cx - x, cy, 1, y);
		fillRect(cx + x, cy, 1, y);
		fillRect(cx - x, cy - y, 1, y);
		fillRect(cx + x, cy - y, 1, y);
		while (x * x + y * y > pr) {
			y--;
			if (y <= 0) break;
		}
		x++;
	}
}

void pOLine(SDL_Surface * sur, int x1, int y1, int x2, int y2, Uint32 color) {
	INIT;

	printf("Draw Oline\n");

	if (x1 > x2) {
		// Make sure x1 is always the left point
		int tmp;
		tmp = x1;x1 = x2;x2 = tmp;
		tmp = y1;y1 = y2;y2 = tmp;
	}

	double ok = ((double)y1 - (double)y2) / ((double)x1 - (double)x2); // 斜率
	double k = ok;
	// 相对与左坐标
	int type, steps = x2 - x1;

	if (k > 1) {
		type = 1;
		k = 1.0 / k;
		steps = y2 - y1;
	} else if (k > 0) {
		type = 2;
	} else if (k > -1) {
		type = 3;
		k = -k;
	} else {
		// k <= -1
		k = 1.0 / -k;
		type = 4;
		steps = y1 - y2;
	}
	printf("k = %lf, ok = %lf, type = %d\n", k, ok, type);

	double d = 0;
	int x = 0, y = 0;
	for (int i = 0; i < steps; i++) {
		x++;
		d += k;

		if (d >= 0.5) {
			y += 1;
			d -= 0.5;
		}
		printf("x, y = %d %d\n", x, y);
		switch (type) {
		case 1: pixels[GI(x1 + y, y1 + x)] = color; break;
		case 2: pixels[GI(x1 + x, y1 + y)] = color; break;
		case 3: pixels[GI(x1 - x, y1 - y)] = color; break;
		case 4: pixels[GI(x1 - y, y1 - x)] = color; break;
		}
	}
}

void pLine(SDL_Surface* scr, int x1, int y1, int x2, int y2, int width, Uint32 color) {
	int r = width / 2;
	if (width % 2) r -= 1;
	pCircleFilled(scr, x1, y1, r, color);
	pCircleFilled(scr, x2, y2, r, color);

	for (int i = 0; i < width; i++) {
		pOLine(scr, x1, y1 - r + i, x2, y2 - r + i, color);
	}
}

*/
import "C"

import (
	"unsafe"

	"jGui/sdl"
)

func DrawBorder(scr * Screen, area *sdl.Rect, width int, color Color) {
	C.pBorder(SCREEN(scr), RECT(area), C.int(width), C.Uint32(color.Map(scr)))
}

func DrawOLine(scr * Screen, x1, y1, x2, y2 int, color Color) {
	C.pOLine(SCREEN(scr), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.Uint32(color.Map(scr)))
}

func DrawLine(scr * Screen, x1, y1, x2, y2, width int, color Color) {
	C.pLine(SCREEN(scr), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(width), C.Uint32(color.Map(scr)))
}

func DrawCircleOBorder(scr *Screen, cx, cy, r int, color Color) {
	C.pCircleOBorder(SCREEN(scr), C.int(cx), C.int(cy), C.int(r), C.Uint32(color.Map(scr)))
}

func DrawCircle(scr *Screen, cx, cy, r int, color Color) {
	C.pCircleFilled(SCREEN(scr), C.int(cx), C.int(cy), C.int(r), C.Uint32(color.Map(scr)))
}

func SCREEN(scr * Screen) *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(scr))
}

func RECT(r *sdl.Rect) *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
}
