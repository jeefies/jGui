#include <SDL2/SDL.h>

struct Point {
	int x;
	int y;
};
typedef struct Point Point;

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
			// printf("set (%d, %d)\n", i, j);
		}
	}

	return 0;
}

int j_drawrect_centered2(SDL_Surface * sur, int x, int y, int r, Uint32 color) {
	// Not Draw Circle, but a rectangle, x, y is the center of the rect, r is half of the width
	// if r is 1, so width would be 2r + 1, notice it won't be even!
	
	Uint32 * pixels = (Uint32 *)sur->pixels;

	int ppr = sur->pitch / sur->format->BytesPerPixel; // Pixel per Row

	for (int i = y - r; i < y + r + 1; i++) {
		for (int j = x - r; j < x + r + 1; j++) {
			pixels[i * ppr + j] = color;
			// printf("set (%d, %d)\n", i, j);
		}
	}
}

int j_draw_circle(SDL_Surface * sur, int cx, int cy, int r, Uint32 color) {
	Uint32 * pixels = (Uint32 *) sur->pixels;

	int ppr = sur->pitch / sur->format->BytesPerPixel;

	Point p = {0, r}; // From left

	Point * ps = (Point *)malloc(sizeof(Point));
	int size = 1, count = 0, pr = r * r;
	int pxy; // Pow of pos

	while (p.y >= 0) {
		pxy = (r - p.x) * (r - p.x) + (r - p.y) * (r -p.y);
		// printf("pxy is %d\n", pxy);

		if (pxy <= pr) {
			ps[count] = p;
			count += 1;
			size += 1;
			ps = (Point *)realloc(ps, sizeof(Point) * size);

			p.y -= 1;
			// printf("so y - 1, now x, y = %d, %d\n", p.x, p.y);
		} else {
			p.x += 1;
			// printf("so x + 1, now x, y = %d, %d\n", p.x, p.y);
		}
	}

	for (int i = 0; i < count; i++) {
		Point ip = ps[i];
		// printf("x, y = %d, %d\n", ip.x, ip.y);

		for (int j = ip.x; j < r; j++) {
			// printf("%d %d %d %d\n", cx - r + j, cx + r - j, cy - r + ip.y, cy + r - ip.y - 1);
			pixels [cx - r + j + (cy - r + ip.y) * ppr] = color;
			pixels [cx + r - j - 1 + (cy - r + ip.y) * ppr] = color;

			pixels [cx - r + j + (cy + r - ip.y - 1) * ppr] = color;
			pixels [cx + r - j - 1 + (cy + r - ip.y - 1) * ppr] = color;
		}
	}

	return 0;
}
