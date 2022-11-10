package main

import (
	"fmt"
	"strconv"
)

const InfoColor = "\033[1;34m%s\033[0m"

func main() {
	x := 5
	fmt.Printf("triangular number where n=%d is %d\n", x, triangularRecursive(x))
	fmt.Printf("%d * (%d + 1) / 2 = %d\n", x, x, triangularGeneral(x))
}

// triangular numbers using recursion
func triangularRecursive(n int) int {
	if n == 1 {
		fmt.Println("hit base case")
		fmt.Printf("n | val\t[ call ]\n")
		return n
	}
	fmt.Printf("n = %d\n", n)
	rv := n + triangularRecursive(n-1)

	fmt.Printf("%d | %d \t[ ", n, rv)
	call := strconv.Itoa(n) + " + tri(" + strconv.Itoa(n-1) + ") ]\n"
	fmt.Printf(InfoColor, call)

	return rv
}

func triangularGeneral(n int) int {
	return n * (n + 1) / 2
}
