package mx.com.app.algorithms.BinaryTree;

import java.util.List;
import java.util.ArrayList;

public class MainBinaryTree{
    
    public static void main(String[] argv) {
        // Creating first binary tree with tree as tree1
        List<Integer> input1 = new ArrayList<Integer>();
        input1.add(100);
        input1.add(50);
        input1.add(200);
        input1.add(25);
        input1.add(125);
        input1.add(350);
        BinaryTree tree1 = new BinaryTree(input1);

        // Creating second binary tree with different values and structure,
        List<Integer> input2 = new ArrayList<Integer>();
        input2.add(4);
        input2.add(2);
        input2.add(6);
        input2.add(1);
        input2.add(5);
        input2.add(7);
        BinaryTree tree2 = new BinaryTree(input2);

        // Creating third binary tree with different structure and same values as
        // tree1
        List<Integer> input3 = new ArrayList<Integer>();
        input3.add(100);
        input3.add(25);
        input3.add(200);
        input3.add(50);
        input3.add(125);
        input3.add(350);
        BinaryTree tree1DiffLayout = new BinaryTree(input3);

        // Defining Test Cases
        BinaryTreeNode[] testCaseRoot1 = {tree1.root, tree1.root, tree1.root, tree1.root, null};
        BinaryTreeNode[] testCaseRoot2 = {tree1.root, tree2.root, tree1DiffLayout.root, null, null};

        for (int i = 0; i < testCaseRoot1.length; i++) {
            if (i > 0) {
                System.out.print("\n");
            }

            // Displaying level-order traversal of trees being tested
            System.out.println((i + 1) + ".\tFirst binary tree:  ");
            // TreePrint.displayTree(testCaseRoot1[i]);
            System.out.println("\n\tSecond binary tree: ");
            //TreePrint.displayTree(testCaseRoot2[i]);
            System.out.print("\n\tIdentical Tree: ");

            // Calling our areIdentical() function to check if tree are identical
            if (areIdentical2(testCaseRoot1[i], testCaseRoot2[i])) {
                System.out.print("true");
            } else {
                System.out.print("false");
            }
            System.out.print(
                    "\n----------------------------------------------------------------------------------------------------\n");
        }
    }
    
    public static boolean areIdentical(BinaryTreeNode root1, BinaryTreeNode root2) {
        // TODO: Write - Your - Code
        BinaryTreeNode temp1 = root1;
        BinaryTreeNode temp2 = root2;
        boolean result = true;

        if (root1 == null && root2 == null) {
            return true;
        }
        if (root1 != null && root2 != null) {
        
        while(temp1 != null && temp2 != null ){
            
            if(temp1.data == temp2.data){
                result = result && true;
            }else{
                result = result && false;
            }

            temp1 = temp1.left;
            temp2 = temp2.left;
            
        }
    }
        return  result;
    }

    public static boolean areIdentical2(BinaryTreeNode root1, BinaryTreeNode root2) {
        // TODO: Write - Your - Code
        
        if (root1 == null && root2 == null) {
            return true;
        }
        if (root1 != null && root2 != null) {
            
            return (areIdentical2(root1.left, root2.left) && areIdentical(root1.right, root2.right) && (root1.data == root2.data));
        
        }
        return  false;
    }

}
