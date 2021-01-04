package main

import (
	"fmt"
	"math/big"
)

func Task33() {
	var Alice = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(345)}
	var Bob = DH {p : big.NewInt(2017), g:big.NewInt(127), sk : big.NewInt(234)}
	Alice.pkGen()
	Bob.pkGen()
	Alice.SharedSecretGen(Bob.pk)
	Bob.SharedSecretGen(Alice.pk)
	fmt.Println("Task33:")
	fmt.Println(Alice.ss)
	fmt.Println(Bob.ss)
}


