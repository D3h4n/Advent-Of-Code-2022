#include "part1.h"

int part1(const char *filename)
{
    Vector vector = newVector();
    readPairs(filename, vector);
    
    int count = countFullyOverlappingPairs(vector);

    freeVector(vector);
    return count;
}

void readPairs(const char *filename, Vector vector)
{
    FILE *file;
    char buffer[BUFSIZ];
    char *pair;
    size_t len;

    if (!(file = fopen(filename, "r")))
        return;

    while (!feof(file) && fgets(buffer, BUFSIZ, file))
    {
        len = strlen(buffer);
        pair = calloc(len, sizeof(char));
        strncpy(pair, buffer, len - 1);
        vector->push(vector, pair);
    }

    fclose(file);
}

int countFullyOverlappingPairs(Vector vector)
{
    char *pair;
    Range firstRange, secondRange, temp;
    int count = 0;

    while ((pair = vector->pop(vector)))
    {
        parseRangePairs(pair, &firstRange, &secondRange);
        free(pair);

        if (secondRange.size > firstRange.size)
        {
            temp = firstRange;
            firstRange = secondRange;
            secondRange = temp;
        }

        if (firstRange.start <= secondRange.start && firstRange.end >= secondRange.end)
        {
            count++;
        }
    }

    return count;
}

void parseRangePairs(char *src, Range *range1, Range *range2)
{
    sscanf(src, "%d-%d,%d-%d", &range1->start, &range1->end, &range2->start, &range2->end);
    range1->size = range1->end - range1->start;
    range2->size = range2->end - range2->start;
}
