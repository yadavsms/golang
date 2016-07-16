// Fibonacci Series Using for recursion
package main

import (
	"fmt"
)

// Using recursion
func fib(n int) int {
	if n <=3 { 
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
        fmt.Println(0)
	for i:=2; i<=10; i++ {
		fmt.Println(fib(i))
	}
}
