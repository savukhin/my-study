package com.task1;

import java.util.Locale;
import java.util.Scanner;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.time.Duration;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        String date = "2022-12-28";
        // String date = "2023-03-28";
        // String date = "2023-06-22";
        // String date = "2023-03-22";
        // String date = "2023-05-22";
        // String date = "2023-03-28";
        java.sql.Date dat = java.sql.Date.valueOf(date);

        double MoscowWidth = 55.0 + 45.0 / 60.0;
        double SabettaWidth = 71.0 + 16 / 60.0;

        Duration duration = Calculator.getDuration(dat, SabettaWidth);

        System.out.println(duration.toHours());

        DateFormat df = new SimpleDateFormat("dd MMM YYYY G", Locale.ENGLISH);
        System.out.println(df.format(dat.getTime()));
    }
}
