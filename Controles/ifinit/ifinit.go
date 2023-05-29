package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := random(); i < 5 { //mesma estrutura que usa no 'for' for(int i=0; i<5; i++)
		fmt.Println("Fixa")
	} else {
		fmt.Println("não Fixa")
	}
}
