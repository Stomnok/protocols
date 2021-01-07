package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Task57() {
	fmt.Println("Task57:")
	var Oracle = DH {p : big.NewInt(2017), g:big.NewInt(127)}
	Oracle.p.SetString("7199773997391911030609999317773941274322764333428698921736339643928346453700085358802973900485592910475480089726140708102474957429903531369589969318716771",10)
	Oracle.g.SetString("4565356397095740655436854503483826832136106141639563487732438195343690437606117828318042418238184896212352329118608100083187535033402010599512641674644143",10)
	Oracle.skGen()
	rMass := make([]*big.Int,0)
	hMass := make([]*big.Int,0)
	bMass := make([]*big.Int,0)
	var j = big.NewInt(0)
	j.SetString("30477252323177606811760882179058908038824640750610513771646768011063128035873508507547741559514324673960576895059570", 10)
	var q = big.NewInt(0)
	q.SetString("236234353446506858198510045061214171961",10)
	var rMul = big.NewInt(1)
	var modResult = big.NewInt(0)
	var div int64 = 65536
	var i = 0
	var gcd = big.NewInt(0)
	var flagNotCoPrime = false
	for div>1 {
		modResult.Mod(j,big.NewInt(div))
		if modResult.Cmp(big.NewInt(0))==0 {
			for j:=0; j<i; j++{
				gcd.GCD(nil,nil,rMass[j], big.NewInt(div))
				if gcd.Cmp(big.NewInt(1)) != 0{
					flagNotCoPrime = true
					break
				}
			}
			if flagNotCoPrime{
				flagNotCoPrime = false
				div--
				continue
			}
			rMass = append(rMass, big.NewInt(div))
			temp := hFromR(div, Oracle.p)
			hMass = append(hMass, &temp)
			Oracle.SharedSecretGen(hMass[i])
			bMass = append(bMass, BruteForce(Oracle.ss, Oracle.p, big.NewInt(div), &temp))
			rMul.Mul(rMul, big.NewInt(div))
			if rMul.Cmp(q)==1 {
				break
			}
			i++
		}
		div--
	}
	solution,err := crt(bMass, rMass)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Oracle.sk)
	fmt.Println(solution)
}

func hFromR(r int64, p * big.Int) big.Int {
	//h := rand(1, p)^((p-1)/r) mod p
	h, err := rand.Int(rand.Reader, p)
	if err != nil {}
	var exp = big.NewInt(0)
	exp.Sub(p,big.NewInt(1))
	exp.Div(exp, big.NewInt(r))
	h.Exp(h,exp,p)
	return *h
}

func BruteForce(ss, p, r, h *big.Int) *big.Int{
	var result = big.NewInt(0)
	var x = big.NewInt(0)
	for x.Cmp(r) == -1{
			result.Exp(h,x,p)
			if result.Cmp(ss) == 0{
				return x
			}
			x.Add(x,big.NewInt(1))
	}
	return nil
}