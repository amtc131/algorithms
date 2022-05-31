package mx.com.app.algorithms.TwoSum;

import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class TwoSum {

    public static void main(String[] args) {
       System.out.println( Arrays.toString(twoSum(new int[] { 2, 7, 11, 15 }, 9)));
       System.out.println( Arrays.toString(twoSum(new int[] { 3,2,4}, 6)));
       System.out.println( Arrays.toString(twoSum(new int[] { 3,3}, 6)));
       System.out.println( Arrays.toString(twoSum(new int[] { 2,5,5,11}, 10)));
    }

    public static int[] twoSum(int[] nums, int target) {
        
        Map<Integer, Integer> s = new HashMap<Integer, Integer>();
        
        for (int i = 0; i < nums.length; i++) {
            int sum = target - nums[i];
            if(s.containsKey(sum)){
                return new int[]{s.get(sum), i};
            }
            s.put(nums[i], i);
        }

        return null;
    }
}
