package main

import "fmt"

func fatorial(n int) int{
	switch{
	case n==0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	fmt.Println(fatorial(5))
	fmt.Println(fatorial(0))
}