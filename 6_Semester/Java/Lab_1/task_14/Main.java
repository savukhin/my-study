/*
Write a program that reads a two-dimensional array of integers and determines
whether it is a magic square (that is, whether the sum of all rows, all columns, and
the diagonals is the same). Accept lines of input that you break up into individual
integers, and stop when the user enters a blank line. For example, with the input
16 3 2 13
3 10 11 8
9 6 7 12
4 15 14 1
(Blank line)
your program should respond affirmatively.

1 2 3 47 23 5 2
1 2 4 5 7
1 2 3 4 5 6 7
9 8 7 6 5 4 3 2 1 
Max sqaure is 
____________
|1 2 3 47 23| 5 2
|1 2 4 5 7  |
|1 2 3 4 5  |6 7
|9 8 7 6 5  |4 3 2 1
------------
*/

package task_14;

import java.util.ArrayList;
import java.util.Scanner;


public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        ArrayList<ArrayList<Integer>> array = readArray(scanner);
        scanner.close();        
            
        if (isMagicMatrix(array)) {
            System.out.println("It is a magic square");
        } else {
            System.out.println("It is not a magic square");
        }
    }

    public static boolean isSquare(ArrayList<ArrayList<Integer>> array, int x, int y, int size) {
        for (int i = x; i < x + size; i++) {
            if (array.get(i).size() < y + size) {
                return false;
            }
        }

        if (array.size() < x + size) {
            return false;
        }

        return true;
    }

    public static boolean isMagicSquare(ArrayList<ArrayList<Integer>> square, int x, int y, int size) {
        int mainDiagonalSum = 0;
        int sideDiagonalSum = 0;
        ArrayList<Integer> rowSums = new ArrayList<Integer>();
        ArrayList<Integer> columnSums = new ArrayList<Integer>();

        for (int i = 0; i < size; i++) {
            int rowSum = 0;
            int columnSum = 0;
            for (int j = 0; j < size; j++) {
                System.out.print("[" + x + i + "][" + y + j + "]");
                System.out.println( columnSums.size() + " " + rowSums.size() + " " + square.get(x + i).get(y + j) + " " + square.get(x + j).get(y + i) + " " + mainDiagonalSum + " " + sideDiagonalSum + "");
                rowSum += square.get(x + i).get(y + j);
                columnSum += square.get(x + j).get(y + i);
            }
            
            mainDiagonalSum += square.get(x + i).get(x + i);
            sideDiagonalSum += square.get(x + i).get(x + size - i - 1);

            rowSums.add(rowSum);
            columnSums.add(columnSum);
        }

        return mainDiagonalSum == sideDiagonalSum && mainDiagonalSum == rowSums.get(0) && mainDiagonalSum == columnSums.get(0);
    }

    public static boolean isMagicMatrix(ArrayList<ArrayList<Integer>> array) {
        for (int x = 0; x < array.size(); x++) {
            for (int y = 0; y < array.get(x).size(); y++) {
                //выбираем все стороны

                for (int size = 2; size <= Math.min(array.size() - x, array.get(x).size() - y); size++) {
                    if (!isSquare(array, x, y, size)) {
                        break;
                    }
                    //просматриваем квадрат по определенным координатам с определенной стороной
                    if (isMagicSquare(array, x, y, size)) {
                        return true;
                    }
                }
            }
        }

        return false;
    }

    
    public static ArrayList<ArrayList<Integer>> readArray(Scanner scanner) {
        ArrayList<ArrayList<Integer>> array = new ArrayList<ArrayList<Integer>>();
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (line.equals("")) {
                break;
            }
            String[] numbers = line.split(" ");
            ArrayList<Integer> row = new ArrayList<Integer>();
            for (String number : numbers) {
                row.add(Integer.parseInt(number));
            }
            array.add(row);
        }
        return array;
    }
}
