/*
 * Write a program that reads an integer and prints it in binary, octal, and hexadecimal.
 * Print the reciprocal as a hexadecimal floating-point number
*/

package task_1;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int number = scanner.nextInt();
        System.out.println(Integer.toBinaryString(number));
        System.out.println(Integer.toOctalString(number));
        System.out.println(Integer.toHexString(number));
        System.out.println(Float.toHexString(1.0f / number));

        scanner.close();
    }
}