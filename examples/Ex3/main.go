package main

import "fmt"

// Calcular percentagem de um número, multiplicado por outro número qualquer. (num * 1.0) * num

type Params struct {
    num1 float64
    num2 float64
	multiplicador float64
}

type MqParams func(*Params) error

func calculaPorcentagem(r *Params) float64 {
	return (r.num1 * r.num2) * 1.0 * r.multiplicador
}

func WithParam(num1 float64, multiplicador float64, num2 float64) MqParams {
    return func(r *Params) error {
		r.num1 = num1
		r.num2 = num2
		r.multiplicador = multiplicador
        return nil
    }

}

func main() {
	params := &Params{}

	err := WithParam(100, 4, 0.12)(params)
	if err != nil {
		fmt.Println("Erro ao atribuir parâmetros:", err)
		return
	}

	res := calculaPorcentagem(params)
	fmt.Println("Calculo da porcentagem:", res)
}