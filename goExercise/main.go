package main

import (
	"fmt"

	"github.com/leon123858/goSamplePackage"
)

type node struct {
	value int
	next  *node
}

type stack struct {
	header *node
	ender  *node
}

func push(st *stack, value int) {

}

func reverse(source *string) {

}

func main() {
	println(goSamplePackage.Fibonacci(15)) 
	fmt.Println("Start!")
	var height string
	fmt.Scanln(&height)
}
