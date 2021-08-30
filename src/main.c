#include <jGui.h>

int _win_update_use_surface = 0;

/* This General Use function. to check result or pointer is right or not */
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
/* End General Use functions */

int j_init(Uint32 flags) {
	if (check(SDL_Init(SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_TIMER | flags), 1)) exit(1);
	winlist.wins = (Window **)malloc(sizeof(Window *));
	winlist.size = 0;
}

void j_append_win(Window * win) {
	winlist.size += 1;
	winlist.wins = (Window **)realloc(winlist.wins, sizeof(Window *) * winlist.size);
	// win->id = winlist.size - 1;
	winlist.wins[(win->id = winlist.size - 1)] = win;
}

Window * j_create_window(char * title, int w, int h, Uint32 flags) {
	Window * win = (Window *)malloc(sizeof(Window));

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

	j_append_win(win);

	return win;
}

void j_close_window(Window * win) {
	SDL_DestroyRenderer(win->ren);
	SDL_DestroyWindow(win->win);
}

Uint32 _timer_present_renderer_for_time(Uint32 interval, void * data) {
	static SDL_Texture * text = NULL;
	for (int i = 0; i < winlist.size; i++) {
		if (_win_update_use_surface) {
			if (!text) printf("Update Window Surface\n");
			SDL_Renderer * ren = winlist.wins[i] -> ren;

			// j_draw_circle (winlist.wins[i]->scr, 100, 50, 40);

			

			if (!text) {
				text = SDL_CreateTextureFromSurface(ren, winlist.wins[i]->scr);
				while (!text) {
					text = SDL_CreateTextureFromSurface(ren, winlist.wins[i]->scr);
					printf("... %s\n", SDL_GetError());
				}
			}

			SDL_RenderCopy(ren, text, NULL, NULL);
			SDL_RenderPresent(ren);

			SDL_DestroyTexture(text);
		} else {
			SDL_RenderPresent(winlist.wins[i]->ren);
		}
	}
	return interval;
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

int j_update_method(int flags) {
	_win_update_use_surface = flags;

	/*
	if (_win_update_use_surface) {
		for (int i = 0; i < winlist.size; i++) {
			SDL_FillRect(winlist.wins[i]->scr, NULL, 0);
		}
	}
	*/
	return flags;
}

int main(int argc, char * argv[]) {
	j_init(0);
	Window * win = j_create_window("Hello J_Gui", 200, 200, WIN_BORDERLESS);

	j_update_method(UPDATE_BY_SURFACE);

	SDL_Rect r = {1, 1, 5, 5};
	SDL_FillRect(win->scr, &r, 0xff00ff);
	j_draw_circle (win->scr, 100, 100, 100, 0xff0000);

	printf("Mainloop start!");

	j_mainloop();
__quit:
	j_close_window(win);
	SDL_Quit();
	return 0;
}
