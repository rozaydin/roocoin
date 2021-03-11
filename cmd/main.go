package main

import (
	"fmt"
	"roo/coin/pkg/finitefield"
)

func main() {
	// elements := finitefield.CalculateFromMultiplication([]int64{1, 3, 7, 13, 18}, 19)
	// for _, elem := range elements {
	// 	fmt.Printf("%+v\n", elem)
	// }

	elements := finitefield.CalculateFromExponent([]int64{4}, 5)
	for _, elem := range elements {
		fmt.Printf("%+v\n", elem)
	}

}
