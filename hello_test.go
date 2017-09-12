package main

import "fmt"
import "testing"

func fib(n int) int {
	if n < 4 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	for i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}
	fmt.Println("hello from spacemacs")
}
