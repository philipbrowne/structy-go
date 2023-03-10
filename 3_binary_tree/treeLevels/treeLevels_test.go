package main

import "testing"

func TestTreeLevelsDFI_A(t *testing.T) {
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

	res := treeLevelsDFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsBFI_A(t *testing.T) {
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

	res := treeLevelsBFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFR_A(t *testing.T) {
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

	res := treeLevelsDFR(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFI_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}
	i := &Node{Val: "i"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right = h
	f.Left = i

	res := treeLevelsDFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsBFI_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}
	i := &Node{Val: "i"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right = h
	f.Left = i

	res := treeLevelsBFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFR_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}
	i := &Node{Val: "i"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right = h
	f.Left = i

	res := treeLevelsDFR(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFI_C(t *testing.T) {
	q := &Node{Val: "q"}
	r := &Node{Val: "r"}
	s := &Node{Val: "s"}
	tt := &Node{Val: "t"}
	u := &Node{Val: "u"}
	v := &Node{Val: "v"}

	q.Left = r
	q.Right = s
	r.Right = tt
	tt.Left = u
	u.Right = v

	res := treeLevelsDFI(q)
	expected := [][]interface{}{
		{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsBFI_C(t *testing.T) {
	q := &Node{Val: "q"}
	r := &Node{Val: "r"}
	s := &Node{Val: "s"}
	tt := &Node{Val: "t"}
	u := &Node{Val: "u"}
	v := &Node{Val: "v"}

	q.Left = r
	q.Right = s
	r.Right = tt
	tt.Left = u
	u.Right = v

	res := treeLevelsBFI(q)
	expected := [][]interface{}{
		{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFR_C(t *testing.T) {
	q := &Node{Val: "q"}
	r := &Node{Val: "r"}
	s := &Node{Val: "s"}
	tt := &Node{Val: "t"}
	u := &Node{Val: "u"}
	v := &Node{Val: "v"}

	q.Left = r
	q.Right = s
	r.Right = tt
	tt.Left = u
	u.Right = v

	res := treeLevelsDFR(q)
	expected := [][]interface{}{
		{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFI_D(t *testing.T) {
	res := treeLevelsDFI(nil)
	expected := [][]interface{}{}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsBFI_D(t *testing.T) {
	res := treeLevelsDFI(nil)
	expected := [][]interface{}{}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFR_D(t *testing.T) {
	res := treeLevelsDFI(nil)
	expected := [][]interface{}{}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func EqualNested(a, b [][]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !EqualInterface(v, b[i]) {
			return false
		}
	}
	return true
}

func EqualInterface(a, b []interface{}) bool {
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