// ui/dashboard.go
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

func DashboardScreen(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel("📊 Tableau de bord")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	welcome := widget.NewLabel("Bienvenue dans MonBudget !")
	listContainer := container.NewVBox()
	scroll := container.NewScroll(listContainer)
	scroll.SetMinSize(fyne.NewSize(400, 400))
	// === BAS : Bouton "Ajouter" ===
	addButton := widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		win.SetContent(DashboardScreen(win))
	})
	addButton.Importance = widget.LowImportance
	buttonBox := container.NewCenter(addButton)

	return container.NewVBox(title, welcome, widget.NewSeparator(), scroll, buttonBox)
}