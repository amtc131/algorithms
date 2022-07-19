package mx.com.app.algorithms.BookCracking.LinkedList.ConvertBinaryNumber;

import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList;
import mx.com.app.algorithms.BookCracking.LinkedList.LinkedList.Node;

public class ConvertBinaryNumber {

  // Give head which is a reference node to a singly-linked list. The value of each
  // node in the linked list is either 0 or 1. The linked list holds the Binary 
  // representation of a number.
  // Return the decimal value of the number in the linked list.
  // The most significant bit is at the head of the linked list.
  // Example 1:
  // 
  // Input:  1 -> 0 -> 1
  // Output: 5
  // Explanation: (101) in base 2 = 5 in base 10

  public static void main(String[] args) {
        LinkedList list = new LinkedList();
        list.add(1);
        list.add(0);
        list.add(1);
        list.printList(list.head);
        int s = solution(list);
        System.out.println(s);
    }

   public static int solution(LinkedList l1){
        Node current = l1.head;
        int d = 0;
        while(current != null  ){
            d = (d * 2) + current.data;
            current = current.next;
        }        
        return d;
    }

}

    // Binary To Decimal Conversion Using Doubling Method
    // we consider the double of the previous digit as 0. 
    // For example, in 101101 2, the left-most digit is '1'. 
    // The double of the previous number is 0. Therefore, we get ((0 × 2) + 1) which is 1.
 
/**
 *  1 0 1 1 0 1 -> (0 *  2 ) + 1 = 1
 *  1 0 1 1 0 1 -> (1 *  2 ) + 0 = 2
 *  1 0 1 1 0 1 -> (2 *  2 ) + 1 = 5
 *  1 0 1 1 0 1 -> (5 *  2 ) + 1 = 11
 *  1 0 1 1 0 1 -> (11 * 2 ) + 0 = 22
 *  1 0 1 1 0 1 -> (22 * 2 ) + 1 = 45
 * -----------------------------------
 * R = 45
 * 
 * 1 0 1 -> (0 * 2) + 1 = 1
 * 1 0 1 -> (1 * 2) + 0 = 2
 * 1 0 1 -> (2 * 2) + 1 = 5
 * -----------------------------------
 * R = 5
 *  */
