package main

import "testing"

func TestMergeListsIterative(t *testing.T) {
	a := &Node{
		Val: 5,
	}
	b := &Node{
		Val: 7,
	}
	c := &Node{
		Val: 10,
	}
	d := &Node{
		Val: 12,
	}
	e := &Node{
		Val:20,
	}
	f := &Node{
		Val: 28,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	
	p := &Node{
		Val: 6,
	}
	q := &Node{
		Val: 8,
	}
	r := &Node{
		Val: 9,
	}
	s := &Node{
		Val: 25,
	}
	p.Next = q
	q.Next = r
	r.Next = s
	mergeListsIterative(a, p)
	if a.Next != p {
		t.Errorf("expected a.Next to p, was %v", a.Next)
	}
	if p.Next != b {
		t.Errorf("expected p.Next to be b, was %v", p.Next)
	}
	if e.Next != s {
		t.Errorf("expected e.Next to be s, was %v", e.Next)
	}
	if s.Next != f {
		t.Errorf("expected s.Next to be f, was %v", s.Next)
	}
}

func TestMergeListsRecursive(t *testing.T) {
	a := &Node{
		Val: 5,
	}
	b := &Node{
		Val: 7,
	}
	c := &Node{
		Val: 10,
	}
	d := &Node{
		Val: 12,
	}
	e := &Node{
		Val:20,
	}
	f := &Node{
		Val: 28,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	
	p := &Node{
		Val: 6,
	}
	q := &Node{
		Val: 8,
	}
	r := &Node{
		Val: 9,
	}
	s := &Node{
		Val: 25,
	}
	p.Next = q
	q.Next = r
	r.Next = s
	mergeListsRecursive(a, p)
	if a.Next != p {
		t.Errorf("expected a.Next to p, was %v", a.Next)
	}
	if p.Next != b {
		t.Errorf("expected p.Next to be b, was %v", p.Next)
	}
	if e.Next != s {
		t.Errorf("expected e.Next to be s, was %v", e.Next)
	}
	if s.Next != f {
		t.Errorf("expected s.Next to be f, was %v", s.Next)
	}
}