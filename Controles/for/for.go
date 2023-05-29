package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i < 5 { //como seria um laço while em GO
		fmt.Println(i)
		i++
	}

	for i := 0; i < 12; i += 2 {
		fmt.Printf("%d", i)
	}

	for {	//laço eterno
		fmt.Println("eterno..")
		time.Sleep(time.Second)
	}
}
