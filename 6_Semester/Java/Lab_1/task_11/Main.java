/*
 * Write a program that reads a line of text and prints all characters that are not ASCII,
together with their Unicode values
 */

package task_11;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String string = scanner.nextLine();
        for (int i = 0; i < string.length(); i++) {
            if (string.charAt(i) > 127) {
                System.out.println(string.charAt(i) + " " + (int) string.charAt(i));
            }
        }
        scanner.close();
    }
}
