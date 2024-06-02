package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Application")
	w.Resize(fyne.NewSize(300, 200))
	label1 := widget.NewLabel("Введите первое число")
	entry1 := widget.NewEntry()
	label2 := widget.NewLabel("Введите второе число")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")
	//responser:= widget.
	btn := widget.NewButton("touch me", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry1.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Error")
		}
		sum := n1 + n2
		div := n1 + n2
		mul := n1 + n2
		sub := n1 + n2

		answer.SetText(fmt.Sprintf("(+) %f\n  (-) %f\n (/) %f\n (*) %f\n", sum, sub, div, mul))
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
	))

	w.ShowAndRun()
}
