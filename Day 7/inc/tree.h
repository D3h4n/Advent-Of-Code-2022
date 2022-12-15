#ifndef TREE_H
#define TREE_H

#include <iostream>
#include <fstream>
#include <bits/stdc++.h>
#include <string>
#include <unordered_map>
#include <list>
#include <ctype.h>
#include <utility> 

using namespace std;

namespace nodeType
{
    typedef enum NodeType
    {
        DIR,
        FILE
    } NodeType;

    string toString(NodeType type);
}

typedef struct Node
{
    struct Node *parent;
    string name;
    nodeType::NodeType type;
    int size;
    list<struct Node *> children;
} Node;

Node *parseFileSystemFromOutputLines(list<string>& lines);
void cleanUpFileSystem(Node* node);
void printTree(Node *node);
#endif
