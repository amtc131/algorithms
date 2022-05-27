package mx.com.app.algorithms.MaxElement;

public class MaxElement {
    public static void main(String[] args) {
        int[] number = { 1, 3, 5, 2, 3, 1, 4, 4, 2 };
        int max = number[0];
        for (int i = 1; i < number.length; i++) {
            if (number[i] > max) {
                max = number[i];
            }
        }

        System.out.println("Max element: " + max);
    }
}
