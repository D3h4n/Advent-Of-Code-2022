#ifndef PART1_H
#define PART1_H

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "vector.h"

typedef struct
{
    int start;
    int end;
    int size;
} Range;

int part1(const char *filename);
void readPairs(const char *filename, Vector vector);
int countFullyOverlappingPairs(Vector vector);
void parseRangePairs(char *src, Range *range1, Range *range2);

#endif
