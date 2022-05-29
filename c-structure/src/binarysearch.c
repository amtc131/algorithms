#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef struct binarysearch binarysearch;
binarysearch *insert(binarysearch *root, int data);
struct binarysearch *newNode(int data);
binarysearch *insertIterative(binarysearch *root, int data);

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

// Iterative function to insert a key into BST
// Root is pased by reference to the function
binarysearch *insertIterative(binarysearch *root, int data)
{
    // start with the root node
    binarysearch *curr = root;

    // pointer to store the parent of the current node
    binarysearch *parent = NULL;

    // if the tree is empty, create a new node and set it as root
    if (root == NULL)
    {
        return newNode(data);
    }

    // traverse the tree and find the parent node of the given key
    while (curr != NULL)
    {
        // update the parent to the current node
        parent = curr;

        // if the given key is less than the current node, go to the
        // left subtree; otherwise, go to the right subtree
        if (data < curr->data)
        {
            curr = curr->left;
        }
        else
        {
            curr = curr->right;
        }
    }

    // construct a node and assign it to the appropiate parent pointer
    if (data < parent->data)
    {
        parent->left = newNode(data);
    }
    else
    {
        parent->right = newNode(data);
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

void search(binarysearch *root, int data, binarysearch *parent)
{
    if (root == NULL)
    {
        printf("Key not found");
        return;
    }

    if (root->data == data)
    {
        if (parent == NULL)
        {
            printf("The node with key %d", data, " is root node \n");
        }
        else if (data < parent->data)
        {
            printf("The given key is the left node of the node with key %d \n", data);
        }
        else
        {
            printf("The given key is the right node of the node with key %d \n", parent->data);
        }
        return;
    }

    if (data < root->data)
    {
        search(root->left, data, root);
    }
    else
    {
        search(root->right, data, root);
    }
}

// Iterative function to search in a given BST
void searchIterative(binarysearch* root, int data){
    // start with the root node
    binarysearch* curr = root;

    // pointer to store the parent of the current node
    binarysearch* parent = NULL;

    // traverse the tree and search for the key
    while (curr != NULL && curr->data != data)
    {
        // Update the parent to the current node
        parent = curr;

        // if the given key is less than the current node, go to the left subtree;
        // otherwise, go to the right subtree
        if(data < curr->data)
        {
            curr = curr->left;
        }else
        {
            curr = curr->right;
        }
    }

    // if the key is not present in the key
    if (curr == NULL)
    {
        printf("Key not found");
        return;
    }
    if(parent == NULL)
    {
        printf("The node with key %d" , data , " is root node");
    }else if(data < parent->data )
    {
        printf("The given key is the left node of the node with key %d", parent->data);
    }else
    {
        printf("The given key is the right node of the node with key %d", parent->data);
    }
}

int main(int argc, char const *argv[])
{
    //           15     15 < 16 go to right
    //         /   \    subtree
    //        /     \
    //       10     20      20 > 16 go
    //      /  \    / \     to left subtree
    //     8   12  18 25
    //              \  \    18 > 16 go
    //              19 30    to left subtree
    //   Insert
    //    16 here
    //
    //    insert(root, 16)

    binarysearch *tree = NULL;
    tree = insertIterative(tree, 15);
    tree = insertIterative(tree, 10);
    tree = insertIterative(tree, 20);
    tree = insertIterative(tree, 8);
    tree = insertIterative(tree, 12);
    tree = insertIterative(tree, 16);
    tree = insertIterative(tree, 25);

    /*binarysearch *tree = NULL;
    tree = insert(tree, 15);
    tree = insert(tree, 10);
    tree = insert(tree, 20);
    tree = insert(tree, 8);
    tree = insert(tree, 12);
    tree = insert(tree, 16);
    tree = insert(tree, 25);
    */
    // Function to perform inorder traversal on the tree
    inOrder(tree);
    printf("\n");
    // Search
    search(tree, 25, NULL);
    // search iterative node
    searchIterative(tree, 25);

    return 0;
}
