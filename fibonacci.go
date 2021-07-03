package main

// Fibonacci Memoization

import (
	"fmt"
	"time"
)

func fibMemo(n int, mem map[int]int) int {
	if mem[n]&1 != 0 {
		return mem[n]
	}
	if n <= 2 {
		return 1
	}
	mem[n] = fibMemo(n-1, mem) + fibMemo(n-2, mem)
	return mem[n]
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	result := fib(n-1) + fib(n-2)
	return result
}

func comparisonFunction() {
	fmt.Println("Fibonacci comparison between ordinary and memoization")
	startA := time.Now()
	a := fibMemo(50, map[int]int{0: 0})
	fmt.Println(a)
	endA := time.Since(startA) * 1000

	startB := time.Now()
	b := fib(50)
	fmt.Println(b)
	endB := time.Since(startB) * 1000
	fmt.Println(endA < endB)

}

func main() {
	fmt.Println("Febonacci Momoization")
	a := fibMemo(50, map[int]int{0: 0})
	fmt.Println(a)
}
