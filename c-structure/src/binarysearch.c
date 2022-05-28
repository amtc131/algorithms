#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef struct binarysearch binarysearch;
binarysearch *insert(binarysearch *root, int data);
struct binarysearch *newNode(int data);

struct binarysearch
{
    int data;
    struct binarysearch *left;
    struct binarysearch *right;
};

struct binarysearch *newNode(int data)
{
    struct binarysearch *node = (struct binarysearch *)malloc(sizeof(struct binarysearch));
    node->data = data;
    node->left = NULL;
    node->right = NULL;
    return (node);
}

binarysearch *insert(binarysearch *root, int data)
{

    if (root == NULL)
    {
        return newNode(data);
    }

    if (data < root->data)
    {
        root->left = insert(root->left, data);
    }
    else
    {
        root->right = insert(root->right, data);
    }

    return root;
}

void inOrder(binarysearch *root)
{
    if (root != NULL)
    {
        inOrder(root->left);
        printf("%d ", root->data);
        inOrder(root->right);
    }
}

int main(int argc, char const *argv[])
{
    binarysearch *tree = insert(NULL, 15);
    tree = insert(tree, 10);
    tree = insert(tree, 20);
    tree = insert(tree, 8);
    tree = insert(tree, 12);
    tree = insert(tree, 16);
    tree = insert(tree, 25);
    inOrder(tree);
    return 0;
}
