package com.task1;

import java.time.Duration;
import java.util.Calendar;
import java.sql.Date;

public class Calculator {
    /** 
     * Угол наклона оси вращения Земли к плоскости эклиптики
     */
    static private double dzeta_rad = Math.toRadians(23.5);

    /**
     * 
     * @param date - Дата, в которую рассчитывается длительнность дня
     * @param width - Широта местности
     * @return - Длительность дня
     */
    public static Duration getDuration(Date date, double width) {
        double width_rad = Math.toRadians(width);

        Calendar cal = Calendar.getInstance();
        cal.setTime(date);
        int yearDay = cal.get(Calendar.DAY_OF_YEAR);

        long hours = (long)(12. - (24. / Math.PI) * Math.asin(Math.tan(width_rad) * Math.tan(dzeta_rad) * Math.cos((yearDay + 10) * 2 * Math.PI / 365.25))) ;
        return Duration.ofHours(hours);
    }
}
