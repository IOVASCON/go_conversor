package conversor

// CelsiusToFahrenheit converte Celsius para Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
    return (celsius * 9 / 5) + 32
}

// FahrenheitToCelsius converte Fahrenheit para Celsius
func FahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5 / 9
}