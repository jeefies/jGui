#include <stdint.h>

typedef uint32_t Uint32;
typedef uint32_t j_window;

typedef uint32_t j_widget;
typedef uint32_t j_label;
typedef uint32_t j_button;

typedef void (*j_event)(j_widget);

extern j_window create_window(char * title, int w, int h, Uint32 flags);
extern j_window get_window_byid(Uint32 id);
extern j_window close_window(j_window id);

extern void mainloop();
extern void delay(Uint32 ms);
extern void j_quit();

extern void regist_event(j_widget wg, char * evtname, j_event);

extern void pack_label(j_label wg, j_window win);
extern void pack_button(j_button wg, j_window win);
extern void pack_widget(j_widget wg, j_window win);

extern j_label create_label(int x, int y, int w, int h, char * text);
extern j_button create_button(int x, int y, int w, int h, char * text);
