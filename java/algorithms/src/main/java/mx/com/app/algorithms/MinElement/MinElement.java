package mx.com.app.algorithms.MinElement;

public class MinElement {

    public static void main(String[] args) {
        int[] arr = { 1, 2, 5, 0, 8, 3, 4, 6, 7, 9 };

        int min = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] < min) {
                min = arr[i];
            }
        }
        System.out.println("Value min: " + min);
    }

}
