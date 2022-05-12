package mx.com.app.algorithms.binary.search;

public class SorterdArray{
    public static void main(String[] args) {
       System.out.println( binarySearch(new int[]{1,3,9,10,12}, 9));
       System.out.println( binarySearch(new int[]{}, 12));
       System.out.println( binarySearch(new int[]{0, 1}, 1));
       System.out.println( binarySearch(new int[]{1, 2, 3}, 3));
       System.out.println( binarySearch(new int[]{-1, 0, 3, 5, 9, 12}, 9));
       System.out.println( binarySearch(new int[]{-1, 0, 3, 5, 9, 12}, 2));
    }

    private static int binarySearch(int[] nums, int target) {
        int low = 0;
        int high = nums.length - 1;

        while(low <= high){
           // Finding the mid using floor division
			int mid = low + ((high - low) / 2);

			// Target value is present at the middle of the array
			if (nums[mid] == target) {
				return mid;
			}

			// Target value is present in the low subarray
			else if (target < nums[mid]) {
				high = mid - 1;
			}

			// Target value is present in the high subarray
			else if (target > nums[mid]) {
				low = mid + 1;
			}
        }
        
        // Target value is not present in the array
        return -1;
    }

}




