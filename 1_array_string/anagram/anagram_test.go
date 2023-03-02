package main

import "testing"

func TestAnagrams(t *testing.T) {
res1 := anagrams("restful", "fluster")
if res1 != true {
	t.Errorf("result was %v should be true", res1)
}
res2 := anagrams("cats", "tocs")
if res2 != false {
	t.Errorf("result was %v should be false", res2)
}
res3 := anagrams("monkeyswrite", "newyorktimes")
if res3 != true {
	t.Errorf("result was %v should be true", res3)
}
res4 := anagrams("paper", "reapa")
if res4 != false {
	t.Errorf("result was %v should be false", res4)
}
res5 := anagrams("elbow", "below") 
if res5 != true {
	t.Errorf("result was %v should be false", res5)
}
res6 := anagrams("tax", "taxi")
if res6 != false {
	t.Errorf("result was %v should be false", res6)
}
}