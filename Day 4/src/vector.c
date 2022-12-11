#include "vector.h"

void _push(_Vector* vec, void *data)
{
    _Node *node = malloc(sizeof(_Node));
    node->data = data;
    node->next = NULL;

    if (vec->size == 0)
    {
        vec->head = node;
    }
    else
    {
        node->next = vec->head;
        vec->head = node;
    }

    vec->size++;
}

void *_pop(_Vector* vec)
{
    if (vec->size < 1)
        return NULL;

    _Node *node = vec->head;
    vec->head = node->next;
    vec->size--;

    void *data = node->data;
    free(node);
    return data;
}


Vector newVector()
{
    _Vector *vec = malloc(sizeof(_Vector));
    
    vec->push = &_push;
    vec->pop = &_pop;

    vec->head = NULL;
    vec->size = 0;
    return vec;
}

void freeVector(_Vector* vec)
{
    void *data;

    while (vec->size > 0)
    {
        data = _pop(vec);
        free(data);
    }

    free(vec);
}
