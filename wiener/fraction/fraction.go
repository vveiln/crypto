package fraction

import (
	"fmt"
	"math/big"
)

//Fraction is a representation of rational numbers
type Fraction struct {
	numerator   *big.Int
	denominator *big.Int
}

func NewFraction(n, d *big.Int) *Fraction {
	if d.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	return &Fraction{n, d}
}

func (f *Fraction) String() string {
	return fmt.Sprint("Fraction(", f.numerator.String(), ", ", f.denominator.String(), ")")
}

func (f *Fraction) Numerator() *big.Int {
	return f.numerator
}

func (f *Fraction) Denominator() *big.Int {
	return f.denominator
}

func (f *Fraction) Reduce() {
	gcd := new(big.Int).GCD(nil, nil, f.numerator, f.denominator)
	if gcd.Cmp(big.NewInt(1)) != 0 {
		f.numerator = new(big.Int).Quo(f.numerator, gcd)
		f.denominator = new(big.Int).Quo(f.denominator, gcd)
	}
}

func (f *Fraction) Equals(ff *Fraction) bool {
	cf := NewFraction(f.numerator, f.denominator)
	cff := NewFraction(ff.numerator, ff.denominator)
	cf.Reduce()
	cff.Reduce()
	return cf.numerator.Cmp(cff.numerator) == 0 && cf.denominator.Cmp(cff.denominator) == 0
}

//Inverse swaps the numerator and denominator and returns the result
//The current fraction is also set to the inverse value
func (f *Fraction) Inverse() *Fraction {
	if f.numerator.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	f.numerator, f.denominator = f.denominator, f.numerator
	return f
}

func (f *Fraction) GetIntegerPart() *big.Int {
	return new(big.Int).Quo(f.numerator, f.denominator)
}

func (f *Fraction) GetFractionalPart() *Fraction {
	intPart := f.GetIntegerPart()
	num := new(big.Int).Sub(f.numerator, new(big.Int).Mul(intPart, f.denominator))
	return NewFraction(num, f.denominator)
}

func (f *Fraction) IsZero() bool {
	return f.numerator.Cmp(big.NewInt(0)) == 0
}

//ContinuedFraction works for rational values only
type ContinuedFraction struct {
	Fraction []*big.Int
}

func NewContinuedFraction(r *Fraction) *ContinuedFraction {
	fraction := computeContinuedFraction(r)
	return &ContinuedFraction{Fraction: fraction}
}

func computeContinuedFraction(r *Fraction) []*big.Int {
	fraction := make([]*big.Int, 0, 1)
	fraction = append(fraction, r.GetIntegerPart())

	f := r.GetFractionalPart()
	if f.IsZero() {
		return fraction
	}

	fraction = append(fraction, computeContinuedFraction(f.Inverse())...)
	return fraction
}

//NextConvergent computes n-th convergent provided two previous convergents
func (cf ContinuedFraction) NextConvergent(c1, c2 *Fraction, n int64) *Fraction {
	num := new(big.Int).Mul(cf.Fraction[n], c1.numerator)
	num.Add(num, c2.numerator)

	den := new(big.Int).Mul(cf.Fraction[n], c1.denominator)
	den.Add(den, c2.denominator)

	return NewFraction(num, den)
}

//GetConvergent computes n-th convergent recursively
func (cf ContinuedFraction) GetConvergent(n int64) *Fraction {
	if n == 0 {
		return NewFraction(cf.Fraction[0], big.NewInt(1))
	}
	if n == 1 {
		a1 := cf.Fraction[1]
		a0 := cf.Fraction[0]
		a0a1 := new(big.Int).Mul(a0, a1)
		num := new(big.Int).Add(a0a1, big.NewInt(1))
		return NewFraction(num, a1)
	}

	return cf.NextConvergent(cf.GetConvergent(n-1), cf.GetConvergent(n-2), n)
}

func (cf ContinuedFraction) Convergents() []*Fraction {
	convergents := make([]*Fraction, 0, len(cf.Fraction))
	convergents = append(convergents, cf.GetConvergent(0), cf.GetConvergent(1))
	for i := 2; i < len(cf.Fraction); i++ {
		convergents = append(convergents, cf.NextConvergent(convergents[i-1], convergents[i-2], int64(i)))
	}
	return convergents
}
