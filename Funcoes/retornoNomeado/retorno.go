package main

import "fmt"

func trocado(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	x, y := trocado(1, 2)

	fmt.Println(x, y)
}
