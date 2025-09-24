// ui/profile.go
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ProfileScreen(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel("👤 Profil")
	title.TextStyle = fyne.TextStyle{Bold: true}
	
	return container.NewVBox(title, widget.NewLabel("Paramètres utilisateur..."))
}