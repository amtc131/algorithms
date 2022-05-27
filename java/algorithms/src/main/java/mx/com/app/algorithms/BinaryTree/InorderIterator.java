package mx.com.app.algorithms.BinaryTree;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

public class InorderIterator {
    // Initializing the stack
    Stack<BinaryTreeNode> stk = new Stack<BinaryTreeNode>();

    // When our iterator is initialized, we populate our stack with all nodes from
    // the root till the left-most node
    public InorderIterator(BinaryTreeNode root) {
        populateStack(root);
    }


    // Function to populate the stack from the root till the left-most node
    public void populateStack(BinaryTreeNode root) {
        while (root != null) {
            // We keep pushing node into the stack until we reach the left-most node
            stk.push(root);
            root = root.left;
        }
    }

    // This function checks if there is a node next in line inside the iterator
    public boolean hasNext() {
        // If the stack is empty, it means that there is no node next in line
        return !stk.isEmpty();
    }

    // getNext returns null if there are no more elements in tree
    public BinaryTreeNode getNext() {
        // Return null if there's no succeeding node to return
        if (stk.isEmpty())
            return null;

        // Extracting and popping the next node from the stack
        BinaryTreeNode rVal = stk.pop();

        // We now populate the stack again from root till the left-most node in the
        // sub-tree of the right child of the node we just extracted
        BinaryTreeNode temp = rVal.right;
        populateStack(temp);

        // Returning the next node
        return rVal;
    }

    // Iterator helper function (Don't modify it)
    // This function returns the in-order list of nodes using the hasNext() and
    // getNext() methods
    public static String inorderUsingIterator(BinaryTreeNode root) {
        InorderIterator iter = new InorderIterator(root);
        String result = "";
        while (iter.hasNext()) {
            result += iter.getNext().data;
            if (iter.hasNext()) {
                result += ", ";
            }
        }
        if (result == "") {
            result = "null";
        } else if (result.contains(", ")) {
            result.substring(0, result.length() - 2);
        }
        return result;
    }

    public static void main(String[] args) {
        // Creating a binary tree
        List<Integer> input1 = new ArrayList<Integer>();
        input1.add(100);
        input1.add(50);
        input1.add(200);
        input1.add(25);
        input1.add(75);
        input1.add(125);
        input1.add(300);
        input1.add(12);
        input1.add(35);
        input1.add(60);
        BinaryTree tree1 = new BinaryTree(input1);

        // Creating a right degenerate binary tree
        List<Integer> input2 = new ArrayList<Integer>();
        input2.add(100);
        input2.add(50);
        input2.add(200);
        input2.add(25);
        input2.add(75);
        input2.add(125);
        input2.add(300);
        input2.add(12);
        input2.add(35);
        input2.add(60);
        Collections.sort(input2);
        BinaryTree tree2 = new BinaryTree(input2);

        // Creating a left degenerate binary tree
        List<Integer> input3 = new ArrayList<Integer>();
        input3.add(100);
        input3.add(50);
        input3.add(200);
        input3.add(25);
        input3.add(75);
        input3.add(125);
        input3.add(300);
        input3.add(12);
        input3.add(35);
        input3.add(60);
        Collections.sort(input3, Collections.reverseOrder());
        BinaryTree tree3 = new BinaryTree(input3);

        // Creating a single node binary tree
        BinaryTree tree4 = new BinaryTree(100);

        // Defining test cases
        BinaryTreeNode[] testCaseRoots = {tree1.root, tree2.root, tree3.root, tree4.root, null};
        String[] testCaseStatements = {"In-Order Traversal of a normal binary search tree: ",
                "In-Order Traversal of a right degenerate binary search tree: ",
                "In-Order Traversal of a left degenerate binary search tree: ",
                "In-Order Traversal of a single node binary tree: ",
                "In-Order Traversal of a null tree: "};

        for (int i = 0; i < testCaseRoots.length; i++) {
            if (i > 0) {
                System.out.print("\n");
            }
            System.out.println((i + 1) + ".\tBinary Tree:");
            System.out.println("\n\t" + testCaseStatements[i]);
            
			// Printing the in-order list using the method we just implemented
            System.out.println("\t" + inorderUsingIterator(testCaseRoots[i]));
            System.out.print(
                    "-------------------------------------------------------------------------------------------------------------------------------\n");
        }
    }
}
