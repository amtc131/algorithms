package mx.com.app.algorithms.RomanToInteger;

import java.util.HashMap;

public class IntegerToRoman {
    public static void main(String[] args) {
        
    }

    public static String intToRoman(int num) {
        String result = "";
        HashMap<Character,Integer> map = new HashMap<Character,Integer>();

        map.put('I', 1);
        map.put('V', 5);
        map.put('X', 10);
        map.put('L', 50);
        map.put('C', 100);
        map.put('D', 500);
        map.put('M', 1000);

        
        return result;
    }

}

/**
 *  3    > III
 *  4    > IV 
 *  58   > LVIII
 *  1994 > MCMXCIV
 * 
 */

