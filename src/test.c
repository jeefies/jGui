#include <SDL2/SDL.h>

#include <stdlib.h>
#define CT SDL_WINDOWPOS_CENTERED

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

int main() {
	SDL_Window * win;
	SDL_Surface * sur;
	win = SDL_CreateWindow("Test", CT, CT, 200, 200, SDL_WINDOW_SHOWN | SDL_WINDOW_BORDERLESS);
	sur = SDL_GetWindowSurface(win);

	// Uint32 * pixels = (Uint32 *)sur->pixels;
	draw_circle(sur, 100, 100, 50);
	SDL_UpdateWindowSurface(win);

	SDL_Delay(2000);

	return 0;
}

