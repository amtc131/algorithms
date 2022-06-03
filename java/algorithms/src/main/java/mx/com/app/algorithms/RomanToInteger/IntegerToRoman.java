package mx.com.app.algorithms.RomanToInteger;

public class IntegerToRoman {
    /**
     * 3 > III
     * 4 > IV
     * 58 > LVIII
     * 1994 > MCMXCIV
     * 
     */

    public  String intToRoman(int num) {
        String romanNumeral = "";
        int numbers[] = { 1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000 };
        String roman[] = { "I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M", };

        int i = numbers.length - 1;

        while (num != 0) {
            if (numbers[i] <= num) {
                romanNumeral += roman[i];
                num = num - numbers[i];
            } else {
                i -= 1;
            }
        }

        return romanNumeral;
    }

}
