package mx.com.app.algorithms.binary.search;

public class RotatedArray {

    /***
     * Example
     * 
     * low                              mid                               high
     *  |                                |                                  |  
     * -----------------------------------------------------------------------
     *  0    1   2   3   4   5  6  7  8  9 10 11 12 13 14  15  16  17  18  19 
     *  176 188 199 200 210 222 1 10 20 47 59 63 75 88 99 107 120 133 155 162
     * -----------------------------------------------------------------------
     * Target: 200
     * 
     * 1.- Let's initialize our low, mid, and high pointers. 
     * 2.- We compare the values of the low, mid, and high pointers. Since array[mid] < array[high], we know that the section from mid to high is sorted.
     * 3.- We also check if the target value lies in this section's range (47 - 162) 
     *      and since it doesn't, we will continue our search in the section between the low and mid pointers.
     * 4.- We update the high pointer to mid -1. We can also discard the section to the right of the mid pointer.Once high pointer is updated, we can also update our mid pointer.
     *
     * low             mid           high      
     *  |               |            |               
     * ---------------------------------
     * 0    1   2   3   4   5  6  7  8 
     * 176 188 199 200 210 222 1 10 20 
     * ---------------------------------
     * 
     * 6.- We compare the values of the low, mid, and high pointers. Since array[mid] > array[low], we know that the section from low to mid is sorted.
     * 7.- We also check if the target value lies in this section's range (176 - 210) and since it does, we will continue our search in the section between the low and mid pointers.
     * 
     * low   mid     high     
     *  |    |       |   
     * ----------------
     * 0    1   2   3  
     * 176 188 199 200 
     * ----------------
     * 
     * 8.- We update the high pointer to mid -1. We can also discard the section to the right of the mid pointer.Once high pointer is updated, we can also update our mid pointer.
     * 9.- We compare the values of the low, mid, and high pointers. Since array[mid] > array[low], we know that the section from low to mid is sorted
     * 10.- We check if the target value lies in this section's range (176 - 188) and since it doesn't, we will continue our search in the section between the mid and high pointers.
     * 11.- We update the low pointer to mid + 1. We can also discard the section to the left of the mid pointer.Once low pointer is updated, we can also update our mid pointer.
     * 12.- We compare the values of the low, mid, and high pointers. Since array[mid] < array[high], we know that the section from mid to high is sorted.
     * 
     * low/mid  high
     * |       / 
     * --------
     *  2   3  
     * 199 200 
     * --------
     * 
     * 13.- We check if the target value lies in this section's range (199 - 200) and since it does, we will continue our search in the section between the mid and high pointers.
     * 14.- We update the low pointer to mid + 1. We can also discard the section to the left of the mid pointer.Once low pointer is updated, we can also update our mid pointer. Our mid pointer now points to the target value.
     *
     *  low/mid high
     *   \    /
     * -------
     *    3
     *   200
     * -------
     * 
     *  @param int[] nums, target int
     */


    public static void main(String[] args) {
        System.out.println( binarySearchRotated(new int[]{6, 7, 1, 2, 3, 4, 5}, 3 ) );        
        System.out.println( binarySearchRotated(new int[]{6, 7, 1, 2, 3, 4, 5}, 6 ) );        
        System.out.println( binarySearchRotated(new int[]{4, 5, 6, 1, 2, 3}, 3 ) );        
        System.out.println( binarySearchRotated(new int[]{4, 5, 6, 1, 2, 3}, 6 ) );        
        System.out.println( binarySearchRotated(new int[]{6, 7, 1, 2, 3, 4, 5}, 3 ) );        
    }

    static int binarySearchRotated(int[] nums, int target) {
		// TODO: Write - Your - Code

        int  start = 0;
        int end = nums.length - 1;

        if(start > end){
            return -1;
        }
        while(start <= end){
            // Fiding the mid using floor division
            int mid = start + ((end - start) / 2);
        
            // Target value is present at the middle of the array
            if(nums[mid] == target){
                return mid;
            }
            //start to mid is sorted
            if(nums[start] <= nums[mid]){
                if(nums[start] <= target && target <= nums[mid]){
                    end = mid - 1;
                }else{
                    start = mid + 1;
                }
            }

            //mid to end is sorted
            else{
                if(nums[mid] < target && target <= nums[end]){
                    start = mid + 1;
                }else{
                    end = mid - 1;
                }
            }
        }
		return -1;
	}


}
