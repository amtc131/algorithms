package mx.com.app.algorithms.MoveNegativeArray;

import static org.junit.Assert.assertArrayEquals;
import static org.mockito.Mockito.doCallRealMethod;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.mockito.runners.MockitoJUnitRunner;

@RunWith(MockitoJUnitRunner.class)
public class MoveNegativeArrayTest {
    
    @Mock
    private MoveNegative moveNegative;

    @Before
    public void init() {
        MockitoAnnotations.initMocks(this);
    }

    @Test
    public void MoveAllNegativeElementsToEnd(){
        int got[]={1 , -1 , 3 , 2 , -7 , -5 , 11 , 6};
        int want[] ={1,  6,  3,  2,  11,  -5,  -7,  -1};
        doCallRealMethod().when(moveNegative).reverseArray(got);
        moveNegative.reverseArray(got);
        assertArrayEquals(want, got);
    }
    
    @Test
    public void MoveAllNegativeElementsToEndTwo(){
        int got[]= {-1,  -5,  3,  2,  -7,  -5,  11,  -6};
        int want[] ={11,  2,  3,  -5,  -7,  -5,  -1,  -6};
        doCallRealMethod().when(moveNegative).reverseArray(got);
        moveNegative.reverseArray(got);
        assertArrayEquals(want, got);
    }
}
