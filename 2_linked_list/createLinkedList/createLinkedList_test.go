package main

import "testing"

func TestCreateLinkedListIterativeA(t *testing.T) {
	head := createLinkedListIterative([]interface{}{"h", "e", "y"})
	if head.Val != "h" {
		t.Errorf("expected %v, got %v", "h", head.Val)
	}
	if head.Next.Val != "e" {
		t.Errorf("expected %v, got %v", "e", head.Next.Val)
	}
	if head.Next.Next.Val != "y" {
		t.Errorf("expected %v, got %v", "y", head.Next.Next.Val)
	}
	if head.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next.Next.Next)
	}
}
func TestCreateLinkedListIterativeB(t *testing.T) {
	head := createLinkedListIterative([]interface{}{1, 7, 1, 8})
	if head.Val != 1 {
		t.Errorf("expected %v, got %v", 1, head.Val)
	}
	if head.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, head.Next.Val)
	}
	if head.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, head.Next.Next.Val)
	}
	if head.Next.Next.Next.Val != 8 {
		t.Errorf("expected %v, got %v", 8, head.Next.Next.Next.Val)
	}
	if head.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next.Next.Next)
	}
}

func TestCreateLinkedListIterativeC(t *testing.T) {
	head := createLinkedListIterative([]interface{}{"a"})
	if head.Val != "a" {
		t.Errorf("expected %v, got %v", "a", head.Val)
	}
	if head.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next)
	}
}

func TestCreateLinkedListIterativeD(t *testing.T) {
	head := createLinkedListIterative([]interface{}{})
	if head != nil {
		t.Errorf("expected %v, got %v", nil, head)
	}
}

func TestCreateLinkedListRecursiveA(t *testing.T) {
	head := createLinkedListRecursive([]interface{}{"h", "e", "y"})
	if head.Val != "h" {
		t.Errorf("expected %v, got %v", "h", head.Val)
	}
	if head.Next.Val != "e" {
		t.Errorf("expected %v, got %v", "e", head.Next.Val)
	}
	if head.Next.Next.Val != "y" {
		t.Errorf("expected %v, got %v", "y", head.Next.Next.Val)
	}
	if head.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next.Next.Next)
	}
}
func TestCreateLinkedListRecursiveB(t *testing.T) {
	head := createLinkedListRecursive([]interface{}{1, 7, 1, 8})
	if head.Val != 1 {
		t.Errorf("expected %v, got %v", 1, head.Val)
	}
	if head.Next.Val != 7 {
		t.Errorf("expected %v, got %v", 7, head.Next.Val)
	}
	if head.Next.Next.Val != 1 {
		t.Errorf("expected %v, got %v", 1, head.Next.Next.Val)
	}
	if head.Next.Next.Next.Val != 8 {
		t.Errorf("expected %v, got %v", 8, head.Next.Next.Next.Val)
	}
	if head.Next.Next.Next.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next.Next.Next)
	}
}

func TestCreateLinkedListRecursiveC(t *testing.T) {
	head := createLinkedListRecursive([]interface{}{"a"})
	if head.Val != "a" {
		t.Errorf("expected %v, got %v", "a", head.Val)
	}
	if head.Next != nil {
		t.Errorf("expected %v, got %v", nil, head.Next)
	}
}

func TestCreateLinkedListRecursiveD(t *testing.T) {
	head := createLinkedListRecursive([]interface{}{})
	if head != nil {
		t.Errorf("expected %v, got %v", nil, head)
	}
}