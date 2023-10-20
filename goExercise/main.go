package main

import (
	"fmt"
	mySamplePackage "github.com/leon123858/go_exercise/mypackage"
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
	mySamplePackage.Fib(0)
	fmt.Println("Start!")
	var height string
	fmt.Scanln(&height)
}
