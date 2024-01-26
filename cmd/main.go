package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func result() int {
	var myInt int
	return myInt
}

func main() {
	fibonacci(10)
	fmt.Println(result())
}
