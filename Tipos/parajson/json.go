package main

import (
	"encoding/json"
	"fmt"
)

type produto struct{
	ID int `json:"ID"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tag []string `json::"tags"`
}

func main(){

	p1 := produto{1,"maicon", 200.00, []string{"douglas","nao tem jeito"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))


}