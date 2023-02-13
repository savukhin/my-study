/*
 * Write a program that reads a string and prints all of its nonempty substrings
 */

package task_8;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String string = scanner.nextLine();
        for (int i = 0; i < string.length(); i++) {
            for (int j = i + 1; j <= string.length(); j++) {
                System.out.println(string.substring(i, j));
            }
        }
        scanner.close();
    }
}
