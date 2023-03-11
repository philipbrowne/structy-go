package main

import "testing"

func TestUndirectedPathDFI_A(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFI(edges, "j", "m")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_A(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathBFI(edges, "j", "m")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_A(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFR(edges, "j", "m")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_B(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFI(edges, "m", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_B(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathBFI(edges, "m", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_B(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFR(edges, "m", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_C(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFI(edges, "l", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_C(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathBFI(edges, "l", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_C(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFR(edges, "l", "j")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_D(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFI(edges, "k", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_D(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathBFI(edges, "k", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_D(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFR(edges, "k", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_E(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFI(edges, "i", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_E(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathBFI(edges, "i", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_E(t *testing.T) {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	
	res := undirectedPathDFR(edges, "i", "o")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_F(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFI(edges, "a", "b")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_F(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathBFI(edges, "a", "b")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_F(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFR(edges, "a", "b")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_G(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFI(edges, "a", "c")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_G(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathBFI(edges, "a", "c")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_G(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFR(edges, "a", "c")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_H(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFI(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_H(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathBFI(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_H(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFR(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_I(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFI(edges, "r", "b")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_I(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathBFI(edges, "r", "b")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_I(t *testing.T) {
	edges := [][]string{
		{"b", "a"},
		{"c", "a"},
		{"b", "c"},
		{"q", "r"},
		{"q", "s"},
		{"q", "u"},
		{"q", "t"},
	}
	
	res := undirectedPathDFR(edges, "r", "b")
	expected := false
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFI_J(t *testing.T) {
	edges := [][]string{
		{"s", "r"},
		{"t", "q"},
		{"q", "r"},
	}
	
	res := undirectedPathDFI(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathBFI_J(t *testing.T) {
	edges := [][]string{
		{"s", "r"},
		{"t", "q"},
		{"q", "r"},
	}
	
	res := undirectedPathBFI(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestUndirectedPathDFR_J(t *testing.T) {
	edges := [][]string{
		{"s", "r"},
		{"t", "q"},
		{"q", "r"},
	}
	
	res := undirectedPathDFR(edges, "r", "t")
	expected := true
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}