package main

import (
	"fmt"
	"math"
	"roo/coin/pkg/finitefield"
)

func main() {
	// elements := finitefield.CalculateFromMultiplication([]int64{1, 3, 7, 13, 18}, 3)
	// for _, elem := range elements {
	// 	fmt.Printf("%+v\n", elem)
	// }

	result := int64(math.Pow(7, 17))
	result = result % 19

	fmt.Printf("result is : %d", result)

	elements := finitefield.CalculateFromExponent([]int64{2}, 3)
	for _, elem := range elements {
		fmt.Printf("%+v\n", elem)
	}

}
