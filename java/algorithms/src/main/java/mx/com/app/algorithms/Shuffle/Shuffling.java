package mx.com.app.algorithms.Shuffle;

import java.util.Arrays;

public class Shuffling {
    public static void main(String[] args) {
        int[] arr = {
                1, 0, -3, 8, 7, 3, 9, 4, 2, 5, 10, 6
        };
        int size = arr.length;
        shuffleElement(arr, size);
        System.out.println(Arrays.toString(arr));
    }

    public static void swapElement(int[] arr, int i, int j) {
        // Get i location element
        int temp = arr[i];
        // Set new values
        arr[i] = arr[j];
        arr[j] = temp;
    }

    // Returns the random location of array elements
    public static int randomLocation(int min, int max) {
        // Calculate random number between given range
        return min + (int) (Math.random() * ((max - min) + 1));
    }

    // Function which is shuffle given array elements
    public static void shuffleElement(int[] arr, int size) {
        // (i,j) indicate locations
        int j = 0;
        int i = 0;
        // Variable which is controlling the
        // execution process of loop
        int counter = 0;
        // Loop which is shuffling random elements in array
        while (counter < size) {
            // Get random location of array index
            i = randomLocation(0, size - 1);
            j = randomLocation(0, size - 1);
            if (i != j) {
                // Swap array elements
                swapElement(arr, i, j);
                counter++;
            }
        }

    }
}
