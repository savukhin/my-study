/*
 * Write a program that computes the factorial n! = 1 × 2 × … × n, using
BigInteger. Compute the factorial of 1000.
 */

package task_6;

import java.math.BigInteger;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int number = scanner.nextInt();
        BigInteger factorial = BigInteger.ONE;
        for (int i = 1; i <= number; i++) {
            factorial = factorial.multiply(BigInteger.valueOf(i));
        }
        System.out.println(factorial);
        scanner.close();
    }
}