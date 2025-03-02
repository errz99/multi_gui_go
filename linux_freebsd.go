//go:build linux || freebsd

package main

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func run() {
	app := gtk.NewApplication("com.example.multiplatform_gui", glib.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		win := gtk.NewApplicationWindow(app)
		win.SetTitle("Multiplatform GUI")
		win.SetDefaultSize(800, 600)
		win.Show()
	})
	app.Run()
}
