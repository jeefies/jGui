/*
A file defines the api of jgui to c.
That means you can build the project into c-shared library and then use it in c
*/
package jgui

/*
#include <stdlib.h>
#include <stdint.h>
typedef uint32_t Uint32;

typedef uint32_t jgui_window;
typedef uint32_t j_widget;

typedef void (*j_event)(j_widget);
 */
import "C"

import "unsafe"

import "sdl"

//export create_window
func create_window(title * C.char, w, h C.int, flags C.Uint32) C.jgui_window {
	gs := C.GoString(title)
	return to_id(CreateWindow(gs, int(w), int(h), uint32(flags)))
}

//export get_window_byid
func get_window_byid(id C.Uint32) C.jgui_window {
	return id
}

//export close_window
func close_window(id C.jgui_window) {
	from_id(id).Close()
}

func from_id(id C.jgui_window) *Window {
	return GetWinById(uint32(id))
}

func to_id(win * Window) C.jgui_window {
	return C.jgui_window(win.id)
}

//export delay
func delay(ms C.Uint32) {
	sdl.Delay(uint32(ms));
}

//export create_label
func create_label(x, y, w, h C.int, text *C.char) C.j_widget {
	lb := NewLabel(int(x), int(y), int(w), int(h), C.GoString(text))
	return C.j_widget(lb.id)
}

//export create_button
func create_button(x, y, w, h C.int, text *C.char) C.j_widget {
	btn := NewButton(int(x), int(y), int(w), int(h), C.GoString(text))
	return C.j_widget(btn.id)
}

//export pack_button
func pack_button(wg C.j_widget, win C.Uint32) {
	widget := GetWidgetById(uint32(wg)).(*Button)
	widget.Pack(from_id(win));
}

//export pack_label
func pack_label(wg C.j_widget, win C.Uint32) {
	widget := GetWidgetById(uint32(wg)).(*Label)
	widget.Pack(from_id(win));
}

//export pack_widget
func pack_widget(wg C.j_widget, win C.Uint32) {
	widget := (*Widget)(unsafe.Pointer(uintptr(wg)))
	widget.Pack(from_id(win));
}

//export regist_event
func regist_event(wg C.j_widget, evtname * C.char, e func(C.j_widget)) {
    evt := func(w Widgets) {
        e(C.j_widget(w.GetId()))
    }
    widget := GetWidgetById(uint32(wg))
    widget.RegistEvent(C.GoString(evtname), evt)
}

//export mainloop
func mainloop() {
    Mainloop()
}

//export j_quit
func j_quit() {
	Quit()
}
