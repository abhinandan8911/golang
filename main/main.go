package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
}
