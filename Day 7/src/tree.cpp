#include "main.h"

namespace nodeType
{
    string toString(NodeType type)
    {
        string str;

        switch (type)
        {
        case FILE:
            str = string("file");
            break;

        case DIR:
            str = string("dir");
            break;
        default:
            error("Unhandled Node Type");
        }

        return str;
    }
}

pair<string, string> splitByCharacter(string line, char c)
{
    stringstream ss(line);
    pair<string, string> p;

    getline(ss, p.first, c);
    getline(ss, p.second, c);

    return p;
}

pair<string, string> parseCommand(string str)
{
    if (str[0] != '$')
        error("Not a command: " + str);

    stringstream ss(str);
    pair<string, string> p;

    getline(ss, p.first, ' '); // skip $
    getline(ss, p.first, ' ');
    getline(ss, p.second, ' ');

    return p;
}

Node *createChildFile(Node *parent, string name, int size)
{
    if (parent == NULL)
        error("A file must always have a parent directory");

    Node *node = new Node();
    node->type = nodeType::FILE;
    node->name = name;
    node->size = size;
    node->parent = parent;
    parent->children.push_back(node);

    return node;
}

Node *createOrFindChildDirectory(Node *parent, string name)
{
    Node *node = NULL;

    if (parent != NULL)
        for (auto child : parent->children)
            if (child->name == name)
            {
                node = child;
                break;
            }

    if (node == NULL)
    {
        node = new Node();
        node->type = nodeType::DIR;
        node->size = 0;
        node->name = name;
        node->parent = parent;

        if (parent != NULL)
            parent->children.push_back(node);
    }

    return node;
}

void updateNodeSize(Node *node)
{
    if (node->type != nodeType::DIR)
        return;

    int size = 0;
    for (auto child : node->children)
    {
        updateNodeSize(child);
        size += child->size;
    }
    node->size = size;
}

Node *handleCdCommand(Node *root, Node *node, string directory)
{
    if (directory == "/")
        return root;

    if (directory == "..")
        return node->parent;

    return createOrFindChildDirectory(node, directory);
}

void handleLsCommand(Node *parent, list<string> &lines)
{
    while (!lines.empty())
    {
        string line = lines.front();

        if (line[0] == '$')
            break;
        else
            lines.pop_front();

        pair<string, string> p = splitByCharacter(line, ' ');

        if (isdigit(line[0]))
            createChildFile(parent, p.second, stoi(p.first));
        else
            createOrFindChildDirectory(parent, p.second);
    }
}

Node *parseFileSystemFromOutputLines(list<string> &lines)
{
    Node *rootNode = createOrFindChildDirectory(NULL, "/");
    Node *currNode = rootNode;

    while (!lines.empty())
    {
        string line = lines.front();
        lines.pop_front();
        pair<string, string> cmd = parseCommand(line);

        if (cmd.first == "cd")
            currNode = handleCdCommand(rootNode, currNode, cmd.second);
        else if (cmd.first == "ls")
            handleLsCommand(currNode, lines);
    }

    updateNodeSize(rootNode);
    return rootNode;
}

void cleanUpFileSystem(Node *node)
{
    for (Node *child : node->children)
        cleanUpFileSystem(child);

    delete node;
}

void printTree(Node *node, string layer)
{
    cout << layer << "- " << node->name << " (" << nodeType::toString(node->type) << ", size=" << node->size << ")" << endl;

    for (auto child : node->children)
        printTree(child, layer + "  ");
}

void printTree(Node *node)
{
    printTree(node, "");
}
