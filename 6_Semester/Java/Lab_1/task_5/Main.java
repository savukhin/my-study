/*
What happens when you cast a double to an int that is larger than the largest
possible int value? Try it out
 */

package task_5;

public class Main {

    public static void main(String[] args) {
        double largestInt = (double) Integer.MAX_VALUE;

        System.out.println(largestInt);
        System.out.println(largestInt + 1);
        System.out.println((int) largestInt);
        System.out.println((int) largestInt + 1);
        System.out.println((int) largestInt + 10);

    }
}