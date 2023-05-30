package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

const (
	PROGNAME              = "Alkmini Management System"
	STATUS_UP             = "Server Status: Host up and running"
	STATUS_DOWN           = "Server Status: Host down"
	STATUS_TUNNELED       = "Server Status: Tunnel running"
	PROCESS_START         = "Process: Starting Server"
	PROCESS_STARTED       = "Process: Server started"
	PROCESS_START_ERROR   = "Process: Server start error"
	PROCESS_SUSPENDED     = "Process: Server suspended"
	PROCESS_SUSPEND_ERROR = "Process: Server suspend error"
	PROCESS_STOP          = "Process: Stopping Server"
	PROCESS_TUNNEL        = "Process: Tunneling"
	PROCESS_TUNNELED      = "Process: Tunnel running"
	PROCESS_TUNNEL_ERROR  = "Process: Tunnel error"
	HIDDEN                = ""
	GETSERVERSTATUSLABEL  = "Get Server Status"
	STARTSERVERLABEL      = "Start Server"
	STOPSERVERLABEL       = "Stop Server"
	TUNNELLABEL           = "Tunnel sshuttle"
)

func main() {
	a := app.New()
	w := a.NewWindow(PROGNAME)
	label := widget.NewLabel(PROGNAME)
	status := widget.NewLabel(HIDDEN)
	process := widget.NewLabel(HIDDEN)
	w.SetContent(container.NewVBox(
		label,
		status,
		widget.NewSeparator(),
		process,

		widget.NewButtonWithIcon(GETSERVERSTATUSLABEL, theme.InfoIcon(), func() {
			status.SetText(getStatus())
		}),
		widget.NewButtonWithIcon(STARTSERVERLABEL, theme.MediaPlayIcon(), func() {
			process.SetText(PROCESS_START)
			if startServer() != nil {
				process.SetText(PROCESS_START_ERROR)
				status.SetText(STATUS_DOWN)
				return
			} else {
				waitUntilUp()
				process.SetText(PROCESS_STARTED)
				status.SetText(STATUS_UP)
			}
			time.Sleep(5 * time.Second)
			process.SetText("")

		}),
		widget.NewButtonWithIcon(STOPSERVERLABEL, theme.MediaStopIcon(), func() {
			process.SetText(PROCESS_STOP)
			if stopServer() != nil {
				process.SetText(PROCESS_SUSPEND_ERROR)
				status.SetText(STATUS_UP)
			} else {
				waitUntilDown()
				process.SetText(PROCESS_SUSPENDED)
				status.SetText(STATUS_DOWN)
			}
			time.Sleep(5 * time.Second)
			process.SetText(HIDDEN)
		}),
		widget.NewButtonWithIcon(TUNNELLABEL, theme.LoginIcon(), func() {
			process.SetText(PROCESS_TUNNEL)
			if createTunnel() != nil {
				process.SetText(PROCESS_TUNNEL_ERROR)
				status.SetText(getStatus())
			} else {
				process.SetText(PROCESS_TUNNELED)
				status.SetText(STATUS_TUNNELED)
			}
			time.Sleep(5 * time.Second)
			process.SetText(HIDDEN)
		}),
	))
	w.ShowAndRun()
}
