package com.task1;

import static org.junit.Assert.assertTrue;

import org.junit.Test;

/**
 * Unit test for simple App.
 */
public class AppTest 
{
    /**
     * Rigorous Test :-)
     */
    @Test
    public void shouldAnswerWithTrue()
    {
        assertTrue( true );
    }

    @Test
    public void checkDatCalculator() {
        String date = "2023-03-28";
        java.sql.Date dat = java.sql.Date.valueOf(date);
        double MoscowWidth = 55.0 + 45.0 / 60.0;

        
    }
}
