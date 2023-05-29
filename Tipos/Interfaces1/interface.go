package main

import "fmt"

type esportiva interface {
	ligarTurbo()
}

type ferrari struct {
	modelo     string
	turbo      bool
	velocidade int
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	carro1 := ferrari{"lamborguine", false, 0}
	carro1.ligarTurbo()

	// nesse tipo de declaraçao, é preciso usar o "&" antes do ferrari pois estamos trabalhando com a interface, e nao diretamente com o tipo, como no carro1. Caso o metodo ligarTurbo mude algum valor direto do tipo ferrari, é necessario passar o endereço ao inves apenas da valor direto.
	var carro2 esportiva = &ferrari{"bmw", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}