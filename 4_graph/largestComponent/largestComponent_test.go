package main

import "testing"

func TestLargestComponent_A(t *testing.T) {
	g := Graph{
		0: []int{8,1,5},
		1: []int{0},
		5: []int{0,8},
		8: []int{0,5},
		2: []int{3,4},
		3: []int{2,4},
		4: []int{3,2},
	}
	res := largestComponent(g)
	expected := 4
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLargestComponent_B(t *testing.T) {
	g := Graph{
		1: []int{2},
		2: []int{1,8},
		6: []int{7},
		9: []int{8},
		7: []int{6,8},
		8: []int{9,7,2},
	}
	res := largestComponent(g)
	expected := 6
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLargestComponent_C(t *testing.T) {
	g := Graph{
		3: []int{},
		4: []int{6},
		6: []int{4,5,7,8},
		8: []int{6},
		7: []int{6},
		5: []int{6},
		1: []int{2},
		2: []int{1},
	}
	res := largestComponent(g)
	expected := 5
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLargestComponent_D(t *testing.T) {
	g := Graph{}
	res := largestComponent(g)
	expected := 0
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestLargestComponent_E(t *testing.T) {
	g := Graph{
		0: []int{4,7},
		1: []int{},
		2: []int{},
		3: []int{6},
		4: []int{0},
		6: []int{3},
		7: []int{0},
		8: []int{},
	}
	res := largestComponent(g)
	expected := 3
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}