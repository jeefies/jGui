/*
 * Draw patterns by setting each pixel.
 * Should bind with 'draw.go'
 * Every function starts with "p" returns (Point *)
 * The functions start with "d" is the actual functions to draw shapes
 */

#include <SDL2/SDL.h>


typedef struct {
	int x, y;
} Point;

typedef struct {
	int cap;
	int index;
	Point * pts;
} Points;

Points * new_ptlist(int size) {
	Points * pts = (Points *)malloc(sizeof(Points));
	pts->pts = (Point *)malloc(sizeof(Point) * size);
	pts->cap = size;
	pts->index = 0;
	return pts;
}

void ptl_append(Points * pts, int x, int y) {
	int i = pts->index;
	if (i >= pts->cap) {
		pts->cap += 5;
		pts->pts = (Point *)realloc(pts->pts, pts->cap * sizeof(Point));
	}
	
	pts->pts[i].x = x; pts->pts[i].y = y;
	pts->index += 1;
}

void ptl_free(Points * pts) {
	free(pts->pts);
	free(pts);
}

Point ptl_i(Points * pts, int i) {
	return pts->pts[i];
}

// Get Index
#define GI(x,y) (y)*ppr+(x)

// Get basic data
#define INIT Uint32 * pixels = (Uint32 *)sur->pixels;\
	int ppr = sur->pitch / sur->format->BytesPerPixel;

#define INIT_RECT int x = area->x;\
	int y = area->y;\
	int h = area->h;\
	int w = area->w;

// Fill Rect
#define fillRect(x,y,w,h) {\
	for (int i = 0; i < w; i++) {\
		for (int j = 0; j < h; j++) {\
			pixels[GI(x+i,y+j)] = color;\
		}\
	}\
}

void dBorder(SDL_Surface * sur, SDL_Rect * area, int width, Uint32 color) {
	INIT;
	INIT_RECT;

	// Top line
	fillRect(x + width, y, w - 2 * width, width);
	// Botton line
	fillRect(x + width, y + h - width, w - 2 * width, width);
	// Left linne
	fillRect(x, y + width, width, h - 2 * width);
	// Right Line
	fillRect(x + w - width, y + width, width, h - 2 * width);

	// Top left rect
	fillRect(x + width, y + width, width, width);
	// Top Right
	fillRect(x + w - 2 * width, y + width, width, width);
	// botton left
	fillRect(x + width, y + h - 2 * width, width, width);
	// botton right
	fillRect(x + w - 2 * width, y + h - 2 * width, width, width);
}

#define CIRCLE_POINT_DRAW pixels[GI(cx + x, cy + y)] = color;\
		pixels[GI(cx + x, cy - y)] = color;\
		pixels[GI(cx - x, cy + y)] = color;\
		pixels[GI(cx - x, cy - y)] = color

// Draw quater of the Circle
Points * pCircleQuater(int r) {
	int x = 0, y = r;
	int pr = r * r;

	Points * pt = new_ptlist(r * 2);

	for (int i = 0; i <= r; i++) {
		ptl_append(pt, x, y);

		while (x * x + y * y >= pr - r) {
			y--;
			if (y <= 0) break;
		}
		x++;
	}

	return pt;
}

void dCircleOBorder(SDL_Surface *sur, int cx, int cy, int r, Uint32 color) {
	INIT;
	int count;
	Points * pt = pCircleQuater(r);

	for (int i = 0; i < pt->index; i++) {
		int x = ptl_i(pt, i).x;
		int y = ptl_i(pt, i).y;
		CIRCLE_POINT_DRAW;
	}
	ptl_free(pt);
}

void dCircleBorder(SDL_Surface *sur, int cx, int cy, int r, int width, Uint32 color) {
	for (int i = 0; i < width; i++) {
		dCircleOBorder(sur, cx, cy, r - i, color);
	}
}

void dCircleFilled(SDL_Surface *sur, int cx, int cy, int r, Uint32 color) {
	INIT;
	Points * pt = pCircleQuater(r);

	for (int i = 0; i < pt->index; i++) {
		fillRect(cx + ptl_i(pt, i).x, cy + ptl_i(pt, i).y, 1, 2 * ptl_i(pt, i).y);
		fillRect(cx - ptl_i(pt, i).x, cy + ptl_i(pt, i).y, 1, 2 * ptl_i(pt, i).y);
	}

	ptl_free(pt);
}

double CalcK(int x1, int y1, int x2, int y2) {
	return ((double)y1 - (double)y2) / ((double)x1 - (double)x2); // 斜率
}

// Draw the Line towards right
Points * pOLineK(double k, int steps) {
	printf("k is %lf\n", k);

	Points * pt = new_ptlist(steps + 1);
	int nk = 0;
	if (k < 0) {
		nk = 1;
		k = -k;
	}

	// actual y
	double d = 0;
	int x = 0, y = 0;

	for (int i = 0; i <= steps; i++) {
		do {
			if (nk == 1)
				ptl_append(pt, x, -y);
			else
				ptl_append(pt, x, y);
			printf("d is %lf add x, y : %d, %d\n", d, x, y);
		} while (d - y > 0.5 && ++y);
		// 利用短路与特性实现递增

		x++;
		d += k;
	}
	return pt;
}

int one(int x) {
	return x>0?1:-1;
}

Points * pOLine(int x1, int y1, int x2, int y2, int * ct) {
	// Make sure x1 is always the left point
	// pOLineK need left point
	double k;
	if (x1 == x2) {
		k = -one(y1 - y2);
	} else
		// y1 can be equal to y2; hence k can be equal to 0
		k = ((double)y1 - (double)y2) / ((double)x1 - (double)x2); // 斜率
	int steps = x2 - x1;
	return pOLineK(k, steps);
}

void dOLine(SDL_Surface *sur, int x1, int y1, int x2, int y2, Uint32 color) {
	INIT;

	if (x1 > x2) {
		int tmp;
		tmp = x1;x1 = x2;x2 = tmp;
		tmp = y1;y1 = y2;y2 = tmp;
	}
	int steps;
	Points * pt = pOLine(x1, y1, x2, y2, &steps);

	for (int i = 0; i < pt->index; i++) {
		int x = ptl_i(pt, i).x, y = ptl_i(pt, i).y;
		printf("at %d %d, add x y = %d %d\n", x1 + x, y1 + y, x, y);
		pixels[GI(x1 + x, y1 + y)] = color;
	}

	free(pt);
}

void dLine(SDL_Surface * sur, int x1, int y1, int x2, int y2, int width, Uint32 color) {
	INIT;

	// Make sure x1 is always the left point
	if (x1 > x2) {
		int tmp;
		tmp = x1;x1 = x2;x2 = tmp;
		tmp = y1;y1 = y2;y2 = tmp;
	}


	double k = CalcK(x1, y1, x2, y2);
	double ok = -1.0 / k;
	printf("k %lf, ok %lf\n", k, ok);
	Points * pt_starts = pOLineK(ok, width);
	int x, y, ex, ey;
	for (int i = 0; i < pt_starts->index; i++) {
		x = x1 + pt_starts->pts[i].x;
		y = y1 + pt_starts->pts[i].y;
		ex = x2 + pt_starts->pts[i].x;
		ey = y2 + pt_starts->pts[i].y;

		printf("left %d %d, point at (%d, %d) to (%d, %d)\n", pt_starts->pts[i].x, pt_starts->pts[i].y, x, y, ex, ey);

		dOLine(sur, x, y, ex, ey, color);
	}
}
