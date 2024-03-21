package main

import (
	"fmt"
	"options-pattern/calcmc"
	"os"
)

func main() {

	mc, err := calcmc.NewCalcMQ(calcmc.WithAltura(10), calcmc.WithBase(10))
	if err != nil {
		fmt.Println("Erro ao criar objeto:", err)
		os.Exit(1)
	}

	fmt.Println("Área do retângulo:", mc.CalculaMQ())

	// continuar os testes aqui para baixo.

}
