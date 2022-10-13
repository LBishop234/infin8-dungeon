package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	anApp := app.New()
	aWindow := anApp.NewWindow("infin8-dungeon")
	aWindow.Resize(fyne.NewSize(1000.0, 1000.0))
	aWindow.SetFixedSize(true)

	displayPanel := newDisplayPanel()

	controlPanel := newControlPanel(anApp, displayPanel)

	aWindow.SetContent(container.NewHBox(
		controlPanel,
		displayPanel,
	))

	aWindow.ShowAndRun()
}
