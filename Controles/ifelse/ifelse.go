package main

import "fmt"

func obterv(nota float64) {
	if nota >= 6 { //no if do GO não precisa de parentese na condição
		fmt.Println("Aprovado com nota :", nota)
	}else{
		fmt.Println("Reprovado com nota:", nota)
	}
}
func main(){
	obterv(6.5)
	obterv(5)

}