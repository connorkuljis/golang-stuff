package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func do() {
	start := time.Now()
	echo1()
	echo2()
	echo3()
	echo4()
	fmt.Printf("\n\t\ttotal %fs elapsed\n", time.Since(start).Seconds())
}

func echo1() {
	fmt.Print("echo1:\n")
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		indexStr := strconv.Itoa(i)
		s += sep + indexStr + " : " + os.Args[i]
		sep = "\n"
		// fmt.Printf("%d : %s\n", i, os.Args[i])
	}
	fmt.Println(s)
	fmt.Printf("\t%fs elapsed\n", time.Since(start).Seconds())
}

func echo2() {
	fmt.Print("echo2: ")
	start := time.Now()
	s, sep := "", ""
	// _ is an unused variable for the current index count
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("\t%fs elapsed\n", time.Since(start).Seconds())
}

func echo3() {
	fmt.Print("echo3: ")
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("\t%fs elapsed\n", time.Since(start).Seconds())
}

func echo4() {
	fmt.Print("echo4: ")
	start := time.Now()
	fmt.Println(os.Args[1:])
	fmt.Printf("\t%fs elapsed\n", time.Since(start).Seconds())
}
