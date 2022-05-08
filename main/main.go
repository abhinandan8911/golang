package main

import "fmt"

type MyInt int

type MyFunc func(message string)

func Log(message string) {
	fmt.Println("Hello ", message)
}

func main() {
	var i interface{}
	var mine MyInt
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	myfunc := MyFunc(Log)
	myfunc("Abhi")
}
