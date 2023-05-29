package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

/*
Usamos o ponteiro no receptor da função pois sem ele o valor dos atributos da struct nao seriam alterados
Sem o ponteiro, seria alterado apenas o valor da variavel da função( a varaivel receptora "p") que é interna a ela e não interage com valores externos.
Por isso fazemos uma referencia ao tipo pessoa, pra ele mudar diretamente no endereço de memoria.

EXEMPLO:
Criamos a pessoa p1, caso o receptor (p pessoa) nao tivesse ponteiro, a funcao alteraria apenas o 'p' e nao o 'p1' pois o valor atribuido esta preso ao escopo da funcao e nao sendo referenciado a uma variavel externa a ela.
*/
func (p *pessoa) setNomeCompleto(nome, sobrenome string) {
	p.nome = nome
	p.sobrenome = sobrenome
}


func main(){
	p1 := pessoa{"Pedro", "Sampaio"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Joao", "Cleber de Almeida")
	fmt.Println(p1.getNomeCompleto())
}