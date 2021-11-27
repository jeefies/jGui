package jgui

import (
    "log"
    "os"

    "jGui/sdl"
)

type Screen = sdl.Surface
type Rect = sdl.Rect

var logger *log.Logger

var loggerFile *os.File

func init() {
	var err error
	loggerFile, err = os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0664)
	check(err)

	logger = log.New(loggerFile,  "Jgui logger:", log.Lmsgprefix | log.Ltime |  log.Lmicroseconds | log.Lshortfile)
}

type Window struct {
	win * sdl.Window
	ren * sdl.Renderer
	_scr * sdl.Surface // scr would change to private function

	childs []Widgets
	current_widget Widgets
	areas [](*Area)

	id uint32
	sid uint32
}

type JguiError struct {
	msg string
}

func (je *JguiError) Error() string {
	return je.msg
}

func NewError(msg string) (*JguiError) {
	return &JguiError{msg}
}
