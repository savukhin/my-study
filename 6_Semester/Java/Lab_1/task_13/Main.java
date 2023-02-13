/*
 * Write a program that prints a lottery combination, picking six distinct numbers
between 1 and 49. To pick six distinct numbers, start with an array list filled with 1
… 49. Pick a random index and remove the element. Repeat six times. Print the
result in sorted order.
*/

package task_13;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Random;

public class Main {

    public static void main(String[] args) {
        ArrayList<Integer> numbers = new ArrayList<>();
        for (int i = 1; i <= 49; i++) {
            numbers.add(i);
        }
        Random random = new Random();
        for (int i = 0; i < 6; i++) {
            int index = random.nextInt(numbers.size());
            System.out.println(numbers.get(index));
            numbers.remove(index);
        }
        Collections.sort(numbers);
        System.out.println(numbers);
    }
}
