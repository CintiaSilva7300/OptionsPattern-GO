package main

import (
	"fmt"
	"options-pattern/calcmc"
)

type Params struct {
	altura float64
	base   float64
}

type MqParams func(*Params) error

func calculaMetroQuadrado(r *Params) float64 {
	return r.altura * r.base
}

func WithParam(altura float64, base float64) MqParams {
	return func(r *Params) error {
		r.altura = altura
		r.base = base
		return nil
	}

}

func main() {
	params := &Params{}

	err := WithParam(10, 10)(params)
	if err != nil {
		fmt.Println("Erro ao atribuir parâmetros:", err)
		return
	}

	area := calculaMetroQuadrado(params)
	fmt.Println("Área do retângulo:", area)

	fmt.Printf("---- \n")

	cc, err := calcmc.NewCalcMQ(calcmc.WithAltura(10), calcmc.WithBase(10))
	if err != nil {
		fmt.Println("Erro ao criar objeto:", err)
		return
	}

	fmt.Println("Área do retângulo:", cc.CalculaMQ())

	fmt.Printf("---- \n")

	var qualquercoisa interface{} // qualquer coia mesmo

	qualquercoisa = 1.0

	val, ok := qualquercoisa.(int)
	if ok {
		fmt.Println("É um inteiro:", val)
	} else {
		fmt.Println("Não é um inteiro")
	}

	switch qualquercoisa.(type) {
	case int, int32, int64:
		fmt.Println("É um inteiro")
	case string:
		fmt.Println("É uma string")
	default:
		fmt.Println("Não sei o que é")
	}

}
