package main

import (
	"arrays"
	"fmt"
)

func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(arrays.FindMaxConsecutiveOnes(input))
}
