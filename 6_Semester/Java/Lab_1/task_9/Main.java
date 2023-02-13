/*
 * Section 1.5.3, “String Comparison,” on p. 21 has an example of two strings s and t
so that s.equals(t) but s != t. Come up with a different example that doesn’t
use substring).
 */

package task_9;

import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        System.out.print("Enter string1:");
        String str1 = in.nextLine();
        System.out.print("Enter string2:");
        String str2 = in.nextLine();

        System.out.println("string1 == string2\nResult ignore case:" + str1.trim().equalsIgnoreCase(str2.trim()));
    }
        
}
