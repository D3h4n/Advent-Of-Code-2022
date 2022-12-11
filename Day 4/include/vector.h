#ifndef VECTOR_H
#define VECTOR_H

#include <stdlib.h>

typedef struct _Node_struct
{
    void *data;
    struct _Node_struct *next;
} _Node;

typedef struct _Vector_struct
{
    _Node *head;
    size_t size;
    void (*push)(struct _Vector_struct*, void *);
    void *(*pop)(struct _Vector_struct*);
} _Vector, *Vector;

Vector newVector();
void freeVector(Vector vec);
#endif
