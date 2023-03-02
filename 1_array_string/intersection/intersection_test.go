package main

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	res1 := intersection([]interface{}{4,2,1,6}, []interface{}{3,6,9,2,10})
	if res1[0] != 6 || res1[1] != 2 {
		t.Errorf(`resulted in %v, expected []interface{}{6,2}`, res1)
	}
	res2 := intersection([]interface{}{2,4,6}, []interface{}{4,2})
	if res2[0] != 4 || res2[1] != 2 {
		t.Errorf(`resulted in %v, expected []interface{}{4,2}`, res2)
	}
	res3 := intersection([]interface{}{4,2,1}, []interface{}{1,2,4,6})
	if res3[0] != 1 || res3[1] != 2 || res3[2] != 4 {
		t.Errorf(`resulted in %v, expected []interface{}{1,2,4}`, res3)
	}
	res4 := intersection([]interface{}{0,1,2}, []interface{}{10,11}) 
	if len(res4) > 0 {
		t.Errorf(`resulted in %v, expected []interface{}`, res4)
	}
}