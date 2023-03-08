package main

import "testing"

func TestTreeValueCount_DFR_A(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_DFR(a, 6)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeValueCount_DFI_A(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_DFI(a, 6)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeValueCount_BF_A(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_BF(a, 6)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestTreeValueCount_DFR_B(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_DFR(a, 12)
	if res != 2 {
		t.Errorf("expected %v, got %v", 2, res)
	}
}

func TestTreeValueCount_DFI_B(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_DFI(a, 12)
	if res != 2 {
		t.Errorf("expected %v, got %v", 2, res)
	}
}

func TestTreeValueCount_BF_B(t *testing.T) {
	a := &Node{Val: 12}
	b := &Node{Val: 6}
	c := &Node{Val: 6}
	d := &Node{Val: 4}
	e := &Node{Val: 6}
	f := &Node{Val: 12}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := treeValueCount_BF(a, 12)
	if res != 2 {
		t.Errorf("expected %v, got %v", 2, res)
	}
}

func TestTreeValueCount_DFR_C(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_DFR(a, 1)
	if res != 4 {
		t.Errorf("expected %v, got %v", 4, res)
	}
}

func TestTreeValueCount_DFI_C(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_DFI(a, 1)
	if res != 4 {
		t.Errorf("expected %v, got %v", 4, res)
	}
}

func TestTreeValueCount_BF_C(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_BF(a, 1)
	if res != 4 {
		t.Errorf("expected %v, got %v", 4, res)
	}
}

func TestTreeValueCount_DFR_D(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_DFR(a, 9)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeValueCount_DFI_D(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_DFI(a, 9)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeValueCount_BF_D(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 5}
	c := &Node{Val: 1}
	d := &Node{Val: 1}
	e := &Node{Val: 8}
	f := &Node{Val: 7}
	g := &Node{Val: 1}
	h := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	g.Right = h

	res := treeValueCount_BF(a, 9)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeValueCount_DFR_E(t *testing.T) {
	res := treeValueCount_DFR(nil, 42)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeValueCount_DFI_E(t *testing.T) {
	res := treeValueCount_DFI(nil, 42)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeValueCount_BF_E(t *testing.T) {
	res := treeValueCount_BF(nil, 42)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}