#include <SDL2/SDL.h>
#include <SDL2/SDL2_gfxPrimitives.h>

#define TRUE 1
#define FALSE 0
#define WIN_BORDERLESS SDL_WINDOW_BORDERLESS
#define INIT_VIDEO SDL_INIT_VIDEO

struct _Window {
	SDL_Window * win;
	SDL_Surface * scr;
	SDL_Renderer * ren;
	short id;
};
typedef struct _Window Window;

struct _winlist {
	Window ** wins;
	Uint32 size;
};

struct Point {
	int x;
	int y;
};
typedef struct Point Point;

#ifndef WINLIST_ON
#define WINLIST_ON
struct _winlist winlist;
#else
extern struct _winlist winlist;
#endif

#ifndef CHECK_FUNCS
#define CHEKC_FUNCS
int check(long r, int sdl_error) {
	if (r) {
		if (sdl_error) {
			fprintf(stderr, "Error: %s\n", SDL_GetError());
		}
		return TRUE;
	}
	return FALSE;
}

int check_p(void * point, int sdl_error) {
	if (point == NULL) {
		if (sdl_error) {
			fprintf(stderr, "Error: %s\n", SDL_GetError());
		}
		return TRUE;
	}
	return FALSE;
}
#else
int check(long r, int sdl_error);
int check_p(void * point, int sdl_error);
#endif

// General Use
int j_init(Uint32 flags);
int j_mainloop();

// For Generating Windows
Window * j_create_window(char * title, int w, int h, Uint32 flags);
void j_close_window(Window * win);
void j_append_win(Window * win);

// For drawing some specific shapes
// Rects
int j_drawrect_centered(SDL_Surface *, int, int, int, Uint32);
int j_drawrect_centered2(SDL_Surface *, int, int, int, Uint32);

// Circles
int j_draw_circle(SDL_Surface * sur, int cx, int cy, int r);
