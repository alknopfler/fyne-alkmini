package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
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
	PROCESS_TUNNEL_DEL    = "Process: Tunnel removing"
	PROCESS_TUNNELED      = "Process: Tunnel running"
	PROCESS_TUNNELED_DEL  = "Process: Tunnel removed"
	PROCESS_TUNNEL_ERROR  = "Process: Tunnel error"
	HIDDEN                = ""
	GETSERVERSTATUSLABEL  = "Get Server Status"
	STARTSERVERLABEL      = "Start Server"
	STOPSERVERLABEL       = "Stop Server"
	TUNNELLABEL           = "Create Tunnel"
	TUNNELREMOVELABEL     = "Remove Tunnel"
	STARTING_SERVER       = "Starting Server..."
)

func main() {
	a := app.New()
	w := a.NewWindow(PROGNAME)
	var m *fyne.Menu
	if desk, ok := a.(desktop.App); ok {
		m = fyne.NewMenu("alkmini",
			fyne.NewMenuItem("Open app", func() {
				w.Show()
			}),
			fyne.NewMenuItem(getStatus(), func() {
				m.Items[1].Label = getStatus()
				m.Refresh()
			}),
			fyne.NewMenuItem("Start Server", func() {
				startServer()
				m.Items[1].Label = getStatus()
				m.Refresh()
			}),
			fyne.NewMenuItem("Stop Server", func() {
				stopServer()
				m.Items[1].Label = getStatus()
				m.Refresh()
			}),
			fyne.NewMenuItem("Create Tunnel", func() {
				createTunnel()
				m.Items[1].Label = getStatus()
				m.Refresh()
			}),
			fyne.NewMenuItem("Remove Tunnel", func() {
				removeTunnel()
				m.Items[1].Label = getStatus()
				m.Refresh()
			}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(theme.ComputerIcon())
		m.Items[1].Label = getStatus()
		m.Refresh()
	}
	w.SetContent(widget.NewLabel("Alkmini Management System"))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.SetIcon(theme.ComputerIcon())
	w.Resize(fyne.NewSize(400, 400))
	status := widget.NewLabel(HIDDEN)
	process := widget.NewLabel(HIDDEN)
	w.SetContent(container.NewVBox(
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
			} else {
				status.SetText(STARTING_SERVER)
				waitUntilUp()
				process.SetText(PROCESS_STARTED)
				status.SetText(STATUS_UP)

			}
			m.Items[1].Label = getStatus()
			m.Refresh()
			time.Sleep(5 * time.Second)
			process.SetText(HIDDEN)
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
			m.Items[1].Label = getStatus()
			m.Refresh()
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
			m.Items[1].Label = getStatus()
			m.Refresh()
			time.Sleep(5 * time.Second)
			process.SetText(HIDDEN)
		}),
		widget.NewButtonWithIcon(TUNNELREMOVELABEL, theme.LogoutIcon(), func() {
			process.SetText(PROCESS_TUNNEL_DEL)
			if removeTunnel() != nil {
				process.SetText(PROCESS_TUNNEL_ERROR)
				status.SetText(getStatus())
				m.Items[1].Label = getStatus()
				m.Refresh()
			} else {
				process.SetText(PROCESS_TUNNELED_DEL)
				status.SetText(getStatus())
			}
			m.Items[1].Label = getStatus()
			m.Refresh()
			time.Sleep(5 * time.Second)
			process.SetText(HIDDEN)
		}),
	))
	m.Items[1].Label = getStatus()
	m.Refresh()
	w.ShowAndRun()
}
