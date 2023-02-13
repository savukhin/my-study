/*
 * Write a program that produces a random string of letters and digits by generating a
random long value and printing it in base 36
 */

package task_10;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long random = (long) (Math.random() * Long.MAX_VALUE);
        System.out.println(random);
        System.out.println(Long.toString(random, 36));
        scanner.close();
    }
}
