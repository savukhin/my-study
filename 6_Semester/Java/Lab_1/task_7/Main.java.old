/*
Write a program that reads in two numbers between 0 and 65535, stores them in
short variables, and computes their unsigned sum, difference, product, quotient,
and remainder, without converting them to int
 */

package task_7;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        // Scanner scanner = new Scanner(System.in);
        // short number1 = scanner.nextShort();
        // short number2 = scanner.nextShort();
        // System.out.println((short) (number1 + number2));
        // System.out.println((short) (number1 - number2));
        // System.out.println((short) (number1 * number2));
        // System.out.println((short) (number1 / number2));
        // System.out.println((short) (number1 % number2));
        // scanner.close();
        int offset=32768;
        Scanner ar=new Scanner(System.in);
        short a=(short)( ar.nextInt()-offset);
        short b=(short)( ar.nextInt()-offset);
        System.out.println("*"+((a+offset)*(b+offset)));
        System.out.println("+"+((a+offset)+(b+offset)));
        System.out.println("+"+((a+offset)-(b+offset)));
        System.out.println("+"+((a+offset)/(b+offset)));
        System.out.println("+"+((a+offset)%(b+offset)));
        ar.close();
    }
}
