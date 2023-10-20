package mySamplePackage

func Fib(num uint) uint {
	if num <= 1 {
		return num
	}
	return Fib(num-1) + Fib(num-2)
}
