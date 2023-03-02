package main

import "testing"

func TestPairProduct(t *testing.T) {
	n1 := []int{3,2,5,4,1}
	n2 := []int{4,7,9,2,5,1}
	n3 := []int{3,2,5,4,1}
	n4 := []int{4,6,8,2}

	res1 := pairProduct(n1, 8)
	if res1[0] != 1 || res1[1] != 3 {
		t.Errorf(`result %v should be []int{1,3}`, res1)
	}
	res2 := pairProduct(n1, 10)
	if res2[0] != 1 || res2[1] != 2 {
		t.Errorf(`result %v should be []int{1,2}`, res2)
	}
	res3 := pairProduct(n2, 5)
	if res3[0] != 4 || res3[1] != 5 {
		t.Errorf(`result %v should be []int{4,5}`, res3)
	}
	res4 := pairProduct(n2, 35)
	if res4[0] != 1 || res4[1] != 4  {
		t.Errorf(`result %v should be []int{1,4}`, res4)
	}
	res5 := pairProduct(n3, 10)
	if res5[0] != 1 || res5[1] != 2  {
		t.Errorf(`result %v should be []int{1,2}`, res5)
	}
	res6 := pairProduct(n4, 16)
	if res6[0] != 2 || res6[1] != 3  {
		t.Errorf(`result %v should be []int{2,3}`, res6)
	}
}