package fraction

import (
	"math/big"
	"testing"
)

func TestFraction_Reduce(t *testing.T) {
	f := NewFraction(big.NewInt(22), big.NewInt(33))
	f.Reduce()
	expected := NewFraction(big.NewInt(2), big.NewInt(3))
	if !f.Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(2), big.NewInt(3))
	f.Reduce()
	expected = NewFraction(big.NewInt(2), big.NewInt(3))
	if !f.Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(0), big.NewInt(3))
	f.Reduce()
	expected = NewFraction(big.NewInt(0), big.NewInt(1))
	if !f.Equals(expected) {
		t.FailNow()
	}
}

func TestFraction_Equals(t *testing.T) {
	f := NewFraction(big.NewInt(0), big.NewInt(3))
	expected := NewFraction(big.NewInt(0), big.NewInt(1))
	if !f.Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(2), big.NewInt(3))
	expected = NewFraction(big.NewInt(4), big.NewInt(6))
	if !f.Equals(expected) {
		t.FailNow()
	}
}

func TestFraction_Inverse(t *testing.T) {
	f := NewFraction(big.NewInt(2), big.NewInt(22))
	expected := NewFraction(big.NewInt(22), big.NewInt(2))
	if !f.Inverse().Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(0), big.NewInt(22))
	if f.Inverse() != nil {
		t.FailNow()
	}
}

func TestFraction_GetIntegerPart(t *testing.T) {
	f := NewFraction(big.NewInt(2), big.NewInt(22))
	expected := big.NewInt(0)
	if f.GetIntegerPart().Cmp(expected) != 0 {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(23), big.NewInt(22))
	expected = big.NewInt(1)
	if f.GetIntegerPart().Cmp(expected) != 0 {
		t.FailNow()
	}
}

func TestFraction_GetFractionalPart(t *testing.T) {
	f := NewFraction(big.NewInt(2), big.NewInt(22))
	expected := NewFraction(big.NewInt(2), big.NewInt(22))
	if !f.GetFractionalPart().Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(23), big.NewInt(22))
	expected = NewFraction(big.NewInt(1), big.NewInt(22))
	if !f.GetFractionalPart().Equals(expected) {
		t.FailNow()
	}

	f = NewFraction(big.NewInt(23), big.NewInt(1))
	expected = NewFraction(big.NewInt(0), big.NewInt(1))
	if !f.GetFractionalPart().Equals(expected) {
		t.FailNow()
	}
}
