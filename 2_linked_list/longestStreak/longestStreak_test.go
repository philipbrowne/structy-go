package main

import (
	"testing"
)

func TestLongestStreakA(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 5}
	c := &Node{Val: 7}
	d := &Node{Val: 7}
	e := &Node{Val: 7}
	f := &Node{Val: 6}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreak(a)
	if res != 3 {
		t.Errorf("expected result to be 3, was %d", res)
	}
}

func TestLongestStreakRecursive_A(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 5}
	c := &Node{Val: 7}
	d := &Node{Val: 7}
	e := &Node{Val: 7}
	f := &Node{Val: 6}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreakRecursive(a, 0, 0)
	if res != 3 {
		t.Errorf("expected result to be 3, was %d", res)
	}
}


func TestLongestStreakB(t *testing.T) {
	a := &Node{Val: 3}
	b := &Node{Val: 3}
	c := &Node{Val: 3}
	d := &Node{Val: 3}
	e := &Node{Val: 9}
	f := &Node{Val: 9}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreak(a)
	if res != 4 {
		t.Errorf("expected result to be 4, was %d", res)
	}
}

func TestLongestStreakRecursive_B(t *testing.T) {
	a := &Node{Val: 3}
	b := &Node{Val: 3}
	c := &Node{Val: 3}
	d := &Node{Val: 3}
	e := &Node{Val: 9}
	f := &Node{Val: 9}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreakRecursive(a, 0, 0)
	if res != 4 {
		t.Errorf("expected result to be 4, was %d", res)
	}
}

func TestLongestStreakC(t *testing.T) {
	a := &Node{Val: 9}
	b := &Node{Val: 9}
	c := &Node{Val: 1}
	d := &Node{Val: 9}
	e := &Node{Val: 9}
	f := &Node{Val: 9}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreak(a)
	if res != 3 {
		t.Errorf("expected result to be 3, was %d", res)
	}
}

func TestLongestStreakRecursive_C(t *testing.T) {
	a := &Node{Val: 9}
	b := &Node{Val: 9}
	c := &Node{Val: 1}
	d := &Node{Val: 9}
	e := &Node{Val: 9}
	f := &Node{Val: 9}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	res := longestStreakRecursive(a, 0, 0)
	if res != 3 {
		t.Errorf("expected result to be 3, was %d", res)
	}
}

func TestLongestStreakD(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 5}

	a.Next = b
	res := longestStreak(a)
	if res != 2 {
		t.Errorf("expected result to be 2, was %d", res)
	}
}

func TestLongestStreakRecursive_D(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 5}

	a.Next = b
	res := longestStreakRecursive(a, 0, 0)
	if res != 2 {
		t.Errorf("expected result to be 2, was %d", res)
	}
}

func TestLongestStreakE(t *testing.T) {
	a := &Node{Val: 4}

	res := longestStreak(a)
	if res != 1 {
		t.Errorf("expected result to be 1, was %d", res)
	}
}

func TestLongestStreakRecursive_E(t *testing.T) {
	a := &Node{Val: 4}

	res := longestStreakRecursive(a, 0, 0)
	if res != 1 {
		t.Errorf("expected result to be 1, was %d", res)
	}
}

func TestLongestStreakF(t *testing.T) {
	res := longestStreak(nil)
	if res != 0 {
		t.Errorf("expected result to be 0, was %d", res)
	}
}

func TestLongestStreak_F(t *testing.T) {
	res := longestStreakRecursive(nil, 0, 0)
	if res != 0 {
		t.Errorf("expected result to be 0, was %d", res)
	}
}