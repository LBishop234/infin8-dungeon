package main

import (
	"image/color"
	"main/room"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newControlPanel(anApp fyne.App, displayPanel *fyne.Container) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Control"),
		widget.NewButton("Generate New Room", func() {
			displayPanel.RemoveAll()

			aRoom := room.GenerateRoom()
			displayRect := canvas.NewRectangle(color.Black)
			displayRect.SetMinSize(fyne.NewSize(float32(aRoom.Width*50), float32(aRoom.Height*50)))
			displayPanel.Add(displayRect)
		}),
		widget.NewButton("Quit", func() {
			anApp.Quit()
		}),
	)
}
