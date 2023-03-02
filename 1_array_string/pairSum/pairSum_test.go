package main

import "testing"

func TestPairSum(t *testing.T) {
	n1 := []int{3,2,5,4,1}
	n2 := []int{4,7,9,2,5,1}
	n3 := []int{1,6,7,2}
	n4 := []int{9,9}
	n5 := []int{6,4,2,8}

	res1 := pairSum(n1, 8)
	if res1[0] != 0 || res1[1] != 2 {
		t.Errorf(`result %v should be []int{0,2}`, res1)
	}
	res2 := pairSum(n2, 5)
	if res2[0] != 0 || res2[1] != 5 {
		t.Errorf(`result %v should be []int{0,5}`, res2)
	}
	res3 := pairSum(n2, 3)
	if res3[0] != 3 || res3[1] != 5 {
		t.Errorf(`result %v should be []int{3,5}`, res3)
	}
	res4 := pairSum(n3, 13)
	if res4[0] != 1 || res4[1] != 2  {
		t.Errorf(`result %v should be []int{1,2}`, res4)
	}
	res5 := pairSum(n4, 18)
	if res5[0] != 0 || res5[1] != 1  {
		t.Errorf(`result %v should be []int{0,1}`, res5)
	}
	res6 := pairSum(n5, 12)
	if res6[0] != 1 || res6[1] != 3  {
		t.Errorf(`result %v should be []int{1,3}`, res6)
	}
}