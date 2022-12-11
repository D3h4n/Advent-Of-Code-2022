CC = gcc
CFLAGS = -std=c11 -Wall -Werror -Wextra -Wpedantic -g

BIN = bin/
SRC = src/
INC = include/

SRC_FILES = $(wildcard $(SRC)*.c)
OBJS = $(patsubst $(SRC)%.c,$(BIN)%.o,$(SRC_FILES))

all: $(BIN)day4

clean:
	rm -f $(BIN)day4 $(BIN)*.o

$(BIN)day4: $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^ 

$(BIN)%.o: $(SRC)%.c
	$(CC) $(CFLAGS) -I$(INC) -c -o $@ $^
