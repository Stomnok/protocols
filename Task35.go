package main

import (
	"fmt"
	"math/big"
)

func Task35Gis1() {
	var Alice = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(345)}
	var Bob = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(234)}
	var MalloryToAlice = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	var MalloryToBob = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	Alice.pkGen()
	Bob.pkGen()

	MalloryToAlice.SharedSecretGen(Alice.pk)
	Bob.SharedSecretGen(Alice.pk)

	MalloryToBob.SharedSecretGen(Bob.pk)
	Alice.SharedSecretGen(Bob.pk)

	fmt.Println("Task35Gis1:")
	fmt.Println(Alice.ss)
	fmt.Println(Bob.ss)
	fmt.Println(MalloryToAlice.ss)
	fmt.Println(MalloryToBob.ss)
}

func Task35GisP() {
	var Alice = DH {p : big.NewInt(2017), g:big.NewInt(2017), sk : big.NewInt(345)}
	var Bob = DH {p : big.NewInt(2017), g:big.NewInt(2017), sk : big.NewInt(234)}
	var MalloryToAlice = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	var MalloryToBob = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	Alice.pkGen()
	Bob.pkGen()

	MalloryToAlice.SharedSecretGen(Alice.pk)
	Bob.SharedSecretGen(Alice.pk)

	MalloryToBob.SharedSecretGen(Bob.pk)
	Alice.SharedSecretGen(Bob.pk)

	fmt.Println("Task35GisP:")
	fmt.Println(Alice.ss)
	fmt.Println(Bob.ss)
	fmt.Println(MalloryToAlice.ss)
	fmt.Println(MalloryToBob.ss)
}

func Task35GisPSub1() {
	var Alice = DH {p : big.NewInt(2017), g:big.NewInt(2016), sk : big.NewInt(121)}
	var Bob = DH {p : big.NewInt(2017), g:big.NewInt(2016), sk : big.NewInt(371)}
	var MalloryToAlice = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	var MalloryToBob = DH {p : big.NewInt(2017), g:big.NewInt(1), sk : big.NewInt(22)}
	Alice.pkGen()
	Bob.pkGen()

	MalloryToAlice.SharedSecretGen(Alice.pk)
	Bob.SharedSecretGen(Alice.pk)

	MalloryToBob.SharedSecretGen(Bob.pk)
	Alice.SharedSecretGen(Bob.pk)

	fmt.Println("Task35GisPSub1:")
	fmt.Println(Alice.ss)
	fmt.Println(Bob.ss)
	fmt.Println(MalloryToAlice.ss)
	fmt.Println(MalloryToBob.ss)
}