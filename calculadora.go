package main

import (
	"fmt"
)

func main() {
	var operacao string
	var num1, num2 float64

	fmt.Print("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&operacao)

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	resultado := 0.0

	switch operacao {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 != 0 {
			resultado = num1 / num2
		} else {
			fmt.Println("Erro: divisão por zero!")
			return
		}
	default:
		fmt.Println("Operação inválida!")
		return
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
