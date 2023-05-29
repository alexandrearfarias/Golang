package main

// Constantes, structs e funções que iniciam em Maiúsculo são componentes públicos(visivel fora do pacote)
import "math"

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) ( catX, catY float64){
	catX = math.Abs(b.x - a.x) 
	catY = math.Abs(b.x - a.x)
	return
}

func Distancia(a,b Ponto) float64{
	cx, cy := catetos(a, b)

	return math.Sqrt(math.Pow(cx,2) + math.Pow(cy,2))
}