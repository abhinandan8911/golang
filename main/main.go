package main

import (
	"github.com/abhinandan8911/golang/arrays"
	"fmt"
)

func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(arrays.FindMaxConsecutiveOnes(input))
}
