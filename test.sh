go build -buildmode=c-archive -o jgui.a main
gcc test/testcapi.c jgui.a -lSDL2 -lSDL2_ttf -lpthread -o TCapis -Iinclude
