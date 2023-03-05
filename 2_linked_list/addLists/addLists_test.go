package main

import "testing"

func TestAddListsIterativeA(t *testing.T) {
	a1 := &Node{Val: 1}
	a2 := &Node{Val: 2}
	a3 := &Node{Val: 6}
	a1.Next = a2
	a2.Next = a3

	b1 := &Node{Val: 4}
	b2 := &Node{Val: 5}
	b3 := &Node{Val: 3}
	b1.Next = b2
	b2.Next = b3

	res := addListsIterative(a1, b1)
	if res.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Val)
	}
	if res.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Val)
	}
	if res.Next.Next.Val != 9 {
		t.Errorf("expected %v, got %v", 9, res.Next.Next.Val)
	}
	if res.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next)
	}
}

func TestAddListsIterativeB(t *testing.T) {
	a1 := &Node{Val: 1}
	a2 := &Node{Val: 4}
	a3 := &Node{Val: 5}
	a4 := &Node{Val: 7}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4

	b1 := &Node{Val: 2}
	b2 := &Node{Val: 3}
	b1.Next = b2

	res := addListsIterative(a1, b1)
	if res.Val != 3 {
		t.Errorf("expected %v, got %v", 3, res.Val)
	}
	if res.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Val)
	}
	if res.Next.Next.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Next.Next.Val)
	}
	if res.Next.Next.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Next.Next.Val)
	}
	if res.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next.Next)
	}
}

func TestAddListsIterativeC(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 3}
	a1.Next = a2

	b1 := &Node{Val: 7}
	b2 := &Node{Val: 4}
	b1.Next = b2

	res := addListsIterative(a1, b1)
	if res.Val != 6 {
		t.Errorf("expected %v, got %v", 6, res.Val)
	}
	if res.Next.Val != 8 {
		t.Errorf("expected %v, got %v", 8, res.Next.Val)
	}
	if res.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next)
	}
}

func TestAddListsIterativeD(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 8}
	a1.Next = a2

	b1 := &Node{Val: 7}
	b2 := &Node{Val: 4}
	b1.Next = b2

	res := addListsIterative(a1, b1)
	if res.Val != 6 {
		t.Errorf("expected %v, got %v", 6, res.Val)
	}
	if res.Next.Val != 3 {
		t.Errorf("expected %v, got %v", 3, res.Next.Val)
	}
	if res.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, res.Next.Next.Val)
	}
	if res.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next)
	}
}

func TestAddListsIterativeE(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 9}
	a3 := &Node{Val: 9}
	a1.Next = a2
	a2.Next = a3

	b1 := &Node{Val: 6}

	res := addListsIterative(a1, b1)
	if res.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Val)
	}
	if res.Next.Val != 0 {
		t.Errorf("expected %v, got %v", 0, res.Next.Val)
	}
	if res.Next.Next.Val != 0 {
		t.Errorf("expected %v, got %v", 0, res.Next.Next.Val)
	}
	if res.Next.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, res.Next.Next.Next.Val)
	}
	if res.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next.Next)
	}
}

func TestAddListsRecursiveA(t *testing.T) {
	a1 := &Node{Val: 1}
	a2 := &Node{Val: 2}
	a3 := &Node{Val: 6}
	a1.Next = a2
	a2.Next = a3

	b1 := &Node{Val: 4}
	b2 := &Node{Val: 5}
	b3 := &Node{Val: 3}
	b1.Next = b2
	b2.Next = b3

	res := addListsRecursive(a1, b1)
	if res.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Val)
	}
	if res.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Val)
	}
	if res.Next.Next.Val != 9 {
		t.Errorf("expected %v, got %v", 9, res.Next.Next.Val)
	}
	if res.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next)
	}
}

func TestAddListsRecursiveB(t *testing.T) {
	a1 := &Node{Val: 1}
	a2 := &Node{Val: 4}
	a3 := &Node{Val: 5}
	a4 := &Node{Val: 7}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4

	b1 := &Node{Val: 2}
	b2 := &Node{Val: 3}
	b1.Next = b2

	res := addListsRecursive(a1, b1)
	if res.Val != 3 {
		t.Errorf("expected %v, got %v", 3, res.Val)
	}
	if res.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Val)
	}
	if res.Next.Next.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Next.Next.Val)
	}
	if res.Next.Next.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, res.Next.Next.Next.Val)
	}
	if res.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next.Next)
	}
}

func TestAddListsRecursiveC(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 3}
	a1.Next = a2

	b1 := &Node{Val: 7}
	b2 := &Node{Val: 4}
	b1.Next = b2

	res := addListsRecursive(a1, b1)
	if res.Val != 6 {
		t.Errorf("expected %v, got %v", 6, res.Val)
	}
	if res.Next.Val != 8 {
		t.Errorf("expected %v, got %v", 8, res.Next.Val)
	}
	if res.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next)
	}
}

func TestAddListsRecursiveD(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 8}
	a1.Next = a2

	b1 := &Node{Val: 7}
	b2 := &Node{Val: 4}
	b1.Next = b2

	res := addListsRecursive(a1, b1)
	if res.Val != 6 {
		t.Errorf("expected %v, got %v", 6, res.Val)
	}
	if res.Next.Val != 3 {
		t.Errorf("expected %v, got %v", 3, res.Next.Val)
	}
	if res.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, res.Next.Next.Val)
	}
	if res.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next)
	}
}

func TestAddListsRecursiveE(t *testing.T) {
	a1 := &Node{Val: 9}
	a2 := &Node{Val: 9}
	a3 := &Node{Val: 9}
	a1.Next = a2
	a2.Next = a3

	b1 := &Node{Val: 6}

	res := addListsRecursive(a1, b1)
	if res.Val != 5 {
		t.Errorf("expected %v, got %v", 5, res.Val)
	}
	if res.Next.Val != 0 {
		t.Errorf("expected %v, got %v", 0, res.Next.Val)
	}
	if res.Next.Next.Val != 0 {
		t.Errorf("expected %v, got %v", 0, res.Next.Next.Val)
	}
	if res.Next.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, res.Next.Next.Next.Val)
	}
	if res.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, res.Next.Next.Next.Next)
	}
}