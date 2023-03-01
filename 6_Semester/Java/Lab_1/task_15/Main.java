/*
 * Write a program that stores Pascal’s triangle up to a given n in an
ArrayList<ArrayList<Integer>>
 */

package task_15;

import java.util.ArrayList;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.out.println("Введите целочисленное значение:");
        Scanner in = new Scanner(System.in);
        while (!in.hasNextInt()) {
            System.out.println("Введите корректное значение!");
            in.next();
        }
        int n = in.nextInt();
        // массив из 'n' строк
        ArrayList<ArrayList<Integer>> arr = new ArrayList<>();
        // обходим строки массива
        for (int i = 0; i < n; i++) {
            // строка из 'n-i' элементов
            ArrayList<Integer> arr1 = new ArrayList<>();
            // обходим элементы строки
            for (int j = 0; j < n - i; j++) {
                if (i == 0 || j == 0) {
                    // элементы первой строки
                    // и колонки равны единице
                    arr1.add(j, 1);
                    arr.add(i, arr1);
                } else {
                    // все остальные элементы — есть сумма двух
                    // предыдущих элементов в строке и в колонке
                    arr1.add(j, arr.get(i).get(j-1) + arr.get(i-1).get(j));
                    arr.add(i, arr1);
                }
            }
        }

        // обходим строки массива
        for(int i=0; i < n; i++){
            for(int j=0; j<arr.get(i).size(); j++) {
                System.out.printf(" %2d", arr.get(i).get(j));
            }
            System.out.println();

        }
    }
}
