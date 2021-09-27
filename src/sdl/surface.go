/*
Package sdl (main file sdl.go)
This is the file of structure of SDL_Surface
Defines the apis to use the functions to edit the surface (serves by SDL)
*/
package sdl

/*
#include<SDL2/SDL.h>
 */
import "C"

// Returns a SDLError (define in types.go)
// Surface.FillRect is the OO typed interface of SDL_FillRect
// C interface is SDL_FillRect(Rect*, Uint32)
func (sur * Surface) FillRect(rect * Rect, color uint32) error {
	if (C.SDL_FillRect(sur, rect, C.Uint32(color)) != C.int(0)) {
		return NewSDLError("Could Not Fill Surface")
	}
	return nil
}

// Surface.Fill is the same as FillRect
func (sur *Surface) Fill(rect * Rect, color uint32) error {
	return sur.FillRect(rect, color)
}

// Surface.Blit returns a SDLError if not success
// This is the OO typed interface of SDL_BlitSurface
// @param src Surface is the surface which would be copied into the calling Surface
// place is a Rect pointer which shows which area of src would be copied to parent surface
// is place is nil, copy all the content of src into the parent surface to top left point
func (sur * Surface) Blit(src * Surface, place * Rect)  error {
    if C.SDL_BlitSurface(src, nil, sur, place) != C.int(0) {
        return NewSDLError("Could not apply surface into surface")
    }
    return nil
}

// Clear means fill the surface with all black
func (sur *Surface) Clear() error {
	return sur.FillRect(nil, 0)
}

// ClearWith want a color (uint32 0xAARRBBGG) to fill all the surface
func (sur *Surface) ClearWith(color uint32) error {
	return sur.FillRect(nil, color)
}

// Size returns two integer (int, int) shows the width and the height of the surface
func (sur *Surface) Size() (int, int) {
	return int(sur.w), int(sur.h)
}

// Close returns nothing but free the surface. That means it should not be used anymore after calling this function
func (sur *Surface) Close() {
	C.SDL_FreeSurface(sur)
}
