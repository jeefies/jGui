package jgui

import sdl "jGui/sdl"

const ID_NULL ID = ^uint32(0)

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
	WIN_UPDATE_SURFACE = iota
	WIN_UPDATE_RENDER
)

const (
	RENDERER_SOFTWARE      = sdl.RENDERER_SOFTWARE
	RENDERER_TARGETTEXTURE = sdl.RENDERER_TARGETTEXTURE
	RENDERER_PRESENTVSYNC  = sdl.RENDERER_PRESENTVSYNC
	RENDERER_ACCELERATED   = sdl.RENDERER_ACCELERATED
)

type WidgetEvent = uint16
const (
	WE_INIT WidgetEvent = ^uint16(0)
	WE_IN WidgetEvent = iota
	WE_OUT
	WE_RESIZE
	WE_FOCUSIN
	WE_FOCUSOUT
	WE_CLICK
	WE_CLICK_DOUBLE
	WE_TEXT_EDITING
	WE_TEXT_INPUT
	WE_KEY
)

const (
	ALIGN_CENTER = 1 << iota
	ALIGN_LEFT
	ALIGN_RIGHT
	ALIGN_TOP
	ALIGN_BOTTOM
)

const (
	REL_X = 1 << iota
	REL_Y
	REL_W
	REL_H
	REL_FLAGS = 0b1111
)

const (
	// format in vim, `gaip `, `gaip=`
	QUIT                uint32 = sdl.QUIT
	KEYUP               uint32 = sdl.KEYUP
	KEYDOWN             uint32 = sdl.KEYDOWN
	WINDOWS_EVENT       uint32 = sdl.WINDOWS_EVENT
	TEXT_EDITING        uint32 = sdl.TEXT_EDITING
	TEXT_INPUT          uint32 = sdl.TEXT_INPUT

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

	BUTTON_LEFT         uint8 = sdl.BUTTON_LEFT
	BUTTON_MIDDLE       uint8 = sdl.BUTTON_MIDDLE
	BUTTON_RIGHT        uint8 = sdl.BUTTON_RIGHT

	// Special keys
	KLEFT               uint32 = sdl.KLEFT
	KRIGHT              uint32 = sdl.KRIGHT
	KUP                 uint32 = sdl.KUP
	KDOWN               uint32 = sdl.KDOWN

	KLSHIFT             uint32 = sdl.KLSHIFT
	KRSHIFT             uint32 = sdl.KRSHIFT
	KLCTRL              uint32 = sdl.KLCTRL
	KRCTRL              uint32 = sdl.KRCTRL
	KRALT               uint32 = sdl.KRALT
	KLALT               uint32 = sdl.KLALT

	KTAB                uint32 = sdl.KTAB
	KESC                uint32 = sdl.KESC
	KBACKSPACE          uint32 = sdl.KBACKSPACE
	KDEL                uint32 = sdl.KDEL

	KRETURN             uint32 = sdl.KRETURN
	KPAUSE              uint32 = sdl.KPAUSE
	KINSERT             uint32 = sdl.KINSERT
	KHOME               uint32 = sdl.KHOME
	KEND                uint32 = sdl.KEND
	KPAGEUP             uint32 = sdl.KPAGEUP
	KPAGEDOWN           uint32 = sdl.KPAGEDOWN
	KF1                 uint32 = sdl.KF1
	KF2                 uint32 = sdl.KF2
	KF3                 uint32 = sdl.KF3
	KF4                 uint32 = sdl.KF4
	KF5                 uint32 = sdl.KF5
	KF6                 uint32 = sdl.KF6
	KF7                 uint32 = sdl.KF7
	KF8                 uint32 = sdl.KF8
	KF9                 uint32 = sdl.KF9
	KF10                uint32 = sdl.KF10
	KF11                uint32 = sdl.KF11
	KF12                uint32 = sdl.KF12
)
