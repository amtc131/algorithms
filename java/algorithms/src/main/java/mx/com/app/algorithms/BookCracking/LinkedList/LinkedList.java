package mx.com.app.algorithms.BookCracking.LinkedList;

/**
 * @author amcr
 *         Creating a Linked List
 *         The code below implements a very basic signgly linked list.
 */
public class LinkedList {

    public Node head;

    public class Node {
        public Node next = null;
        public int data;
    
        public Node() {}
    }
    

    public Node add(int data) {
        Node newNode = new Node();
        newNode.data = data;
        newNode.next = head;
        head = newNode;
        return head;
    }

    public void printList(Node n) {
       while (n != null) {
            System.out.print("{ " + n.data + " }");
            n = n.next;
        }
    }

    /**
     * Deleting a Node from a singly Linked List
     * Given a node current , we find the previous node prev and set prev.
     * next eqauls to current.next. If the list is doubly linked , we must also
     * update current.next to set current.next.prev
     * equals to current.prev
     */
    public Node deleteNode(Node node, int data) {
        if (node == null)
            return null;
        Node current = node;

        if (current.data == data) {
            return node.next; // moved head
        }

        while (current.next != null) {
            if (current.next.data == data) {
                current.next = current.next.next;
                return node; // head didn't change
            }
            current = current.next;
        }
        return node;
    }
}

