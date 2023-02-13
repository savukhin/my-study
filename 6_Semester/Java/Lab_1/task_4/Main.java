/*
Write a program that prints the smallest and largest positive double value. Hint:
Look up Math.nextUp in the Java API
 */

package task_4;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println(Double.MIN_VALUE);
        System.out.println(Double.MAX_VALUE);
        System.out.println(Math.nextUp(Double.MAX_VALUE));
        scanner.close();
    }
}
