package main

import (
	"fmt"

	"github.com/hashlab/dukptkit/lib"
)

func main() {
	keyBytes1, key1, _ := lib.GenerateComponentKey(16, true)
	keyBytes2, key2, _ := lib.GenerateComponentKey(16, true)
	keyBytes3, key3, _ := lib.GenerateComponentKey(16, true)

	kcvBytes1, kcv1, _ := lib.CalculateKcv(keyBytes1)
	kcvBytes2, kcv2, _ := lib.CalculateKcv(keyBytes2)
	kcvBytes3, kcv3, _ := lib.CalculateKcv(keyBytes3)

	cKeyBytes, cKey := lib.GenerateCombinedKey(keyBytes1, keyBytes2, keyBytes3)
	cKcvBytes, cKcv, _ := lib.CalculateKcv(cKeyBytes)

	fmt.Println("Component key 1: ")
	fmt.Println(keyBytes1)
	fmt.Println(key1)
	fmt.Println(kcvBytes1)
	fmt.Println(kcv1)

	fmt.Println("Component key 2: ")
	fmt.Println(keyBytes2)
	fmt.Println(key2)
	fmt.Println(kcvBytes2)
	fmt.Println(kcv2)

	fmt.Println("Component key 3: ")
	fmt.Println(keyBytes3)
	fmt.Println(key3)
	fmt.Println(kcvBytes3)
	fmt.Println(kcv3)

	fmt.Println("Combined key: ")
	fmt.Println(cKeyBytes)
	fmt.Println(cKey)
	fmt.Println(cKcvBytes)
	fmt.Println(cKcv)

	ihex1, _ := lib.HexToBytes("1C4F20AB1CF2169EA4BCDF29ABD351AB")
	ihex2, _ := lib.HexToBytes("29C73E13EF515E37A7649DF220A2B9D0")
	ihex3, _ := lib.HexToBytes("9ECD54310B7A26E9328FDA893158E585")

	icKeyBytes, icKey := lib.GenerateCombinedKey(ihex1, ihex2, ihex3)
	icKcvBytes, icKcv, _ := lib.CalculateKcv(icKeyBytes)

	fmt.Println("Inputed Combined key: ")
	fmt.Println(icKeyBytes)
	fmt.Println(icKey)
	fmt.Println(icKcvBytes)
	fmt.Println(icKcv)

	isOdd := lib.IsOddParityAdjusted(cKeyBytes)

	fmt.Println("Is Odd Parity adjusted: ")
	fmt.Println(isOdd)

}
