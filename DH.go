package main

import (
	"crypto/rand"
	"math/big"
)


type DH struct {
	p *big.Int
	g *big.Int
	sk *big.Int
	pk *big.Int
	ss *big.Int
}

func (DHthis *DH) SharedSecretGen(B *big.Int) {
	DHthis.ss = big.NewInt(1)
	DHthis.ss.Exp(B, DHthis.sk, DHthis.p)
}

func (DHthis *DH) pkGen() {
	DHthis.pk = big.NewInt(1)
	DHthis.pk.Exp(DHthis.g, DHthis.sk, DHthis.p)
}

func (DHthis *DH) skGen() {
	DHthis.sk = big.NewInt(0)
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {}
	DHthis.sk.Mod(n,DHthis.p)
}