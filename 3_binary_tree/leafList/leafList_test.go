package main

import "testing"

func TestLeafListDFI_A(t *testing.T) {
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

	res := leafListDFI(a)
	expected := []interface{}{"d", "e", "f"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFR_A(t *testing.T) {
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

	res := leafListDFR(a)
	expected := []interface{}{"d", "e", "f"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFI_B(t *testing.T) {
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

	res := leafListDFI(a)
	expected := []interface{}{"d", "g", "h"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFR_B(t *testing.T) {
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

	res := leafListDFR(a)
	expected := []interface{}{"d", "g", "h"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFI_C(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 54}
	d := &Node{Val: 20}
	e := &Node{Val: 15}
	f := &Node{Val: 1}
	g := &Node{Val: 3}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	e.Left = f
	e.Right = g

	res := leafListDFI(a)
	expected := []interface{}{20, 1, 3, 54}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFR_C(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 54}
	d := &Node{Val: 20}
	e := &Node{Val: 15}
	f := &Node{Val: 1}
	g := &Node{Val: 3}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	e.Left = f
	e.Right = g

	res := leafListDFR(a)
	expected := []interface{}{20, 1, 3, 54}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFI_D(t *testing.T) {
	x := &Node{Val: "x"}
	
	res := leafListDFI(x)
	expected := []interface{}{"x"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFR_D(t *testing.T) {
	x := &Node{Val: "x"}
	
	res := leafListDFR(x)
	expected := []interface{}{"x"}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFI_E(t *testing.T) {
	res := leafListDFI(nil)
	expected := []interface{}{}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLeafListDFR_E(t *testing.T) {
	res := leafListDFI(nil)
	expected := []interface{}{}
	if !EqualInterfaceSlice(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func EqualInterfaceSlice(a, b []interface{}) bool {
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