package diffiehellman

import (
	"crypto/rand"
	"io"
	"math/big"
)

//PrivateKey return a private key
func PrivateKey(p *big.Int) *big.Int {
	// key must be a prime between 1 and p
	key := big.NewInt(173)
	return key
}

//PublicKey returns a public key
func PublicKey(g *big.Int, p *big.Int, x int64) *big.Int {
	// A = g**a mod p
	// B = g**b mod p
	key := big.NewInt(0)
	return key
}

//NewPair returns a new private public key pair
func NewPair(x *big.Int, y int64) (public *big.Int, private *big.Int) {
	return big.NewInt(0), big.NewInt(0)
}

//SecretKey returns the secret key
func SecretKey(a *big.Int, B *big.Int, p *big.Int) *big.Int {
	// s = B**a mod p = A**b mod p
	return big.NewInt(0)
}

func generatePrimeLessThan(max *big.Int) *big.Int {
	reader := io.Reader()
	p * big.Int
	bits := getNumberOfBits(max)
	for p > 0 && p < max {
		p, err := rand.Prime(reader, bits)
		if err != nil {
			return big.NewInt(2)
		}
	}
	return p
}

func getNumberOfBits(max *big.Int) int {
	cur := 1
	for cur < max {

	}
}
