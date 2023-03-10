package main

import "testing"

func TestLevelAveragesA(t *testing.T) {
	a := &Node{Val: 3}
	b := &Node{Val: 11}
	c := &Node{Val: 4}
	d := &Node{Val: 4}
	e := &Node{Val: -2}
	f := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := levelAverages(a)
	expected := []float64{3, 7.5, 1}
	if !EqualFloatSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLevelAveragesB(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 54}
	d := &Node{Val: 20}
	e := &Node{Val: 15}
	f := &Node{Val: 1}
	g := &Node{Val:3}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	e.Left = f
	e.Right = g

	res := levelAverages(a)
	expected := []float64{5, 32.5, 17.5, 2}
	if !EqualFloatSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLevelAveragesC(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: 0}
	f := &Node{Val: 45}
	g := &Node{Val:-1}
	h := &Node{Val:-2}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := levelAverages(a)
	expected := []float64{-1, -5.5, 14, -1.5}
	if !EqualFloatSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLevelAveragesD(t *testing.T) {
	q := &Node{Val: 13}
	r := &Node{Val: 4}
	s := &Node{Val: 2}
	tt := &Node{Val: 9}
	u := &Node{Val: 2}
	v := &Node{Val: 42}
	
	q.Left = r
	q.Right = s
	r.Right = tt
	tt.Left = u
	u.Right = v

	res := levelAverages(q)
	expected := []float64{13, 3, 9, 2, 42}
	if !EqualFloatSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLevelAveragesE(t *testing.T) {
	res := levelAverages(nil)
	expected := []float64{}
	if !EqualFloatSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}


func EqualFloatSlice(a, b []float64) bool {
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