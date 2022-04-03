package jgui

import sdl "jGui/sdl"

const (
	WIN_FULLSCREEN   = sdl.WIN_FULLSCREEN
	WIN_OPENGL       = sdl.WIN_OPENGL
	WIN_SHOWN        = sdl.WIN_SHOWN
	WIN_HIDDEN       = sdl.WIN_HIDDEN
	WIN_BORDERLESS   = sdl.WIN_BORDERLESS
	WIN_RESIZABLE    = sdl.WIN_RESIZABLE
	WIN_DEFAULT      = (WIN_SHOWN | WIN_RESIZABLE)
	// No complete

	WINPOS_CENTERED  = sdl.WINPOS_CENTERED
	WINPOS_UNDIFINED = sdl.WINPOS_UNDIFINED
)

const (
	RENDERER_SOFTWARE      = sdl.RENDERER_SOFTWARE
	RENDERER_TARGETTEXTURE = sdl.RENDERER_TARGETTEXTURE
	RENDERER_PRESENTVSYNC  = sdl.RENDERER_PRESENTVSYNC
	RENDERER_ACCELERATED   = sdl.RENDERER_ACCELERATED
)

const (
	// format in vim, `gaip `, `gaip=`
	QUIT                uint32 = sdl.QUIT
	KEYUP               uint32 = sdl.KEYUP
	KEYDOWN             uint32 = sdl.KEYDOWN
	WINDOWS_EVENT       uint32 = sdl.WINDOWS_EVENT

	// Type of window events
	WINDOW_SHOWN        uint32 = sdl.WINDOW_SHOWN
	WINDOW_HIDDEN       uint32 = sdl.WINDOW_HIDDEN
	WINDOW_EXPOSED      uint32 = sdl.WINDOW_EXPOSED
	WINDOW_MOVED        uint32 = sdl.WINDOW_MOVED
	WINDOW_RESIZED      uint32 = sdl.WINDOW_RESIZED
	WINDOW_SIZE_CHANGED uint32 = sdl.WINDOW_SIZE_CHANGED
	WINDOW_MINIMIZED    uint32 = sdl.WINDOW_MINIMIZED
	WINDOW_MAXIMIZED    uint32 = sdl.WINDOW_MAXIMIZED
	WINDOW_RESTORED     uint32 = sdl.WINDOW_RESTORED
	WINDOW_ENTER        uint32 = sdl.WINDOW_ENTER
	WINDOW_LEAVE        uint32 = sdl.WINDOW_LEAVE
	WINDOW_FOCUS_GAINED uint32 = sdl.WINDOW_FOCUS_GAINED
	WINDOW_FOCUS_LOST   uint32 = sdl.WINDOW_FOCUS_LOST
	WINDOW_CLOSE        uint32 = sdl.WINDOW_CLOSE

    // Type of MOUSE EVENT
    MOUSE_MOTION        uint32 = sdl.MOUSE_MOTION
    MOUSE_DOWN          uint32 = sdl.MOUSE_DOWN
    MOUSE_UP            uint32 = sdl.MOUSE_UP
    MOUSE_WHEEL         uint32 = sdl.MOUSE_WHEEL

	// Special keys
	KLEFT               uint32 = sdl.KLEFT
	KRIGHT              uint32 = sdl.KRIGHT
	KUP                 uint32 = sdl.KUP
	KDOWN               uint32 = sdl.KDOWN

	KLSHIFT             uint32 = sdl.KLSHIFT
	KRSHIFT             uint32 = sdl.KRSHIFT
	KLCTRL              uint32 = sdl.KLCTRL
	KRCTRL              uint32 = sdl.KRCTRL

	KTAB                uint32 = sdl.KTAB
	KESC                uint32 = sdl.KESC
	KBACKSPACE          uint32 = sdl.KBACKSPACE
	KDEL                uint32 = sdl.KDEL
)
