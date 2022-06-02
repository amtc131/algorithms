package mx.com.app.algorithms.RomanToInteger;

import java.util.HashMap;

public class RomanToInteger {
    public static void main(String[] args) {
        System.out.println(  romanToInt("MCMXCIV") );
        System.out.println(  romanToInt("LVIII") );
        System.out.println(  romanToInt("III") );
    }

    public static int romanToInt(String s) {
        int result = 0;

        HashMap<Character,Integer> map = new HashMap<Character,Integer>();

        map.put('I', 1);
        map.put('V', 5);
        map.put('X', 10);
        map.put('L', 50);
        map.put('C', 100);
        map.put('D', 500);
        map.put('M', 1000);

        int l = s.length();
        for (int i = 0; i < l; i++) {
            // get the value index 0 and 0 + 1
            // validate if i + 1  < l 
            if(i + 1 < l && map.get(s.charAt(i)) < map.get(s.charAt(i + 1))){
                result -= map.get(s.charAt(i));
            }else{
                result +=  map.get(s.charAt(i));
            }
        }


        return result;
    }


}


/**
 * 
 * MCMXCI-
 * 1000 < 100   FALSE  1000    1000  1000
 * 100  < 1000  TRUE    100    -100   900 
 * 1000 < 10    FALSE  1000    1000  1900
 * 10   < 100   TRUE     10     -10  1890
 * 100  < 1     FALSE   100     100  1990
 * 1    <  5    TRUE      1      -1  1989 
 * i + 1 < l  < TRUE      5       5  1984
 *                        
 */                   

