package main

import "fmt"

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
}