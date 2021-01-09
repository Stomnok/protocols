package main

import (
	"crypto/rand"
	//"crypto/rand"
	"fmt"
	"math/big"
)

func Task58() {
	fmt.Println("Task58:")
	var Oracle = DH {p : big.NewInt(2017), g:big.NewInt(127)}
	Oracle.p.SetString("11470374874925275658116663507232161402086650258453896274534991676898999262641581519101074740642369848233294239851519212341844337347119899874391456329785623",10)
	Oracle.g.SetString("622952335333961296978159266084741085889881358738459939978290179936063635566740258555167783009058567397963466103140082647486611657350811560630587013183357",10)
	Oracle.skGen()
	Oracle.pkGen()
	var j = big.NewInt(0)
	j.SetString("34233586850807404623475048381328686211071196701374230492615844865929237417097514638999377942356150481334217896204702", 10)
	var q = big.NewInt(0)
	q.SetString("335062023296420808191071248367701059461",10)
	var div int64 = 3
	var modResult = big.NewInt(0)
	for true{
		modResult.Mod(j,big.NewInt(div))
		if modResult.Cmp(big.NewInt(0))==0 {
			break
		}
		div++
	}
	//fmt.Println(div)
	h := hFromR(div, Oracle.p)
	Oracle.SharedSecretGen(&h)
	n := BruteForce(Oracle.ss, Oracle.p, big.NewInt(div), &h)
	var b = big.NewInt(0)
	var temp = big.NewInt(0)
	var y = big.NewInt(0)
	b.Sub(q,big.NewInt(1))
	b.Div(b,big.NewInt(div))
	y.Exp(Oracle.g, n, Oracle.p)//y = Oracle.pk
	temp.Sub(q,n)
	temp.Exp(Oracle.g, temp, Oracle.p)
	y.Mul(y,temp)
	y.Mod(y,Oracle.p)
	temp = Oracle.g
	temp.Exp(temp, big.NewInt(div),nil)
	result, err := Kangaroo(big.NewInt(0),b, temp, y, Oracle.p)
	if err != nil {
		fmt.Println(err.Error())
	}
	result.Mul(result,big.NewInt(div))
	result.Add(result, n)
	fmt.Println(Oracle.sk)
	fmt.Println(result)
}

func Kangaroo(a,b,g,y,p *big.Int) (*big.Int,error){
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(16), nil).Sub(max, big.NewInt(1))
	k, err := rand.Int(rand.Reader, max)
	if err != nil {}
	var xT = big.NewInt(0)
	var yT = big.NewInt(0)
	yT.Exp(g,b, p)
	N := GetN(k,p)
	var i = big.NewInt(1)
	var temp = big.NewInt(0)
	for i.Cmp(N)<1{
		temp = f(yT,k,p)
		xT.Add(xT,temp)
		temp.Exp(g,temp,p)
		yT.Mul(yT, temp)
		i.Add(i,big.NewInt(1))
	}
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
		temp1.Exp(g,temp1,p)
		yW.Mul(yW,temp1)
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
	Sum.Mul(Sum, big.NewInt(4))
	return Sum
}

func f(y,k,p *big.Int) *big.Int{
	var temp = big.NewInt(0)
	temp.Mod(y, k)
	temp.Exp(big.NewInt(2),temp,p)
	return temp
}
