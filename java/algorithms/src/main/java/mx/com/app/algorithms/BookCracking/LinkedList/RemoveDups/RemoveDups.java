package mx.com.app.algorithms.BookCracking.LinkedList.RemoveDups;

import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList;
import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList.Node;

public class RemoveDups {

    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        list.add(10);
        list.add(12);
        list.add(11);
        list.add(11);
        list.add(12);
        list.add(11);
        list.add(10);
        System.out.println("Linked List before removing duplicates : \n ");
        list.printList(list.head);
        System.out.println();
        solution(list);
        System.out.println("Linked List after removing duplicates : \n ");
        list.printList(list.head);
    }

    public static void solution(LinkedList list) {
        Node current = list.head, temp = null;
        while (current != null && current.next != null) {
            temp = current;
            while (temp.next != null) {
                if (current.data == temp.next.data) {
                    temp.next = temp.next.next;
                    System.gc();
                } else {
                    temp = temp.next;
                }
            }
            current = current.next;
        }
    }
}
