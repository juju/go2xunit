package mmath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	x, y := 1, 2
	z := Add(x, y)
	if z != x+y {
		t.Fatalf("%d + %d != %d\n", x, y, x+y)
	}
}

func TestSub(t *testing.T) {
	x, y := 1, 2
	z := Sub(x, y)
	if z != x-y {
		t.Fatalf("%d-%d != %d\n", x, y, x-y)
	}
}

func TestMul(t *testing.T) {
	x, y := 2, 3
	z := Mul(x, y)
	if z != x*y {
		t.Fatalf("%d*%d != %d\n", x, y, x*y)
	}
}

func TestDiv(t *testing.T) {
	x, y := 8, 2
	z := Div(x, y)
	if z != x/y {
		t.Fatalf("%d/%d != %d\n", x, y, x/y)
	}
}

func TestDivInt(t *testing.T) {
	x, y := 2, 3
	z := Div(x, y)
	if z != 0 {
		t.Fatalf("%d/%d != %d\n", x, y, 0)
	}
}
