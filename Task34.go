package main

import (
	"fmt"
	"math/big"
)

func Task34() {
	var Alice = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(345)}
	var Bob = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(234)}
	var MalloryToAlice = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(22)}
	var MalloryToBob = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(22)}
	Alice.pkGen()
	Bob.pkGen()

	MalloryToAlice.SharedSecretGen(Alice.pk)
	Bob.SharedSecretGen(Alice.p)


	MalloryToBob.SharedSecretGen(Bob.pk)
	Alice.SharedSecretGen(Bob.p)

	fmt.Println("Task34:")
	fmt.Println(Alice.ss)
	fmt.Println(Bob.ss)
	fmt.Println(MalloryToAlice.ss)
	fmt.Println(MalloryToBob.ss)
}