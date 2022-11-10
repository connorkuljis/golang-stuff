package main

import "fmt"

// triangular numbers using recursion

func main() {
	fmt.Println(triangularRecursive(6))
}

func triangularRecursive(n int) int {
	if n == 1 {
		return n
	}
	rv := n + triangularRecursive(n-1)
	fmt.Printf("%d | %d\n", n, rv)
	return rv
}
