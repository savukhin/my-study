/*
Write a program that reads an integer angle (which may be positive or negative) and
normalizes it to a value between 0 and 359 degrees. Try it first with the % operator,
then with floorMod.
 */

package task_2;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int angle = scanner.nextInt();
        System.out.println(angle % 360);
        System.out.println(Math.floorMod(angle, 360));
        scanner.close();
    }
}