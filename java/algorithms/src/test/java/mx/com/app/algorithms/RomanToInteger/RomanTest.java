package mx.com.app.algorithms.RomanToInteger;

import static org.junit.Assert.assertEquals;
import static org.mockito.Mockito.when;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.mockito.runners.MockitoJUnitRunner;

@RunWith(MockitoJUnitRunner.class)
public class RomanTest {
    
    @Mock
    private IntegerToRoman integerToRoman;

    @Before
    public void init(){
        MockitoAnnotations.initMocks(this);
    }

    @Test
    public void intToRomanWhenValueis3ThenReturnStringRomanNumeral(){
        when(integerToRoman.intToRoman(3)).thenReturn("III");
        
        assertEquals("III", integerToRoman.intToRoman(3));
    }
    @Test
    public void intToRomanWhenInto4ThenReturnStringRomanNumeral(){
        when(integerToRoman.intToRoman(4)).thenReturn("IV");
        
        assertEquals("IV", integerToRoman.intToRoman(4));
    }

    @Test
    public void intToRomanWhenInto58ThenReturnStringRomanNumeral(){
        when(integerToRoman.intToRoman(58)).thenReturn("LVIII");
        
        assertEquals("LVIII", integerToRoman.intToRoman(58));
    }

    @Test
    public void intToRomanWhenInto1994ThenReturnStringRomanNumeral(){
        when(integerToRoman.intToRoman(1994)).thenReturn("MCMXCIV");
        
        assertEquals("MCMXCIV", integerToRoman.intToRoman(1994));
    }
}
