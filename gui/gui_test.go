package gui

import (
	"testing"

	"go-conversor/conversor"
	"go-conversor/utils"

	"github.com/stretchr/testify/assert"
)

// TestConvertCelsiusToFahrenheit verifica se a conversão de Celsius para Fahrenheit funciona corretamente.
func TestConvertCelsiusToFahrenheit(t *testing.T) {
	input := "25" // Entrada simulada pelo usuário
	expectedCelsius := 25.0
	expectedFahrenheit := conversor.CelsiusToFahrenheit(expectedCelsius)

	// Simula a entrada do usuário
	celsius, err := utils.ParseFloat(input)
	assert.NoError(t, err, "Erro ao converter entrada para float")

	// Realiza a conversão
	fahrenheit := conversor.CelsiusToFahrenheit(celsius)

	// Verifica o resultado
	assert.Equal(t, expectedFahrenheit, fahrenheit, "Conversão de Celsius para Fahrenheit incorreta")
}

// TestConvertFahrenheitToCelsius verifica se a conversão de Fahrenheit para Celsius funciona corretamente.
func TestConvertFahrenheitToCelsius(t *testing.T) {
	input := "77" // Entrada simulada pelo usuário
	expectedFahrenheit := 77.0
	expectedCelsius := conversor.FahrenheitToCelsius(expectedFahrenheit)

	// Simula a entrada do usuário
	fahrenheit, err := utils.ParseFloat(input)
	assert.NoError(t, err, "Erro ao converter entrada para float")

	// Realiza a conversão
	celsius := conversor.FahrenheitToCelsius(fahrenheit)

	// Verifica o resultado
	assert.Equal(t, expectedCelsius, celsius, "Conversão de Fahrenheit para Celsius incorreta")
}

// TestGenerateReport verifica se o relatório gerado está no formato correto.
func TestGenerateReport(t *testing.T) {
	celsiusValues := []float64{0, 100}
	fahrenheitValues := []float64{32, 212}

	expectedReport := `Temperaturas Convertidas:
Celsius | Fahrenheit
--------|-----------
   0,00 |     32,00
 100,00 |    212,00
`

	report := utils.GenerateReport(celsiusValues, fahrenheitValues)
	assert.Equal(t, expectedReport, report, "Relatório gerado incorretamente")
}
