package sdl

/*
#include<SDL2/SDL.h>
 */
import "C"

func (c Color) Map(sur * Surface) uint32 {
	var color MapColor = C.SDL_MapRGB(sur.format, C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B))
	return uint32(color)
}

func (c Color) MapA(sur * Surface) uint32 {
	var color MapColor = C.SDL_MapRGBA(sur.format, C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B), C.Uint8(c.A))
	return uint32(color)

}

func (c Color) SDLColor() _Color {
    return _Color{C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B), C.Uint8(c.A)}
}
