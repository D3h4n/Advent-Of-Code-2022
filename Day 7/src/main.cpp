#include "main.h"

bool compare_node_sizes(const Node *first, const Node *second)
{
    return first->size - second->size < 0;
}

int main(int argc, char **argv)
{
    if (argc < 2)
    {
        cout << "Format: " << argv[0] << " [InputFile]\n";
        return 1;
    }

    list<string> outputLines = readOutputFileLines(argv[1]);
    Node *root = parseFileSystemFromOutputLines(outputLines);

    // PART 1
    int sum = totalDirsMaxSizeN(root, 100000);
    cout << "Part 1: " << sum << endl;

    // PART 2
    list<Node *> dirs;
    int requiredSpace = 30000000 - (70000000 - root->size);

    dirsMinSizeN(root, requiredSpace, dirs);
    dirs.sort(compare_node_sizes);

    cout << "Part 2: " << dirs.front()->name << endl;

    cleanUpFileSystem(root);
    return 0;
}

list<string> readOutputFileLines(string fileName)
{
    ifstream f(fileName);
    string line;
    list<string> lines;

    while (getline(f, line))
    {
        lines.push_back(line);
    }

    f.close();
    return lines;
}

int totalDirsMaxSizeN(Node *root, int n)
{
    if (!root || root->type != nodeType::DIR)
        return 0;

    int total = 0;

    if (root->size <= n)
    {
        total += root->size;
    }

    for (auto child : root->children)
    {
        total += totalDirsMaxSizeN(child, n);
    }

    return total;
}

void dirsMinSizeN(Node *root, int n, list<Node *> &dirs)
{
    if (root == NULL || root->type != nodeType::DIR)
        return;

    for (auto child : root->children)
    {
        dirsMinSizeN(child, n, dirs);
    }

    if (root->size >= n)
    {
        dirs.push_back(root);
    }
}
