#ifndef MAIN_H
#define MAIN_H

#include <iostream>
#include <string>
#include <list>

#include "logger.h"
#include "tree.h"

using namespace std;

list<string> readOutputFileLines(string fileName);
int totalDirsMaxSizeN(Node *node, int n);
void dirsMinSizeN(Node *root, int n, list<Node*>& dirs);
#endif
