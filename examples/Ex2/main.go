package main

import "fmt"

type Params struct {
    altura float64
    base   float64
	largura float64
}

type MqParams func(*Params) error

func calculaMetroCubico(r *Params) float64 {
	return r.altura * r.base  * r.largura
}

func WithParam(altura float64, base float64, largura float64) MqParams {
    return func(r *Params) error {
		r.altura = altura
		r.base = base
		r.largura = largura
        return nil
    }

}

func main() {
	params := &Params{}

	err := WithParam(10, 10, 10)(params)
	if err != nil {
		fmt.Println("Erro ao atribuir parâmetros:", err)
		return
	}

	area := calculaMetroCubico(params)
	fmt.Println("Área do metrocubico:", area)
}