package main

import "testing"

func TestUncompress(t *testing.T) {
	str1 := "2c3a1t"
	str2 := "4s2b"
	str3 := "2p1o5p"

	res1 := uncompress(str1)
	if res1 != "ccaaat" {
		t.Errorf(`response is %s, should be "ccaaat"`, res1)
	}
	res2 := uncompress(str2)
	if res2 != "ssssbb" {
		t.Errorf(`response is %s, should be "ssssbb"`, res2)
	}
	res3 := uncompress(str3)
	if res3 != "ppoppppp" {
		t.Errorf(`response is %s, should be ppoppppp`, res3)
	}
}

