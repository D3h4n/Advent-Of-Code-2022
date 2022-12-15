CXX=g++
CPPFLAGS=-Wall -Wextra -Wpedantic -g

BIN=bin/
INC=inc/
SRC=src/

TARGET = $(BIN)day7

SRC_FILES = $(wildcard $(SRC)*.cpp)
OBJS = $(patsubst $(SRC)%.cpp,$(BIN)%.o,$(SRC_FILES))


all: $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CPPFLAGS) -I$(INC) $^ -o $@

$(BIN)%.o: $(SRC)%.cpp | $(BIN)
	$(CXX) $(CPPFLAGS) -I$(INC) -c $< -o $@

$(BIN):
	@mkdir $(BIN)

clean:
	rm -rf $(BIN)
