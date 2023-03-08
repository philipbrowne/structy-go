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

// TODO - Add more tests