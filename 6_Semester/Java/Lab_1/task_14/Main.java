package task_14;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int[][] array = new int[4][4];

        for(int i=0; i< array.length; i++){
            for(int j=0; j<array[i].length; j++){
                System.out.println("Введите целочисленное значение array[" + i + "][" + j + "]:");
                Scanner in = new Scanner(System.in);
                while (!in.hasNextInt()) {
                    System.out.println("Введите корректное значение!");
                    in.next();
                }
                array[i][j]= in.nextInt();
            }
        }
        //check column
        int sum=0;
        int[] array_int = new int[10];
        //string
        for(int i=0; i< array.length; i++){
            for (int j=0; j<array[i].length; j++){
                sum+=array[i][j];
            }
            array_int[i] = sum;
            sum=0;
        }
        //column
        for(int i=0; i< array.length; i++){
            for (int j=0; j<array.length; j++){
                sum+=array[i][j];
            }
            array_int[i+array.length] = sum;
            sum=0;
        }
        //diagon
        for(int i=0; i<array.length; i++){
            sum+=array[i][i];
        }
        array_int[8] = sum;
        sum=0;
        for(int i=(array.length-1); i>=0; i--){
            sum+=array[i][i];
        }
        array_int[9] = sum;

        //check
        int flag = 0;
        for(int j=0; j<9; j++){
            if(array_int[j]!= array_int[j+1]){
                flag =1;
            }
        }
        if(flag==1){
            System.out.println(false);
        }else {
            System.out.println(true);
        }
    }
}