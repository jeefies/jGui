package jgui

/*
Package jgui.Widgets
A file for all widgets like Button or Label
*/

import (
    _ "fmt"

    "sdl"
    "sdl/ttf"
)

type Color = uint32

// The interface of all Widgets
type Widgets interface {
    // Check if mouse is in the widget
    IsIn(x, y int) bool
    // regist the event like active or deactive.
    RegistEvent(evtName string, f func(Widgets))
    // Jgui would call this func when a event found
    Call(evtName string)
    // A Function to get the parent window
    Pack(win * Window)
    // state can be define as your own, but 0 always stands for default state
    Draw(state int)
}

type Button struct {
    childs []Widgets
    parent Widgets

    x, y, w, h int
    bg, fg Color

    win_id uint32
    events map[string]func(Widgets)

    text string
    text_size int
    text_color Color
    active_bg, active_fg Color

    // event's functions
    onclick func(*Button)
}

var default_font * ttf.Font

func init() {
    var err error
    default_font, err = ttf.OpenFont("font.ttf", 14)
    if (err != nil) {
        panic(err)
    }
}

func (btn * Button) IsIn(x, y int) bool {
    // Check if it's not in
    if (btn.x >= x) || (x >= btn.x + btn.w) || (btn.y >= y) || (y >= btn.y + btn.h) {
        return false
    } else {
        return true
    }
}

func (btn * Button) RegistEvent(evtName string, f func(Widgets)) {
    btn.events[evtName] = f
}

func (btn * Button) Call(evtName string) {
    if f, ok := btn.events[evtName]; ok {
        // TODO: Call Event Here
        f(btn)
    }
}

func (btn * Button) Pack(win * Window) {
    btn.win_id = win.id
    win.childs = append(win.childs, btn)
    btn.Draw(0)
}

func (btn * Button) Draw(state int) {
    var err error
    textsur, err := default_font.RenderText(btn.text, ToSDLColor(btn.text_color))
    check(err)
    win := GetWinById(btn.win_id)

    tw, th := textsur.Size()
    sw, sh := max(tw, btn.w), max(th, btn.h)

    var fg, bg Color
    if state == 0 { // Deactive
        fg, bg = btn.fg, btn.bg
    } else {
        fg, bg = btn.active_fg, btn.active_bg
    }

    win.DrawRect(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, fg)
    win.DrawRectBorder(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, 2, bg)

    f := func(n, a, b int) int {
        return n + (b - a) / 2
    }

    r := sdl.NewRect(f(btn.x, tw, sw) + 2, f(btn.y, th, sh) + 2, 0, 0)
    // fmt.Printf("rect: %v ; btn: x, y = %d, %d ; sw, sh = %d %d; tw, th = %d, %d\n", r, btn.x, btn.y, sw, sh, tw, th)
    win._scr.Blit(textsur, r)
}


func NewButton(x, y, w, h int, text string) (*Button) {
    btn := new(Button)
    btn.x, btn.y, btn.w, btn.h = x, y, w, h
    btn.text = text
    btn.bg = 0xffffff
    btn.fg = 0x000000
    btn.active_bg = 0xffffff
    btn.active_fg = 0xff0000
    btn.text_color = 0xffffff
    btn.events = make(map[string]func(Widgets))
    btn.events["active"] = func(Widgets) {
        btn.Draw(1)
    }
    btn.events["deactive"] = func(Widgets) {
        btn.Draw(0)
    }
    return btn
}

func ToRGB(c uint32) (r, g, b uint8) {
    r = uint8((c | 0xff0000) >> 4)
    g = uint8((c | 0x00ff00) >> 2)
    b = uint8((c | 0x0000ff) >> 0)
    return
}

func ToSDLColor(c uint32) sdl.Color {
    r, g, b := ToRGB(c)
    return *sdl.NewColor(r, g, b)
}
