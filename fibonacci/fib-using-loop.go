// Fibonacci Series Using for loop 
package main

import (
	"fmt"
)

func fib(n int) int {
	x , y := 0, 1
	for i:= 0; i < n; i++ {
		x , y = y , x + y
	}
	return x
}

func main() {
	for i:=0; i<=10; i++ {
		fmt.Println(fib(i))
	}
}
