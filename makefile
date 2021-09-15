CC = gcc

BIN = ./bin
SRC = ./src
INC = ./include

SDLFLAGS = `pkg-config SDL2_gfx sdl2 --cflags --libs`

all:
	$(CC) $(SRC)/main.c -o $(BIN)/main -I$(INC) $(SDLFLAGS)
