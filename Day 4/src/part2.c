#include "part2.h"

int part2(char *filename)
{
    Vector vector = newVector();
    readPairs(filename, vector);

    int count = countOverlappingPairs(vector);
    
    freeVector(vector);
    return count;
}

int countOverlappingPairs(Vector vector)
{
    char *pair;
    Range firstRange, secondRange, temp;
    int count = 0;

    while ((pair = vector->pop(vector)))
    {
        parseRangePairs(pair, &firstRange, &secondRange);
        free(pair);

        if (secondRange.start <= firstRange.start)
        {
            temp = firstRange;
            firstRange = secondRange;
            secondRange = temp;
        }

        if (firstRange.end >= secondRange.start)
            count++;
    }

    return count;
}
