package main

import "testing"

func TestMostFrequent(t *testing.T) {
	res1 := mostFrequentChar("bookeeper") 
	if res1 != "e" {
		t.Errorf(`result is %v, should be "e"`, res1)
	}
	res2 := mostFrequentChar("david") 
	if res2 != "d" {
		t.Errorf(`result is %v, should be "d"`, res2)
	}
	res3 := mostFrequentChar("abby") 
	if res3 != "b" {
		t.Errorf(`result is %v, should be "b"`, res3)
	}
	res4 := mostFrequentChar("mississippi") 
	if res4 != "i" {
		t.Errorf(`result is %v, should be "i"`, res4)
	}
	res5 := mostFrequentChar("potato") 
	if res5 != "o" {
		t.Errorf(`result is %v, should be "o"`, res5)
	}
	res6 := mostFrequentChar("eleventennine") 
	if res6 != "e" {
		t.Errorf(`result is %v, should be "e"`, res6)
	}
	res7 := mostFrequentChar("riverbed") 
	if res7 != "r" {
		t.Errorf(`result is %v, should be "r"`, res7)
	}
}