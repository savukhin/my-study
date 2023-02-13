/*
 * Write a program that stores Pascal’s triangle up to a given n in an
ArrayList<ArrayList<Integer>>
 */

package task_15;

import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    
        public static void main(String[] args) {
            Scanner scanner = new Scanner(System.in);
            int n = scanner.nextInt();
            scanner.close();
            
            ArrayList<ArrayList<Integer>> pascalTriangle = new ArrayList<ArrayList<Integer>>();
            
            for (int i = 0; i < n; i++) {
                ArrayList<Integer> row = new ArrayList<Integer>();
                for (int j = 0; j <= i; j++) {
                    if (j == 0 || j == i) {
                        row.add(1);
                    } else {
                        row.add(pascalTriangle.get(i - 1).get(j - 1) + pascalTriangle.get(i - 1).get(j));
                    }
                }
                pascalTriangle.add(row);
            }
            
            for (ArrayList<Integer> row : pascalTriangle) {
                for (Integer element : row) {
                    System.out.print(element + " ");
                }
                System.out.println();
            }
        }
}
