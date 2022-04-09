package sdl

/*
#include<SDL2/SDL.h>

Uint32 getEtype(SDL_Event * e) {
	return e->type;
}

Uint32 getKey(SDL_Event * e) {
    return e->key.keysym.sym;
}


// It's unable to use unicode in SDL2 (but can if in SDL1.2)
// Uint32  getKeyUCode(SDL_Event * e) {
//      return e->key.keysym.unicode;
// }

Uint32 getWinId(SDL_Event * e) {
	return e->window.windowID;
}

Uint32 getWinEType(SDL_Event * e) {
	return e->window.event;
}

void getMousePositon(SDL_Event * e, int * x, int * y) {
    *x = e->motion.x;
    *y = e->motion.y;
}

unsigned char getButton(SDL_Event * e) {
	return e->button.button;
}

char * getInputText(SDL_Event * e) {
	return e->text.text;
}
*/
import "C"

// const SDL Event
const (
	// format in vim, `gaip `, `gaip=`
	QUIT                uint32 = uint32(C.SDL_QUIT)
	KEYUP               uint32 = uint32(C.SDL_KEYUP)
	KEYDOWN             uint32 = uint32(C.SDL_KEYDOWN)
	WINDOWS_EVENT       uint32 = uint32(C.SDL_WINDOWEVENT)
	TEXT_EDITING        uint32 = uint32(C.SDL_TEXTEDITING)
	TEXT_INPUT          uint32 = uint32(C.SDL_TEXTINPUT)

	// Type of window events
	WINDOW_SHOWN        uint32 = uint32(C.SDL_WINDOWEVENT_SHOWN)
	WINDOW_HIDDEN       uint32 = uint32(C.SDL_WINDOWEVENT_HIDDEN)
	WINDOW_EXPOSED      uint32 = uint32(C.SDL_WINDOWEVENT_EXPOSED)
	WINDOW_MOVED        uint32 = uint32(C.SDL_WINDOWEVENT_MOVED)
	WINDOW_RESIZED      uint32 = uint32(C.SDL_WINDOWEVENT_RESIZED)
	WINDOW_SIZE_CHANGED uint32 = uint32(C.SDL_WINDOWEVENT_SIZE_CHANGED)
	WINDOW_MINIMIZED    uint32 = uint32(C.SDL_WINDOWEVENT_MINIMIZED)
	WINDOW_MAXIMIZED    uint32 = uint32(C.SDL_WINDOWEVENT_MAXIMIZED)
	WINDOW_RESTORED     uint32 = uint32(C.SDL_WINDOWEVENT_RESTORED)
	WINDOW_ENTER        uint32 = uint32(C.SDL_WINDOWEVENT_ENTER)
	WINDOW_LEAVE        uint32 = uint32(C.SDL_WINDOWEVENT_LEAVE)
	WINDOW_FOCUS_GAINED uint32 = uint32(C.SDL_WINDOWEVENT_FOCUS_GAINED)
	WINDOW_FOCUS_LOST   uint32 = uint32(C.SDL_WINDOWEVENT_FOCUS_LOST)
	WINDOW_CLOSE        uint32 = uint32(C.SDL_WINDOWEVENT_CLOSE)

    // Type of MOUSE EVENT
    MOUSE_MOTION        uint32 = uint32(C.SDL_MOUSEMOTION)
    MOUSE_DOWN          uint32 = uint32(C.SDL_MOUSEBUTTONDOWN)
    MOUSE_UP            uint32 = uint32(C.SDL_MOUSEBUTTONUP)
    MOUSE_WHEEL         uint32 = uint32(C.SDL_MOUSEWHEEL)

	BUTTON_LEFT         uint8 = uint8(C.SDL_BUTTON_LEFT)
	BUTTON_MIDDLE       uint8 = uint8(C.SDL_BUTTON_MIDDLE)
	BUTTON_RIGHT        uint8 = uint8(C.SDL_BUTTON_RIGHT)

	// Special keys
	KLEFT               uint32 = uint32(C.SDLK_LEFT)
	KRIGHT              uint32 = uint32(C.SDLK_RIGHT)
	KUP                 uint32 = uint32(C.SDLK_UP)
	KDOWN               uint32 = uint32(C.SDLK_DOWN)

	KLSHIFT             uint32 = uint32(C.SDLK_LSHIFT)
	KRSHIFT             uint32 = uint32(C.SDLK_RSHIFT)
	KLCTRL              uint32 = uint32(C.SDLK_LCTRL)
	KRCTRL              uint32 = uint32(C.SDLK_RCTRL)

	KTAB                uint32 = uint32(C.SDLK_TAB)
	KESC                uint32 = uint32(C.SDLK_ESCAPE)
	KBACKSPACE          uint32 = uint32(C.SDLK_BACKSPACE)
	KDEL                uint32 = uint32(C.SDLK_DELETE)
)

var mx, my int

func init() {
    // enables unicode translation for event (only in SDL1.2)
    // C.SDL_EnableUNICODE(1)
}

// Create an Event instance (a pointer)
func NewEvent() (*Event) {
	return new(Event)
}

// Poll a event from sdl events list
// if true is returned, means the event is avaliable
func (e *Event) Poll() bool {
	if (C.SDL_PollEvent(e) != C.int(0)) {
		return true
	}

	if e.Type() == MOUSE_MOTION {
        mx, my = e.MousePosition()
    }

	return false
}

// It's a function like Poll but it would wait until a new event appear
func (e *Event) WaitPoll() {
    C.SDL_WaitEvent(e)
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

// KeyCode returns the translated unicode character by sdl (only in 1.2)
// func (e *Event) KeyCode() (rune) {
//     return rune(C.getKeyUCode())
// }

// WinId returns the windows id in sdl if the event type is window event
func (e *Event) WinId() (uint32) {
	return uint32(C.getWinId(e))
}

// The sub type of the window event
func (e *Event) WinEvent() (uint32) {
	return uint32(C.getWinEType(e))
}

// Returns the relative position of the window
func (e *Event) MousePosition() (int, int) {
    var x, y C.int
    C.getMousePositon(e, &x, &y)
    return int(x), int(y)
}
func (e *Event) GetButton() uint8 {
	return uint8(C.getButton(e))
}

func GetRootMousePosition() (int,  int) {
    var x, y C.int
	C.SDL_GetGlobalMouseState(&x, &y)
	return int(x), int(y)
}

func (e *Event) InputText() string {
	return C.GoString(C.getInputText(e))
}
