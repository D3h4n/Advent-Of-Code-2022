CC = gcc
CFLAGS = -std=c11 -Wall -Werror -Wextra -Wpedantic -g

BIN = bin/
SRC = src/
INC = include/

SRC_FILES = $(wildcard $(SRC)*.c)
OBJS = $(patsubst $(SRC)%.c,$(BIN)%.o,$(SRC_FILES))

all: $(BIN)day4

$(BIN)day4: $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^ 

$(BIN)%.o: $(SRC)%.c | $(BIN)
	$(CC) $(CFLAGS) -I$(INC) -c -o $@ $^

$(BIN):
	@mkdir -p $@

clean:
	rm -rf $(BIN)
