#include <jGui.h>


int j_init(Uint32 flags) {
	if (check(SDL_Init(SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_TIMER | flags), 1)) exit(1);
	winlist.wins = (Window **)malloc(sizeof(Window *));
	winlist.size = 0;
}

void j_append_win(Window * win) {
	winlist.size += 1;
	winlist.wins = (Window **)realloc(winlist.wins, sizeof(Window *) * winlist.size);
	win->id = winlist.size - 1;
}

Window * j_create_window(char * title, int w, int h, Uint32 flags) {
	Window * win = (Window *)malloc(sizeof(Window));
	j_append_win(win);

	if (check_p(win, 0)) {
		perror("Could not create windows");
		exit(1);
	}

	win->win = SDL_CreateWindow(title, SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED,
			w, h, SDL_WINDOW_SHOWN | flags);
	if (check_p(win->win, 1)) exit(1);

	win->ren = SDL_CreateRenderer(win->win, -1, SDL_RENDERER_PRESENTVSYNC);
	if (check_p(win->ren, 1)) exit(1);

	win->scr = SDL_GetWindowSurface(win->win);
	if (check_p(win->scr, 1)) exit(1);

	SDL_SetRenderDrawColor(win->ren, 255, 0, 0, 255);
	SDL_RenderFillRect(win->ren, NULL);
	SDL_RenderPresent(win->ren);
	
	return win;
}

void j_close_window(Window * win) {
	SDL_DestroyRenderer(win->ren);
	SDL_DestroyWindow(win->win);
}

Uint32 _timer_present_renderer_for_time(Uint32 interval, void * data) {
	for (int i = 0; i < winlist.size; i++) {
		SDL_RenderPresent(winlist.wins[i]->ren);
		printf("RenderPresent %d\n", i);
	}
	return interval;
}

int j_drawrect_centered(SDL_Surface * sur, int x, int y, int r, Uint32 color) {
	// Not Draw Circle, but a rectangle, x, y is the center of the rect, r is half of the width
	// if r is 1, so width would be 2r, notice it won't be odd!
	// Center would be the left down side by the real center
	Uint32 * pixels = (Uint32 *)sur->pixels;

	// printf("%p\n", sur->pixels);

	// printf("sizeof(pixels): %d, sizeof(*pixels): %d, sizeof(&pixels): %d\n",
	// 		sizeof(pixels), sizeof(*pixels), sizeof(&pixels));

	// printf("Get All pixels\n");

	int ppr = sur->pitch / sur->format->BytesPerPixel;

	for (int i = y - r; i < y + r; i++) {
		for (int j = x - r; j < x + r; j++) {
			pixels[i * ppr + j] = color;
			printf("set (%d, %d)\n", i, j);
		}
	}
}

int j_drawrect_centered2(SDL_Surface * sur, int x, int y, int r, Uint32 color) {
	// Not Draw Circle, but a rectangle, x, y is the center of the rect, r is half of the width
	// if r is 1, so width would be 2r + 1, notice it won't be even!
	
	Uint32 * pixels = (Uint32 *)sur->pixels;

	int ppr = sur->pitch / sur->format->BytesPerPixel; // Pixel per Row

	for (int i = y - r; i < y + r + 1; i++) {
		for (int j = x - r; j < x + r + 1; j++) {
			pixels[i * ppr + j] = color;
			printf("set (%d, %d)\n", i, j);
		}
	}
}

int draw_circle(SDL_Surface * sur, int cx, int cy, int r) {
	Uint32 * pixels = (Uint32 *) sur->pixels;

	int ppr = sur->pitch / sur->format->BytesPerPixel;

	Point p = {0, r}; // From left

	Point * ps = (Point *)malloc(sizeof(Point));
	int size = 1, count = 0, pr = r * r;
	int pxy; // Pow of pos

	while (p.y >= 0) {
		pxy = (r - p.x) * (r - p.x) + (r - p.y) * (r -p.y);
		printf("pxy is %d\n", pxy);

		if (pxy <= pr) {
			ps[count] = p;
			count += 1;
			size += 1;
			ps = (Point *)realloc(ps, sizeof(Point) * size);

			p.y -= 1;
			printf("so y - 1, now x, y = %d, %d\n", p.x, p.y);
		} else {
			p.x += 1;
			printf("so x + 1, now x, y = %d, %d\n", p.x, p.y);
		}
	}

	for (int i = 0; i < count; i++) {
		Point ip = ps[i];
		printf("x, y = %d, %d\n", ip.x, ip.y);

		for (int j = ip.x; j < r; j++) {
			pixels [cx - r + j + (cy - r + ip.y) * ppr] = 0xff0000;
		}
	}

	return 0;
}

int j_mainloop() {
	SDL_Event e;

	SDL_TimerID _timer_present_renderer_id = SDL_AddTimer(1000 / 24,_timer_present_renderer_for_time, NULL);

	while (TRUE) {
		while (SDL_PollEvent(&e)) {
			switch (e.type) {
				case SDL_QUIT:
					SDL_RemoveTimer(_timer_present_renderer_id);
					return 0;
			}
		}
		SDL_Delay(20);
	}

	return 1;
}

int main(int argc, char * argv[]) {
	j_init(0);
	Window * win = j_create_window("Hello J_Gui", 200, 200, WIN_BORDERLESS);


	j_mainloop();
__quit:
	j_close_window(win);
	SDL_Quit();
	return 0;
}
