package jgui

import "sdl"

type Event struct {
    EventName string
    EventCode int
    SDLEvent * sdl.Event
}
