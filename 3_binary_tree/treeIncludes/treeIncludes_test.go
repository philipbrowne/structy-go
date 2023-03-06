package main

import "testing"

func TestTreeIncludesDFI_A(t *testing.T) {
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

	res := treeIncludesDFI(a, "e")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFI_B(t *testing.T) {
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

	res := treeIncludesDFI(a, "a")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFI_C(t *testing.T) {
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

	res := treeIncludesDFI(a, "n")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}
func TestTreeIncludesDFI_D(t *testing.T) {
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

	res := treeIncludesDFI(a, "f")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFI_E(t *testing.T) {
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

	res := treeIncludesDFI(a, "p")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesDFI_F(t *testing.T) {
	res := treeIncludesDFI(nil, "b")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesDFR_A(t *testing.T) {
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

	res := treeIncludesDFR(a, "e")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFR_B(t *testing.T) {
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

	res := treeIncludesDFR(a, "a")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFR_C(t *testing.T) {
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

	res := treeIncludesDFR(a, "n")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesDFR_D(t *testing.T) {
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

	res := treeIncludesDFR(a, "f")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesDFR_E(t *testing.T) {
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

	res := treeIncludesDFR(a, "p")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesDFR_F(t *testing.T) {
	res := treeIncludesDFR(nil, "b")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesBF_A(t *testing.T) {
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

	res := treeIncludesBF(a, "e")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesBF_B(t *testing.T) {
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

	res := treeIncludesBF(a, "a")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesBF_C(t *testing.T) {
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

	res := treeIncludesBF(a, "n")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesBF_D(t *testing.T) {
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

	res := treeIncludesBF(a, "f")
	if res != true {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestTreeIncludesBF_E(t *testing.T) {
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

	res := treeIncludesBF(a, "p")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}

func TestTreeIncludesBF_F(t *testing.T) {
	res := treeIncludesBF(nil, "b")
	if res != false {
		t.Errorf("expected %v, got %v", false, res)
	}
}
