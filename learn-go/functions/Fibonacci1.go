package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		fmt.Print(n, " ")
		return n
	}

	a := 0
	b := 1
	fmt.Print(a, " ", b, " ")

	var temp int
	for i := 2; i < n; i++ {
		temp = a + b
		fmt.Print(temp, " ")
		a = b
		b = temp
	}
	fmt.Println()
	return b
}

func main() {
	result := fib(10)
	fmt.Println("10th Fibonacci number is:", result)
}
