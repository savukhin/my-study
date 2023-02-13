package task_12;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

/*
 * The Java Development Kit includes a file src.zip with the source code of the
Java library. Unzip and, with your favorite text search tool, find usages of the
labeled break and continue sequences. Take one and rewrite it without a labeled
statement
 */

 // Not found in Java Development Kit So I've done it with my own code

public class Main {

    public static void main(String[] args) {
        List<List<Integer>> arr = createRandomArray();
        withOuter(arr);
        withoutOuter(arr);
    }

    private static List<List<Integer>> createRandomArray() {
        Random rand = new Random();
        int randCountI = rand.nextInt(10) + 5;
        List<List<Integer>> arr = new ArrayList<>();
        for (int i = 0; i < randCountI; i++) {
            arr.add(new ArrayList<>());
            int randCountJ = rand.nextInt(10) + 1;
            for (int j = 0; j < randCountJ; j++) {
                arr.get(i).add(rand.nextInt(randCountJ * 2));
                System.out.print(arr.get(i).get(j) + "\t");
            }
            System.out.println("<-" + arr.get(i).size() + " elements");
        }

        return arr;
    }

    public static void withOuter(List<List<Integer>> arr) {
        outer:
        for (int i = 0; i < arr.size(); i++) {
            if (arr.get(i).size() > 5) {
                System.out.printf("Arr[%d]:\n", i);
                for (int j = 0; j < arr.get(i).size(); j++) {
                    if (arr.get(i).get(j) > arr.get(i).size()) {
                        System.out.println(arr.get(i).get(j));
                    } else {
                        break outer;
                    }
                }
            }
        }
    }

    public static void withoutOuter(List<List<Integer>> arr) {
        for (int i = 0; i < arr.size(); i++) {
            if (arr.get(i).size() > 5) {
                System.out.printf("Arr[%d]:\n", i);
                for (int j = 0; j < arr.get(i).size(); j++) {
                    if (arr.get(i).get(j) > arr.get(i).size()) {
                        System.out.println(arr.get(i).get(j));
                    } else {
                        break;
                    }
                }
                break;
            }
        }

    }
}