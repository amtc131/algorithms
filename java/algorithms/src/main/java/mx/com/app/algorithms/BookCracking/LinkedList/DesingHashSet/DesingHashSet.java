package mx.com.app.algorithms.BookCracking.LinkedList.DesingHashSet;

import java.util.HashSet;
import java.util.Set;

public class DesingHashSet {
        
    public static void main(String[] args) {
        MyHashSet myHashSet= new MyHashSet();
        myHashSet.add(1);      // set = [1]
        myHashSet.add(2);      // set = [1, 2]
        myHashSet.contains(1); // return True
        myHashSet.contains(3); // return False, (not found)
        myHashSet.add(2);      // set = [1, 2]
        myHashSet.contains(2); // return True
        myHashSet.remove(2);   // set = [1]
        myHashSet.contains(2); // return False, (already removed)
        
    }


}
class MyHashSet {

    Set<Integer> hash;

    public MyHashSet() {
        hash = new HashSet<Integer>();
        hash.add(null);
    }
    
    public void add(int key) {
        if(!contains(key)){
            hash.add(key);
        }
        hash.add(null);
    }
    
    public void remove(int key) {
        if(contains(key)){
            hash.remove(key);
        }
    }
    
    public boolean contains(int key) {
        return hash.contains(key);
    }
}
