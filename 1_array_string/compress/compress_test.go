package main

import "testing"

func TestCompress(t *testing.T) {
	res1 := compress("ccaaatsss")
	if res1 != "2c3at3s" {
		t.Errorf(`result was %v, should be "2c3at3s"`, res1)
	}
	res2 := compress("ssssbbz")
	if res2 != "4s2bz" {
		t.Errorf(`result was %v, should be "4s2bz"`, res2)
	}
	res3 := compress("ppoppppp")
	if res3 != "2po5p" {
		t.Errorf(`result was %v, should be "2po5p"`, res3)
	}
	res4 := compress("nnneeeeeeeeeeeezz")
	if res4 != "3n12e2z" {
		t.Errorf(`result was %v, should be "3n12e2z"`, res4)
	}
}