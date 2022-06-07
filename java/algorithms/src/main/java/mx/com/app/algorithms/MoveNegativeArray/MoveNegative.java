package mx.com.app.algorithms.MoveNegativeArray;

import java.util.Arrays;

public class MoveNegative {
    // public static void main(String[] args) {
// //        arr1 []  = {1 , -1 , 3 , 2 , -7 , -5 , 11 , 6}
// //                   [1,  6,  3,  2,  11,  -5,  -7,  -1]
// //                   [1, -1,  3,  2,  -7,  -5,  11,  6]
        // reverseArray(new int[]{1 , -1 , 3 , 2 , -7 , -5 , 11 , 6});
// //                  {-1,  -5,  3,  2,  -7,  -5,  11,  -6}
// //                  {11,  2,  3,  -5,  -7,  -5,  -1,  -6}
// //                  [11,  2,  3,  -5,  -7,  -5,  -1,  -6]
//         MoveNegative.reverseArray(new int[]{-1,  -5,  3,  2,  -7,  -5,  11,  -6});
    // }

    public void reverseArray(int[] arr){
        
        // L= last index i= first index
        int l = arr.length - 1, i=0;

        while (i < l){
            int tmp = arr[i];
            if (arr[i] < 0 && arr[l] >= 0) {
                //when [i] index are have negative values 
                // and [l] is positive then swapping elements values
                // swap the given array elements
                arr[i] = arr[l];
                arr[l] = tmp;
                // update index
                i++;
                l--;
            }else if(arr[i] >= 0){
                // when element of [i] is not negative 
                i++;
            }else{
                l--;
            }

        }

        System.out.println(Arrays.toString(arr));
    }






}
