#include <jgui.h>

int main() {
	j_window win = create_window("Test jGui C api", 100, 100, 0);

	j_label lb = create_label(10, 10, 0, 0, "Test Label");

	pack_label(lb, win);
	// pack_widget(lb, win);

	mainloop();

	// close_window(win);
}
