package utils

import (
	"fmt"
	"strconv" // Converte strings para outros tipos, como float64

	"golang.org/x/text/language" // Fornece suporte para localização (ex.: idioma e formato regional)
	"golang.org/x/text/message"  // Fornece formatação de texto com base no idioma (ex.: números no formato brasileiro)
)

var printer = message.NewPrinter(language.BrazilianPortuguese)

// FormatDecimal formata um número no sistema decimal brasileiro
func FormatDecimal(value float64) string {
	return printer.Sprintf("%.2f", value)
}

// ParseFloat converte uma string para float64
func ParseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// GenerateReport gera um relatório em formato tabular
func GenerateReport(celsiusValues, fahrenheitValues []float64) string {
	report := "Temperaturas Convertidas:\n"
	report += "Celsius | Fahrenheit\n"
	report += "--------|-----------\n"

	for i := range celsiusValues {
		celsius := FormatDecimal(celsiusValues[i])
		fahrenheit := FormatDecimal(fahrenheitValues[i])
		report += fmt.Sprintf("%7s | %10s\n", celsius, fahrenheit)
	}

	return report
}
