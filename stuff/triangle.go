package main

import (
	"fmt"
	"strconv"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	tri(5)
}

func tri(n int) int {
	return triangularRecursive(n)
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
