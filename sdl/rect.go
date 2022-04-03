package sdl

/*
#include<SDL2/SDL.h>
 */
import "C"

import "fmt"

// Rect part
func NewRect(x, y, w, h int) (*Rect) {
	return &Rect{C.int(x), C.int(y), C.int(w), C.int(h)}
}

func (r * Rect) FromPoint(pos Point) {
	r.x, r.y = C.int(pos.X), C.int(pos.Y)
}

func (r * Rect) String() string {
    return fmt.Sprintf("x, y, w, h = %d, %d, %d, %d", r.x, r.y, r.w, r.h)
}

func (p Point) ToRect() (*Rect) {
	return NewRect(p.X, p.Y, 0, 0)
}
