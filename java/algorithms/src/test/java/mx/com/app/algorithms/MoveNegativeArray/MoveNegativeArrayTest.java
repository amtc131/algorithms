package mx.com.app.algorithms.MoveNegativeArray;

import static org.junit.Assert.assertArrayEquals;
import static org.mockito.Mockito.doNothing;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

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
        int arr[]={1 , -1 , 3 , 2 , -7 , -5 , 11 , 6};
        doNothing().when(moveNegative).reverseArray(arr);
        moveNegative.reverseArray(arr);

        assertArrayEquals(new int[]{1,  6,  3,  2,  11,  -5,  -7,  -1}, arr);

    }
}
