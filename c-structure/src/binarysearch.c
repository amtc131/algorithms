#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef struct binarysearch binarysearch;
binarysearch* insert(binarysearch* root, int data);

struct binarysearch
{
    int data;
    binarysearch* left;
    binarysearch* right;
};

binarysearch* insert(binarysearch* root, int data){
    binarysearch* temp = NULL;
    if(root == NULL){
        temp = (binarysearch *)malloc(sizeof(binarysearch));
        temp->data =data;
        root = temp;
    }

    if(root->left != NULL){
        root->left = insert(root->left, data);
    }else{
        root->right = insert(root->right, data);
    }

    return root;
}

void inOrder(binarysearch* root){
    if(root != NULL){
        inOrder(root->left);
        
        inOrder(root->right);
    }
}

int main(int argc, char const *argv[])
{
     binarysearch* tree = insert(NULL, 15);
     tree = insert(tree, 10);
    return 0;
}



