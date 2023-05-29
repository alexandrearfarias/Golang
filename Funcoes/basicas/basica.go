package main

import "fmt"

func f1() {
	fmt.Println("olá")

}

func f2(p1 string, p2 string) {
	fmt.Printf("%s e %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("%s ou %s", p1, p2)
}

func f5() (string, string) {
	return "retorno 1", "retorno 2"
}

func main() {
	f1()
	f2("awa", "uwu")
	x, y := f3(), f4("awa", "uwu")
	fmt.Println(x)
	fmt.Println(y)
	f5()
}
