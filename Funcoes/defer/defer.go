package main

import "fmt"

func obterV(num int ){
	defer fmt.Println("-----")

	if num < 5000{
		fmt.Println("Aprovado")
	}else {
		fmt.Println("Reprovado")
	}
}

func main(){
	obterV(1000)
	obterV(7000)
}