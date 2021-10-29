#include <stdio.h>
#include <stdlib.h>
#include <jgui.h>

void btn_onclick(j_widget wg) {
    printf("Button Onclick\n");
}

int main() {
	j_window win = create_window("Test jGui C api", 100, 100, 0);

	j_label lb = create_label(10, 10, 0, 0, "Test Label");
	pack_label(lb, win);
    
    j_button btn = create_button(10, 30, 0, 0, "Test Button");
    regist_event(btn, "mouse up", btn_onclick);
    pack_button(btn, win);
	// pack_widget(lb, win);

	mainloop();

	// close_window(win);
}
