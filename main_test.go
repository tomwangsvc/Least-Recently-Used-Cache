package main

import "testing"

func Test_cache(t *testing.T) {
	cache := NewCache(2)
	cache.put(1, 1)
	cache.put(2, 2)
	result1 := cache.get(1)
	if result1 != 1 {
		t.Errorf("test result1 %d is not 1", result1)
	}
	cache.put(3, 3)
	result2 := cache.get(2)
	if result2 != -1 {
		t.Errorf("test result2 %d is not -1", result2)
	}
	cache.put(4, 4)
	result3 := cache.get(1)
	if result3 != -1 {
		t.Errorf("test result3 %d is not -1", result3)
	}
	result4 := cache.get(3)
	if result4 != 3 {
		t.Errorf("test result4 %d is not 3", result3)
	}
	result5 := cache.get(4)
	if result5 != 4 {
		t.Errorf("test result5 %d is not 4", result3)
	}
	result6 :=cache.delete(3)
	if result6 != 3 {
		t.Errorf("test result6 %d is not 3", result6)
	}
	result7 := cache.get(3)
	if result7 != -1 {
		t.Errorf("test result7 %d is not -1", result7)
	}
}
