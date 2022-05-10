package test

import (
	"testing"

	"github.com/amtc131/algorithms/structure/hashmap"
)

func TestHashMap(t *testing.T) {

	mp := hashmap.New()

	t.Run("Test 1: Put(10) and cheking if Get() is correct", func(t *testing.T) {
		mp.Put("test", 10)

		got := mp.Get("test")
		if got != 10 {
			t.Fatalf("Put() = %v, Got: %v", 10, got)
		}
	})

}
