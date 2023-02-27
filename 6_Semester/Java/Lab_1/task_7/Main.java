/*
Write a program that reads in two numbers between 0 and 65535, stores them in
short variables, and computes their unsigned sum, difference, product, quotient,
and remainder, without converting them to int
 */

package task_7;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.out.println("Введите целочисленное значение от 0 до " + Short.MAX_VALUE+1 +":");
        Scanner in = new Scanner(System.in);
        int number1 = in.nextInt();
        while (number1<0 || number1>Short.MAX_VALUE+1) {
            System.out.println("Введите корректное значение!");
            number1 = in.nextInt();
        }


        int number2 = in.nextInt();
        while (number2<0 || number2>(Short.MAX_VALUE+1-number1)) {
            System.out.println("Введите корректное значение!");
            number2 = in.nextInt();
        }


        short sum = (short) (number1+number2);
        System.out.println("Сумма: " + sum);

        short multip = (short) (number1*number2);
        System.out.println("Произведение: " + multip);

        short del = (short) (number1/number2);
        System.out.println("Деление: " + del);

        short ost = (short) (number1 % number2);
        System.out.println("Остаток: " + ost);
    }
}
