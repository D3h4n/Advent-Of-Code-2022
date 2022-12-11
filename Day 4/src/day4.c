#include "day4.h"

int main(int argc, char **argv)
{
    char *inputFile = "puzzle_input.txt";

    if (argc >= 2)
    {
        inputFile = argv[1];
    }

    int result = part1(inputFile);
    printf("Part 1: %d\n", result);

    result = part2(inputFile);
    printf("Part 2: %d\n", result);
    return 0;
}
