package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Alkmini Management system")
	label := widget.NewLabel("Alkmini Management System")
	status := widget.NewLabel("")
	process := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		label,
		status,
		widget.NewSeparator(),
		process,

		widget.NewButtonWithIcon("Get Server Status", theme.InfoIcon(), func() {
			status.SetText(getStatus())
		}),
		widget.NewButtonWithIcon("Start Server", theme.MediaPlayIcon(), func() {
			process.SetText("Starting Server")
			startServer()
			for getStatus() != "Host up and running" {
				time.Sleep(5 * time.Second)
			}
			process.SetText("Server started")
			status.SetText("Host up and running")
			time.Sleep(5 * time.Second)
			process.SetText("")

		}),
		widget.NewButtonWithIcon("Stop Server", theme.MediaStopIcon(), func() {
			process.SetText("stopping")
			status.SetText("inactivo")
		}),
		widget.NewButtonWithIcon("Tunnel sshuttle", theme.LoginIcon(), func() {
			process.SetText("tunnel 192.168.122.0/24")
			status.SetText("Tunnel activo")
		}),
	))
	w.ShowAndRun()
}
