package main

import "fmt"

func main() {
	//a1 := [3]int{1, 2, 3} //um array
	//s1 := []int{1, 2, 3}  //um slice

	a2 := [5]int{1, 2, 3, 4,5}

	//slice é um pedaço de um array e não um array
	s2 := a2[1:3] // s2 é um pedaço do a2 que vai do indice 1 ao 3 sem pegar o valor do indice 3
	fmt.Println(a2,"\n",s2)

	s3 := a2[:2] // vai do inicio do array
	fmt.Println(s3, s2)
}