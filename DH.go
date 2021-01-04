package main

import "math/big"


type DH struct {
	p *big.Int
	g *big.Int
	sk *big.Int
	pk *big.Int
	ss *big.Int
}

func (DHthis *DH) SharedSecretGen(B *big.Int) {
	var temp = big.NewInt(2)
	DHthis.ss = temp.Exp(B,DHthis.sk,nil)
	DHthis.ss = temp.Mod(DHthis.ss, DHthis.p)
}

func (DHthis *DH) pkGen() {
	DHthis.pk = DHthis.g.Exp(DHthis.g,DHthis.sk,nil)
	DHthis.pk = DHthis.g.Mod(DHthis.pk,DHthis.p)
}