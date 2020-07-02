//Implementation of Wiener's attack https://en.wikipedia.org/wiki/Wiener%27s_attack
package wiener

import (
	"math/big"

	"github.com/vveiln/crypto/wiener/fraction"
)

func solveQuadraticEquation(phi, n *big.Int) (*big.Int, *big.Int) {
	b := new(big.Int).Sub(n, phi)
	b.Add(b, big.NewInt(1))
	b.Neg(b)
	c := new(big.Int).Set(n)

	discr := new(big.Int).Mul(b, b)
	discr.Sub(discr, new(big.Int).Mul(big.NewInt(4), c))
	if discr.Cmp(big.NewInt(0)) == -1 {
		return nil, nil
	}
	discr.Sqrt(discr)

	x1 := new(big.Int).Neg(b)
	x1.Add(x1, discr)
	x1.Rsh(x1, 1)

	x2 := new(big.Int).Neg(b)
	x2.Sub(x2, discr)
	x2.Rsh(x2, 1)

	return x1, x2
}

func FindKey(e, n *big.Int) *big.Int {
	cf := fraction.NewContinuedFraction(fraction.NewFraction(e, n))
	convergents := cf.Convergents()
	for i := 0; i < len(convergents); i++ {
		conv := convergents[i]

		//Check if we can compute phi
		if conv.Numerator().Cmp(big.NewInt(0)) == 0 {
			continue
		}
		probable_phi := new(big.Int).Mul(e, conv.Denominator())
		probable_phi.Sub(probable_phi, big.NewInt(1))

		//Checks if probable_phi is integer
		divisible := new(big.Int).Mod(probable_phi, conv.Numerator())
		if divisible.Cmp(big.NewInt(0)) == 0 {

			//Integer phi candidate
			probable_phi := probable_phi.Div(probable_phi, conv.Numerator())
			p, q := solveQuadraticEquation(probable_phi, n)
			if p != nil && q != nil && new(big.Int).Mul(p, q).Cmp(n) == 0 {
				return conv.Denominator()
			}
		}
	}

	return nil
}
