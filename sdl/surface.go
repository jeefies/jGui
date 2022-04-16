/*
Package sdl (main file sdl.go)
This is the file of structure of SDL_Surface
Defines the apis to use the functions to edit the surface (serves by SDL)
*/
package sdl

/*
#include <SDL2/SDL.h>

int get_place(int ppr, int x, int y) {
	// ppr: pixels per row
	return y * ppr + x;
}

int j_border(SDL_Surface * sur, SDL_Rect * area, int width, Uint32 color) {
	Uint32 * pixels = (Uint32 *)sur->pixels;

	int ppr = sur->pitch / sur->format->BytesPerPixel;

	int x = area->x;
	int y = area->y;
	int w = area->w;
	int h = area->h;

	for (int i = 0; i < w; i++) {
		for (int j = 0; j < width; j++) {
			pixels[get_place(ppr, x + i, y + j)] = color;
		}
		for (int j = 0; j < width; j++) {
			pixels[get_place(ppr, x + i, y + h - j - 1)] = color;
		}
	}

	for (int i = width; i < h - width; i++) {
		for (int j = 0; j < width; j++) {
			pixels[get_place(ppr, x + j, y + i)] = color;
		}
		for (int j = 0; j < width; j++) {
			pixels[get_place(ppr, x + w - 1 - j, y + i)] = color;
		}
	}

	return 0;
}

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

func (sur * Surface) BlitPart(src * Surface, origin * Rect, into * Rect) error {
    if C.SDL_BlitSurface(src, origin, sur, into) != C.int(0) {
        return NewSDLError("Could not apply surface into surface")
    }
    return nil
}

// Clear means fill the surface with all black
func (sur *Surface) Clear() error {
    return sur.FillRect(nil, 0)
}

// ClearWith want a color (type Color) to fill all the surface
func (sur *Surface) ClearWith(color Color) error {
    return sur.FillRect(nil, color.MapA(sur))
}

// Size returns two integer (int, int) shows the width and the height of the surface
func (sur *Surface) Size() (int, int) {
    return int(sur.w), int(sur.h)
}

// Close returns nothing but free the surface. That means it should not be used anymore after calling this function
func (sur *Surface) Close() {
    C.SDL_FreeSurface(sur)
}

// Create a new surface like current surface
func CreateSurface(width, height int) (*Surface, error) {
	// With default RGBAmasks
    sur := C.SDL_CreateRGBSurface(0, C.int(width), C.int(height), 32, 0, 0, 0, 0)
    if sur == nil {
            return nil, NewSDLError("Could Create Surface")
    }
    return sur, nil
}

func (sur * Surface) ToTexture(ren * Renderer) (*Texture, error) {
	return ren.CreateTextureFromSurface(sur)
}

// Draw Old series see https://github.com/jeefies/jGui/blob/a8b1a0c95dd9638e6ef528a8ddac42c0665ac8b9/jgui/draw.go
func (sur * Surface) DrawBorder(area * Rect, width int, cl Color) {
	var w C.int = C.int(width)
	var color C.Uint32 = C.Uint32(cl.MapA(sur))
	C.j_border(sur, area, w, color)
}

func (text *Texture) Close() {
	C.SDL_DestroyTexture(text)
}
