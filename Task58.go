package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Task58() {
	var q = big.NewInt(0)
	var b = big.NewInt(0)
	var temp = big.NewInt(0)
	var y = big.NewInt(0)
	q.SetString("335062023296420808191071248367701059461",10)
	n,r,Oracle := Task57forTask58()
	//b
	b.Sub(q,big.NewInt(1))
	b.Div(b,r)
	//y,y'
	y = Oracle.pk
	temp.Sub(q,n)
	temp.Exp(Oracle.g, temp, Oracle.p)
	y.Mul(y,temp)
	y.Mod(y,Oracle.p)
	//g'
	temp = Oracle.g
	temp.Exp(temp, r, Oracle.p)
	result, err := Kangaroo(big.NewInt(0),b, temp, y, Oracle.p)
	if err != nil {
		fmt.Println(err.Error())
	}
	result.Mul(result,r)
	result.Add(result, n)
	fmt.Println(Oracle.sk)
	fmt.Println(result)
}

func Kangaroo(a,b,g,y,p *big.Int) (*big.Int,error){
	//k
	//var k = big.NewInt(21)
	fmt.Println("Task58:")
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(5), nil).Sub(max, big.NewInt(1))
	k, err := rand.Int(rand.Reader, max)
	if err != nil {}
	//trained kangaroo
	var xT = big.NewInt(0)
	var yT = big.NewInt(0)
	yT.Exp(g,b, p)
	N := GetN(k,p)
	var i = big.NewInt(1)
	var temp = big.NewInt(0)
	fmt.Println("trained kangaroo is started")
	fmt.Println("k is:", k)
	fmt.Println("N is:", N)
	for i.Cmp(N)<1{
		temp = f(yT,k,p)
		xT.Add(xT,temp)
		exp := temp
		temp.Exp(g,exp,p)
		yT.Mul(yT, temp)
		yT.Mod(yT, p)
		i.Add(i,big.NewInt(1))
	}
	//wild kangaroo
	var xW = big.NewInt(0)
	var yW = big.NewInt(0)
	yW = y
	temp.Sub(b,a)
	temp.Add(temp,xT)
	var temp1 = big.NewInt(0)
	fmt.Println("wild kangaroo is started")
	for xW.Cmp(temp)==-1{
		temp1 = f(yW,k,p)
		xW.Add(xW,temp1)
		exp := temp1
		temp1.Exp(g,exp,p)
		yW.Mul(yW,temp1)
		yW.Mod(yW, p)
		if yW.Cmp(yT)==0{
			temp1.Add(b,xT)
			temp1.Sub(temp1,xW)
			return temp1,nil
		}
	}
	return nil, fmt.Errorf("k is bad")
}

func GetN(k,p *big.Int) *big.Int{
	var Sum = big.NewInt(2)
	Sum.Exp(Sum,k,p)
	Sum.Sub(Sum, big.NewInt(1))
	Sum.Div(Sum,k)
	Sum.Mul(Sum, big.NewInt(4))
	return Sum
}

func f(y,k,p *big.Int) *big.Int{
	var temp = big.NewInt(0)
	temp.Mod(y, k)
	temp.Exp(big.NewInt(2),temp,p)
	return temp
}

func Task57forTask58() (*big.Int,*big.Int, *DH) {
	//fmt.Println("Task57:")
	var Oracle = DH {p : big.NewInt(2017), g:big.NewInt(127)}
	Oracle.p.SetString("11470374874925275658116663507232161402086650258453896274534991676898999262641581519101074740642369848233294239851519212341844337347119899874391456329785623",10)
	Oracle.g.SetString("622952335333961296978159266084741085889881358738459939978290179936063635566740258555167783009058567397963466103140082647486611657350811560630587013183357",10)
	Oracle.skGen()
	Oracle.pkGen()
	rMass := make([]*big.Int,0)
	hMass := make([]*big.Int,0)
	bMass := make([]*big.Int,0)
	var j = big.NewInt(0)
	j.SetString("34233586850807404623475048381328686211071196701374230492615844865929237417097514638999377942356150481334217896204702", 10)
	var q = big.NewInt(0)
	q.SetString("335062023296420808191071248367701059461",10)
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
			h := hFromR(div, Oracle.p)
			hMass = append(hMass, &h)
			Oracle.SharedSecretGen(&h)
			bMass = append(bMass, BruteForce(Oracle.ss, Oracle.p, big.NewInt(div), &h))
			rMul.Mul(rMul, big.NewInt(div))
			if rMul.Cmp(q)==1 {
				break
			}
			i++
		}
		div--
	}
	//fmt.Println(rMass)
	//fmt.Println(hMass)
	//fmt.Println(bMass)
	solution,err := crt(bMass, rMass)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(Oracle.sk)
	//fmt.Println(solution)
	return solution, rMul, &Oracle
}