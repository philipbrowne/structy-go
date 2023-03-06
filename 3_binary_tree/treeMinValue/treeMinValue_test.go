package main

import "testing"

func TestTreeMinValDFI_A(t *testing.T) {
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

	res := treeMinValueDFI(a)
	if res != -2 {
		t.Errorf("expected %v, got %v", -2, res)
	}
}
func TestTreeMinValDFI_B(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 3}
	d := &Node{Val: 4}
	e := &Node{Val: 14}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeMinValueDFI(a)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeMinValDFI_C(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: -4}
	f := &Node{Val: -13}
	g := &Node{Val: -2}
	h := &Node{Val: -2}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := treeMinValueDFI(a)
	if res != -13 {
		t.Errorf("expected %v, got %v", -13, res)
	}
}

func TestTreeMinValDFI_D(t *testing.T) {
	a := &Node{Val: 42}

	res := treeMinValueDFI(a)
	if res != 42 {
		t.Errorf("expected %v, got %v", 42, res)
	}
}


func TestTreeMinValDFR_A(t *testing.T) {
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

	res := treeMinValueDFR(a)
	if res != -2 {
		t.Errorf("expected %v, got %v", -2, res)
	}
}
func TestTreeMinValDFR_B(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 3}
	d := &Node{Val: 4}
	e := &Node{Val: 14}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeMinValueDFR(a)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeMinValDFR_C(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: -4}
	f := &Node{Val: -13}
	g := &Node{Val: -2}
	h := &Node{Val: -2}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := treeMinValueDFR(a)
	if res != -13 {
		t.Errorf("expected %v, got %v", -13, res)
	}
}

func TestTreeMinValDFR_D(t *testing.T) {
	a := &Node{Val: 42}

	res := treeMinValueDFR(a)
	if res != 42 {
		t.Errorf("expected %v, got %v", 42, res)
	}
}

func TestTreeMinValBF_A(t *testing.T) {
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

	res := treeMinValueBF(a)
	if res != -2 {
		t.Errorf("expected %v, got %v", -2, res)
	}
}
func TestTreeMinValBF_B(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 3}
	d := &Node{Val: 4}
	e := &Node{Val: 14}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeMinValueBF(a)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeMinValBF_C(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: -4}
	f := &Node{Val: -13}
	g := &Node{Val: -2}
	h := &Node{Val: -2}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := treeMinValueBF(a)
	if res != -13 {
		t.Errorf("expected %v, got %v", -13, res)
	}
}

func TestTreeMinValBF_D(t *testing.T) {
	a := &Node{Val: 42}

	res := treeMinValueBF(a)
	if res != 42 {
		t.Errorf("expected %v, got %v", 42, res)
	}
}
