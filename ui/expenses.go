// ui/expenses.go
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ExpensesScreen(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabel("💸 Dépenses")
	title.TextStyle = fyne.TextStyle{Bold: true}
	
	return container.NewVBox(title, widget.NewLabel("Liste des dépenses..."))
}