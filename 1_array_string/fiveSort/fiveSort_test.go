package main

import "testing"

func TestFiveSort(t *testing.T) {
	res1 := fiveSort([]int{12, 5, 1, 5, 12, 7})
	if !Equal(res1, []int{12, 7, 1, 12, 5, 5}) {
		t.Errorf(`resulted in %v, expected []int{12,7,1,12,5,5}`, res1)
	}

 	res2 := fiveSort([]int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5})
	if !Equal(res2, []int{2, 2, 10, 6, 1, 5, 5, 5, 5, 5}) {
		t.Errorf(`resulted in %v, expected []int{{2, 2, 10, 6, 1, 5, 5, 5, 5, 5}}`, res2)
	}
	res3 := fiveSort([]int{5, 5, 5, 1, 1, 1, 4})
	if !Equal(res3, []int{4, 1, 1, 1, 5, 5, 5}) {
		t.Errorf(`resulted in %v, expected []int{{4, 1, 1, 1, 5, 5, 5}}`, res3)
	}
	res4 := fiveSort([]int{5, 5, 6, 5, 5, 5, 5}) 
	if !Equal(res4, []int{6, 5, 5, 5, 5, 5, 5}) {
		t.Errorf(`resulted in %v, expected []int{6, 5, 5, 5, 5, 5, 5}`, res4)
	}

	res5 := fiveSort([]int{5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5})
	if !Equal(res5, []int{4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5}) {
		t.Errorf(`resulted in %v, expected []int{4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5}`, res5)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}