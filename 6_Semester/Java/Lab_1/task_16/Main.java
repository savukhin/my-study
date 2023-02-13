/*
 * Improve the average method so that it is called with at least one parameter
 */

package task_16;

import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Main {
    
        public static void main(String[] args) {
            //ArrayList
            System.out.print("Enter string numbers:");
            Scanner in = new Scanner(System.in);
            String str = in.nextLine();
            if (str.length() > 0) {
                List<Double> list = Arrays.stream(str.trim().split("\\s+")).map(Double::parseDouble).collect(Collectors.toList());
                System.out.println(average(list));
            }
        }
        
        public static double average(int... numbers) {
            double sum = 0;
            for (int number : numbers) {
                sum += number;
            }
            
            return sum / numbers.length;
        }

        public static double average(List<Double> arr) {
            double sum = 0;
            for (double it : arr) {
                sum += it;
            }
            return (double) sum / arr.size();
        }
    
}
