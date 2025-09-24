// ui/menu.go
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

// MainMenu crée le menu de navigation en bas avec 5 onglets
func MainMenu(win fyne.Window) *container.AppTabs {
	// Créer les onglets
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Tableau de bord", theme.HomeIcon(), DashboardScreen(win)),
		container.NewTabItemWithIcon("Dépenses", theme.SearchIcon(), ExpensesScreen(win)),
		container.NewTabItemWithIcon("Catégories", theme.ViewRefreshIcon(), CategoriesScreen(win)),
		container.NewTabItemWithIcon("Budgets", theme.ContentPasteIcon(), BudgetScreen(win)),
		container.NewTabItemWithIcon("Profil", theme.AccountIcon(), ProfileScreen(win)),
	)

	// Style : menu en bas
	tabs.SetTabLocation(container.TabLocationBottom)

	return tabs
}