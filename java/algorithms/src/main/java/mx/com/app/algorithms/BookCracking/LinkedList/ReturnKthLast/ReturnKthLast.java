package mx.com.app.algorithms.BookCracking.LinkedList.ReturnKthLast;

import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList;
import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList.Node;

public class ReturnKthLast {
    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        list.add(4);
        list.add(0);
        list.add(7);
        list.add(3);
        list.add(8);
        list.add(2);
        list.add(1);
        list.printList(list.head);
        System.out.println();
        int s=solution(list, 3);
        System.out.println("Element from the end is: "+s);
    }

    // Return Kht to List
    // Implement an algorithm to find the kth to last element of singly linkedList
    public static int solution(LinkedList list, int k){
        if(list.head == null){
            return 0;
        }
        Node current = list.head;
        while(k>0){
            current = current.next;
            k--;
        }
        Node temp = list.head;
        while (current != null) {
            current= current.next;
            temp = temp.next;
        }
        
        return temp.data;
    }
}
