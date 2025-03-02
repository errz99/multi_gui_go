//go:build windows

package main

import (
	"github.com/gen2brain/iup-go/iup"
)

func run() {
	iup.Open()
	defer iup.Close()

	dlg := iup.Dialog(iup.Vbox(
		iup.Label("Multiplatform GUI"),
	)).SetAttributes(`TITLE=Multiplatform GUI, SIZE=800x600`)

	// dlg.Show()
	iup.ShowXY(dlg, iup.CENTER, iup.CENTER)
	iup.MainLoop()
}
