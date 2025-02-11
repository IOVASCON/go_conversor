package conversor

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {0, 32},
        {100, 212},
        {-40, -40},
    }

    for _, test := range tests {
        result := CelsiusToFahrenheit(test.input)
        if result != test.expected {
            t.Errorf("CelsiusToFahrenheit(%v) = %v; expected %v", test.input, result, test.expected)
        }
    }
}

func TestFahrenheitToCelsius(t *testing.T) {
    tests := []struct {
        input    float64
        expected float64
    }{
        {32, 0},
        {212, 100},
        {-40, -40},
    }

    for _, test := range tests {
        result := FahrenheitToCelsius(test.input)
        if result != test.expected {
            t.Errorf("FahrenheitToCelsius(%v) = %v; expected %v", test.input, result, test.expected)
        }
    }
}