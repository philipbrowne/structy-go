package main

import "testing"

func TestPathFinderA_A(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := pathFinderA(a, "e")
	if !Equal(res, []interface{}{"a", "b", "e"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "e"}, res)
	}
}

func TestPathFinderB_A(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := pathFinderB(a, "e")
	if !Equal(res, []interface{}{"a", "b", "e"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "e"}, res)
	}
}

func TestPathFinderA_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := pathFinderA(a, "p")
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
}

func TestPathFinderB_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := pathFinderB(a, "p")
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
}

func TestPathFinderA_C(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := pathFinderA(a, "c")
	if !Equal(res, []interface{}{"a", "c"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "c"}, res)
	}
}

func TestPathFinderB_C(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := pathFinderB(a, "c")
	if !Equal(res, []interface{}{"a", "c"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "c"}, res)
	}
}

func TestPathFinderA_D(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := pathFinderA(a, "h")
	if !Equal(res, []interface{}{"a", "c", "f", "h"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "c", "f", "h"}, res)
	}
}

func TestPathFinderB_D(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := pathFinderB(a, "h")
	if !Equal(res, []interface{}{"a", "c", "f", "h"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "c", "f", "h"}, res)
	}
}

func TestPathFinderA_E(t *testing.T) {
	x := &Node{Val: "x"}

	res := pathFinderA(x, "x")
	if !Equal(res, []interface{}{"x"}) {
		t.Errorf("expected %v, got %v", []interface{}{"x"}, res)
	}
}

func TestPathFinderB_E(t *testing.T) {
	x := &Node{Val: "x"}

	res := pathFinderB(x, "x")
	if !Equal(res, []interface{}{"x"}) {
		t.Errorf("expected %v, got %v", []interface{}{"x"}, res)
	}
}

func TestPathFinderA_F(t *testing.T) {
	res := pathFinderA(nil, "x")
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
	
}

func TestPathFinderB_F(t *testing.T) {
	res := pathFinderB(nil, "x")
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
}

func TestPathFinderA_G(t *testing.T) {
	root := &Node{Val: 0}
	curr := root
	for i := 1; i <= 6000; i++ {
		curr.Right = &Node{Val: i}
		curr = curr.Right
	}
	nums := []interface{}{}
	for i := 0; i <= 3451; i++ {
		nums = append(nums, i)
	}
	res := pathFinderA(root, 3451)
	if !Equal(res, nums) {
		t.Errorf("expected %v, got %v", nums, res)
	}
}

func TestPathFinderB_G(t *testing.T) {
	root := &Node{Val: 0}
	curr := root
	for i := 1; i <= 6000; i++ {
		curr.Right = &Node{Val: i}
		curr = curr.Right
	}
	nums := []interface{}{}
	for i := 0; i <= 3451; i++ {
		nums = append(nums, i)
	}
	res := pathFinderB(root, 3451)
	if !Equal(res, nums) {
		t.Errorf("expected %v, got %v", nums, res)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []interface{}) bool {
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
