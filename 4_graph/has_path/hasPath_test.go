package main

import "testing"

func TestHasPathDFI_A(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFI(g, "f", "k")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathBFI_A(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathBFI(g, "f", "k")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFR_A(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFR(g, "f", "k")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFI_B(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFI(g, "f", "j")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathBFI_B(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathBFI(g, "f", "j")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFR_B(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFR(g, "f", "j")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFI_C(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFI(g, "i", "h")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathBFI_C(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathBFI(g, "i", "h")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFR_C(t *testing.T) {
	g := Graph{
		"f": []string{"g", "i"},
		"g": []string{"h"},
		"h": []string{},
		"i": []string{"g", "k"},
		"j": []string{"i"},
		"k": []string{},
	}

	res := hasPathDFR(g, "i", "h")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFI_D(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathDFI(g, "v", "w")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathBFI_D(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathBFI(g, "v", "w")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFR_D(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathDFR(g, "v", "w")
	expected := true
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFI_E(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathDFI(g, "v", "z")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathBFI_E(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathBFI(g, "v", "z")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestHasPathDFR_E(t *testing.T) {
	g := Graph{
		"v": []string{"x", "w"},
		"w": []string{},
		"x": []string{},
		"y": []string{"z"},
		"z": []string{},
	}

	res := hasPathDFR(g, "v", "z")
	expected := false
	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}