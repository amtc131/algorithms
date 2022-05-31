package mx.com.app.algorithms.PalindromeNumber;

public class PalindromeNumber {

    public static void main(String[] args) {
        System.out.println(isPalindrome(121));
        System.out.println(isPalindrome(-121));
        System.out.println(isPalindrome(10));
        System.out.println(isPalindrome(123));
    }

    public static boolean isPalindrome(int x) {

        String str = String.valueOf(x);
        StringBuilder b = new StringBuilder();
        for (int i = str.toCharArray().length - 1; i >= 0; i--) {
            b.append(str.charAt(i));
        }
        return b.toString().equals(str);
    }

}
