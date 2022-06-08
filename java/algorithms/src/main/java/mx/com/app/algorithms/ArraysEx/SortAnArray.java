package mx.com.app.algorithms.ArraysEx;

import java.util.Arrays;

public class SortAnArray {

    public static void main(String[] args) {
        int arr[] = { 1, 2, 0, 0, 1 };
        sort(arr);
        System.out.println(" - " + Arrays.toString(arr));
    }

    public static void sort(int arr[]) {

        int l = arr.length;

        for (int j = 0; j < l; j++) {
            for (int i = 0; i < l; i++) {
                // index 0 of j 
                int tmp = arr[j];
                // if index arr[j] is < arr[i]
                if (tmp < arr[i]) {
                    //update in index arr[j] in arr[i]  
                    arr[j] = arr[i];
                    // update arr[i] in arr[j]
                    arr[i] = tmp;
                }
            }
        }
    }
}

/*                        
 * [1,2,0,0,1]           
 * [2, 1, 0, 0, 1] change 2 and 1
 * [1, 2, 0, 0, 1] change 2 and 1
 * [0, 1, 2, 0, 1] change 0 and 1, change 1 and 2
 * [0, 0, 1, 1, 2] change 0 and 1, change 1 and 2, change 1 and 2
 * 
 */