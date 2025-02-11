package main

import (
	"fmt"
	"log"

	"go-conversor/conversor"
	"go-conversor/gui"
)

func main() {
	// Exibe o menu principal
	fmt.Println("=== Conversor de Temperatura ===")
	fmt.Println("1. Modo Terminal")
	fmt.Println("2. Modo Interface Gráfica")
	fmt.Print("Escolha uma opção (1 ou 2): ")

	var opcao int
	if _, err := fmt.Scan(&opcao); err != nil {
		log.Fatalf("Erro ao ler a opção: %v", err)
	}

	switch opcao {
	case 1:
		runTerminalMode()
	case 2:
		// Chama a função que inicia a interface gráfica
		gui.RunGUI()
	default:
		fmt.Println("Opção inválida!")
	}
}

func runTerminalMode() {
	// Menu para o modo terminal
	fmt.Println("=== Modo Terminal ===")
	fmt.Println("1. Celsius para Fahrenheit")
	fmt.Println("2. Fahrenheit para Celsius")
	fmt.Print("Escolha uma opção (1 ou 2): ")

	var opcao int
	if _, err := fmt.Scan(&opcao); err != nil {
		log.Fatalf("Erro ao ler a opção: %v", err)
	}

	switch opcao {
	case 1:
		fmt.Print("Digite a temperatura em Celsius: ")
		var celsius float64
		if _, err := fmt.Scan(&celsius); err != nil {
			log.Fatalf("Erro ao ler a temperatura: %v", err)
		}
		fahrenheit := conversor.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
	case 2:
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		var fahrenheit float64
		if _, err := fmt.Scan(&fahrenheit); err != nil {
			log.Fatalf("Erro ao ler a temperatura: %v", err)
		}
		celsius := conversor.FahrenheitToCelsius(fahrenheit)
		fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
	default:
		fmt.Println("Opção inválida!")
	}
}
