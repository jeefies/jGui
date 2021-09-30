package sdl

/*
#include<SDL2/SDL.h>

Uint32 getEtype(SDL_Event * e) {
	return e->type;
}

Uint32 getKey(SDL_Event * e) {
    return e->key.keysym.sym;
}

Uint32 getWinId(SDL_Event * e) {
	return e->window.windowID;
}

Uint32 getWinEType(SDL_Event * e) {
	return e->window.event;
}
*/
import "C"

// const SDL Event
const (
	QUIT uint32 = uint32(C.SDL_QUIT)
	KEYUP uint32 = uint32(C.SDL_KEYUP)
	KEYDOWN uint32 = uint32(C.SDL_KEYDOWN)
	WINDOWS_EVENT uint32 = uint32(C.SDL_WINDOWEVENT)

	// Type of window events
	WINDOW_SHOWN uint32 = uint32(C.SDL_WINDOWEVENT_SHOWN)
	WINDOW_HIDDEN uint32 = uint32(C.SDL_WINDOWEVENT_HIDDEN)
	WINDOW_EXPOSED uint32 = uint32(C.SDL_WINDOWEVENT_EXPOSED)
	WINDOW_MOVED uint32 = uint32(C.SDL_WINDOWEVENT_MOVED)
	WINDOW_RESIZED uint32 = uint32(C.SDL_WINDOWEVENT_RESIZED)
	WINDOW_SIZE_CHANGED uint32 = uint32(C.SDL_WINDOWEVENT_SIZE_CHANGED)
	WINDOW_MINIMIZED uint32 = uint32(C.SDL_WINDOWEVENT_MINIMIZED)
	WINDOW_MAXIMIZED uint32 = uint32(C.SDL_WINDOWEVENT_MAXIMIZED)
	WINDOW_RESTORED uint32 = uint32(C.SDL_WINDOWEVENT_RESTORED)
	WINDOW_ENTER uint32 = uint32(C.SDL_WINDOWEVENT_ENTER)
	WINDOW_LEAVE uint32 = uint32(C.SDL_WINDOWEVENT_LEAVE)
	WINDOW_FOCUS_GAINED uint32 = uint32(C.SDL_WINDOWEVENT_FOCUS_GAINED)
	WINDOW_FOCUS_LOST uint32 = uint32(C.SDL_WINDOWEVENT_FOCUS_LOST)
	WINDOW_CLOSE uint32 = uint32(C.SDL_WINDOWEVENT_CLOSE)

	// Special keys
	KLEFT uint32 = uint32(C.SDLK_LEFT)
	KRIGHT uint32 = uint32(C.SDLK_RIGHT)
	KUP uint32 = uint32(C.SDLK_UP)
	KDOWN uint32 = uint32(C.SDLK_DOWN)

	KLSHIFT uint32 = uint32(C.SDLK_LSHIFT)
	KRSHIFT uint32 = uint32(C.SDLK_RSHIFT)
	KLCTRL uint32 = uint32(C.SDLK_LCTRL)
	KRCTRL uint32 = uint32(C.SDLK_RCTRL)

	KTAB uint32 = uint32(C.SDLK_TAB)
	KESC uint32 = uint32(C.SDLK_ESCAPE)
	KBACKSPACE uint32 = uint32(C.SDLK_BACKSPACE)
	KDEL uint32 = uint32(C.SDLK_DELETE)
)

// Create an Event instance (a pointer)
func NewEvent() (*Event) {
	return &Event{}
}

// Poll a event from sdl events list
// if true is returned, there's still another event waiting for you!
func (e *Event) Poll() bool {
	if (C.SDL_PollEvent(e) != C.int(0)) {
		return true
	}
	return false
}

// Type returns the type of the event.
// Constants in Prev content
func (e *Event) Type() (uint32) {
	return uint32(C.getEtype(e))
}

// Key returns the keysym if the event is keyboard event
func (e *Event) Key() (uint32) {
	return uint32(C.getKey(e))
}

// WinId returns the windows id in sdl if the event type is window event
func (e *Event) WinId() (uint32) {
	return uint32(C.getWinId(e))
}

// The sub type of the window event
func (e *Event) WinEvent() (uint32) {
	return uint32(C.getWinEType(e))
}