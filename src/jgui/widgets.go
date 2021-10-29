package jgui

/*
Package jgui.Widgets
A file for all widgets like Button or Label
*/

import (
    "sync"
    
    "sdl"
    "sdl/ttf"
)

var wgrefs struct {
    sync.Mutex
    left int
    wgs []Widgets
    avaliables []int
}

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
    // When 'deactive' event called, must Draw as the default state (state = 1)
    Draw(state int)
    // The close function
    Close()
    // A function to change the position of the widget
    MoveTo(x, y int)
    // A function to set the width and the height
    SetSize(w, h int)
    // A function to get the size of the widget
    GetSize() (int, int)
    
    GetId() uint32
}

type Widget struct {
    // childs []Widgets
    // parent Widgets

    x, y, w, h int
    bg, fg Color

    win_id uint32
    id uint32
    events map[string]func(Widgets)
}

type Label struct {
    *Widget
    text string
    // fg is the text color, default to white
    // bg is the background color, default to black
}


type Button struct {
    *Widget
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
    
    wgrefs.wgs = make([]Widgets,  0)
    wgrefs.avaliables = make([]int, 0)
    wgrefs.left = 0
}

func NewWidgetId(wg Widgets) uint32 {
    wgrefs.Lock()
    defer wgrefs.Unlock()
    var id int
    if wgrefs.left == 0 {
        id = len(wgrefs.wgs)
        wgrefs.wgs = append(wgrefs.wgs, wg)
    } else {
        id = wgrefs.avaliables[0]
        wgrefs.avaliables = wgrefs.avaliables[1:]
        wgrefs.wgs[id] = wg
    }
    
    return uint32(id)
}

func GetWidgetById(id uint32) Widgets {
    iid := int(id)
    if (iid > len(wgrefs.wgs)) {
        return nil
    }
    
    if wg := wgrefs.wgs[iid]; wg != nil {
        return wg
    }
    
    return nil
}

// to check it the position (x, y) is in the area of the widget
func (wg * Widget) IsIn(x, y int) bool {
    // Check if it's not in
    if (wg.x >= x) || (x >= wg.x + wg.w) || (wg.y >= y) || (y >= wg.y + wg.h) {
        return false
    } else {
        return true
    }
}

// A function to close the wigdet (may still hold the memory)
func (wg * Widget) Close() {
    wgrefs.Lock()
    defer wgrefs.Unlock()
    id := wg.id
    wgrefs.wgs[id] = nil
    wgrefs.left += 1
    wgrefs.avaliables = append(wgrefs.avaliables, int(id))
}

// A function to change the widgets position
func (wg * Widget) MoveTo(x, y int) {
    wg.x, wg.y = x,  y
}

// A function to set the widget's width and height
func (wg * Widget) SetSize(w, h int) {
    wg.w, wg.h = w, h
}

// A function returns to int of the widgets width and height
func (wg * Widget) GetSize() (int, int) {
    return wg.w, wg.h
}

// Regist a Event's function by the name (you can define your own event name)
func (wg * Widget) RegistEvent(evtName string, f func(Widgets)) {
    wg.events[evtName] = f
}

// A function to call the event function by the event name
func (wg * Widget) Call(evtName string) {
    if f, ok := wg.events[evtName]; ok {
        // TODO: Call Event Here
        f(wg)
    }
}

// A function returns the id of the widget
func (wg * Widget) GetId() uint32 {
    return wg.id
}

// A function to pack in the window
func (wg * Widget) Pack(win * Window) {
    wg.win_id = win.id
    win.childs = append(win.childs, wg)
    // BUG: this would not call the current widgets function Draw, but structure Widget's
    // 不会调用子类Draw方法，依据测试，这里会调用Widget结构体下的Draw方法，在log.txt中回出现错误信息
    // 导致在main函数中Pack时，屏幕不会显示。利用创建新Widget时函数闭包的性质，以及Call(不会继承)解决。
    // wg.Draw(0)
}

// A function should be overwrited. 需要被重写！
func (wg * Widget) Draw(state int) {
    logger.Print("Should not be called!\n")
    // Should be overwrite
}

// A function to draw the button according to the state number
// (you can define your own state number meaning, but 0 always should means the default or inactive state)
func (btn * Button) Draw(state int) {
    // Print("Button Drew\n")
    var err error
    textsur, err := default_font.RenderText(btn.text, ToSDLColor(btn.text_color))
    check(err)
    win := GetWinById(btn.win_id)

    tw, th := textsur.Size()
    sw, sh := max(tw, btn.w), max(th, btn.h)
    btn.w, btn.h = sw, sh

    var fg, bg Color
    if state == 0 { // Deactive
        fg, bg = btn.fg, btn.bg
    } else {
        fg, bg = btn.active_fg, btn.active_bg
    }

    win.DrawRect(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, bg)
    win.DrawRectBorder(btn.x, btn.y, btn.x + sw + 4, btn.y + sh + 4, 2, fg)

    f := func(n, a, b int) int {
        return n + (b - a) / 2
    }

    r := sdl.NewRect(f(btn.x, tw, sw) + 2, f(btn.y, th, sh) + 2, 0, 0)
    // fmt.Printf("rect: %v ; btn: x, y = %d, %d ; sw, sh = %d %d; tw, th = %d, %d\n", r, btn.x, btn.y, sw, sh, tw, th)
    win._scr.Blit(textsur, r)
}

// Set the Background Color of the button
func (btn * Button) SetBgColor(color Color) {
    btn.bg = color
}

// Set the Forground Color of the button (or the border color)
func (btn * Button) SetFgColor(color Color) {
    btn.fg = color
}

// Set the background Color of the button when mouse hover on it
func (btn * Button) SetActiveBgColor(color Color) {
    btn.active_bg = color
}

// Set the foreground color of the button (or the border color)
func (btn * Button) SetActiveFgColor(color Color) {
    btn.active_fg = color
}

// Returns a pointer to a Button instance, remember to use Pack after New...
// Button has two default event "active" and "deactive"
func NewButton(x, y, w, h int, text string) (*Button) {
    btn := &Button{Widget: new(Widget)}
    btn.id = NewWidgetId(btn)
    btn.x, btn.y, btn.w, btn.h = x, y, w, h
    btn.text = text
    btn.fg = 0xffffff
    btn.bg = 0x000000
    btn.active_fg = 0xffffff
    btn.active_bg = 0xff0000
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

// Returns a pointer to a Label instance
// Label has no default events
func NewLabel(x, y, w, h int, text string) (*Label) {
    lb := &Label{Widget: new(Widget), text: text}
    lb.id = NewWidgetId(lb)
    lb.x, lb.y, lb.w, lb.h = x, y, w, h
    lb.fg = 0xffffff
    lb.bg = 0x000000
    lb.events = make(map[string]func(Widgets))
    lb.events["deactive"] = func(Widgets) {
        lb.Draw(0)
    }
    return lb
}

func (lb * Label) Draw(state int) {
    // Print("Label Drew\n")
    if state == 0 {
        var err error
        textsur, err := default_font.RenderText(lb.text, ToSDLColor(lb.fg))
        check(err)

        tw, ty := textsur.Size()
        sw, sh := max(lb.w, tw), max(lb.h, ty)
        // FIXED: mistake update label's x, y to w, h
        lb.w, lb.h = sw, sh

        r := sdl.NewRect(lb.x, lb.y, sw, sh)
        win := GetWinById(lb.win_id)
        win.DrawRect(lb.x, lb.y, sw, sh, lb.bg)
        win._scr.Blit(textsur,  r)
    }
}

// Map uint32 to r, g, b color
func ToRGB(c uint32) (r, g, b uint8) {
    r = uint8((c | 0xff0000) >> 4)
    g = uint8((c | 0x00ff00) >> 2)
    b = uint8((c | 0x0000ff) >> 0)
    return
}

// Map r, b, g color to SDL_Color instance (define in sdl)
func ToSDLColor(c uint32) sdl.Color {
    r, g, b := ToRGB(c)
    return *sdl.NewColor(r, g, b)
}
