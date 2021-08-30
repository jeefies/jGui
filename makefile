CC = gcc

BIN = ./bin
SRC = ./src
INC = ./include
BLD = ./build

SDLFLAGS = `pkg-config SDL2_gfx sdl2 --cflags --libs`
SDL_HEAD = `pkg-config sdl2 --cflags`

all: $(BLD)/main.o $(BLD)/draw.o
	$(CC) $(BLD)/main.o $(BLD)/draw.o -o $(BIN)/main $(SDLFLAGS)

$(BLD)/main.o: $(SRC)/main.c $(INC)/jGui.h
	$(CC) -c $(SRC)/main.c -o $(BLD)/main.o -I$(INC) $(SDL_HEAD)

$(BLD)/draw.o: $(SRC)/draw.c $(INC)/jGui.h
	$(CC) -c $(SRC)/draw.c -o $(BLD)/draw.o -I$(INC) $(SDL_HEAD)
