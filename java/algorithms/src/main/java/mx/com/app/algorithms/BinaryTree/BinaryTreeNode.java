package mx.com.app.algorithms.BinaryTree;

public class BinaryTreeNode {
    public int data;
    public BinaryTreeNode left;
    public BinaryTreeNode right;

    // below data members used only for some of the problems
    public BinaryTreeNode next;
    public BinaryTreeNode parent;
    public int count;

    public BinaryTreeNode(int data){
        this.data   = data;
        this.left   = null;
        this.right  = null;
        this.next   = null;
        this.parent = null;
        this.count = 0;
    }
    
}
