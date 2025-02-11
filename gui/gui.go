package gui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"go-conversor/conversor"
	"go-conversor/utils"
)

func RunGUI() {
	os.Setenv("FYNE_BACKEND", "software")

	myApp := app.New()
	myWindow := myApp.NewWindow("Conversor de Temperatura")

	input := widget.NewEntry()
	output := widget.NewLabel("")
	resultTable := widget.NewLabel("")

	convertCtoF := func() {
		celsius, err := utils.ParseFloat(input.Text)
		if err != nil {
			output.SetText("Entrada inválida!")
			return
		}
		fahrenheit := conversor.CelsiusToFahrenheit(celsius)
		formattedCelsius := utils.FormatDecimal(celsius)
		formattedFahrenheit := utils.FormatDecimal(fahrenheit)
		output.SetText(fmt.Sprintf("%s°C = %s°F", formattedCelsius, formattedFahrenheit))
		resultTable.SetText(utils.GenerateReport([]float64{celsius}, []float64{fahrenheit}))
	}

	convertFtoC := func() {
		fahrenheit, err := utils.ParseFloat(input.Text)
		if err != nil {
			output.SetText("Entrada inválida!")
			return
		}
		celsius := conversor.FahrenheitToCelsius(fahrenheit)
		formattedCelsius := utils.FormatDecimal(celsius)
		formattedFahrenheit := utils.FormatDecimal(fahrenheit)
		output.SetText(fmt.Sprintf("%s°F = %s°C", formattedFahrenheit, formattedCelsius))
		resultTable.SetText(utils.GenerateReport([]float64{celsius}, []float64{fahrenheit}))
	}

	btnCtoF := widget.NewButton("Celsius -> Fahrenheit", convertCtoF)
	btnFtoC := widget.NewButton("Fahrenheit -> Celsius", convertFtoC)

	content := container.NewVBox(
		widget.NewLabel("Digite a temperatura:"),
		input,
		btnCtoF,
		btnFtoC,
		output,
		resultTable,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
