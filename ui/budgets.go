// ui/budgets.go
package ui

import (
	"context"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"expense-tracker/data"
	"expense-tracker/ent"

)

// BudgetScreen crée l'écran de gestion des budgets
func BudgetScreen(win fyne.Window) fyne.CanvasObject {
	client := data.GetDB()
	ctx := context.Background()

	// Récupérer l'utilisateur (pour cet exemple, on prend le premier)
	currentUser, err := client.User.Query().First(ctx)
	if err != nil {
		// Si aucun utilisateur, on en crée un par défaut
		currentUser, err = client.User.Create().
			SetName("Utilisateur").
			SetCurrency("€").
			Save(ctx)
		if err != nil {
			log.Fatalf("Impossible de créer l'utilisateur: %v", err)
		}
	}

	// Récupérer les catégories existantes ou les créer
	categories, err := client.Category.Query().All(ctx)
	if err != nil || len(categories) == 0 {
		// Création des catégories par défaut
		defaultCategories := []struct {
			Name string
			Icon string
		}{
			{"Alimentation", "🍽️"},
			{"Transport", "🚗"},
			{"Divertissement", "🎬"},
			{"Shopping", "🛍️"},
			{"Logement", "🏠"},
		}
		for _, cat := range defaultCategories {
			_, err := client.Category.Create().
				SetName(cat.Name).
				SetIcon(cat.Icon).
				Save(ctx)
			if err != nil {
				log.Printf("Erreur lors de la création de la catégorie %s: %v", cat.Name, err)
			}
		}
		categories, _ = client.Category.Query().All(ctx)
	}

	// Liste des budgets
	list := widget.NewList(
		func() int {
			budgets, _ := client.Budget.Query().WithCategory().All(ctx)
			return len(budgets)
		},
		func() fyne.CanvasObject {
			card := widget.NewCard("", "", widget.NewLabel(""))
			return card
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			budgets, _ := client.Budget.Query().WithCategory().All(ctx)
			if id >= len(budgets) {
				return
			}
			b := budgets[id]
			card := obj.(*widget.Card)
			card.SetTitle(fmt.Sprintf("%s %s", b.Edges.Category.Icon, b.Edges.Category.Name))
			card.SetSubTitle(fmt.Sprintf("Budget: %.2f %s", b.Amount, currentUser.Currency))
		},
	)

	// Bouton pour ajouter un budget
	addButton := widget.NewButton("Ajouter Budget", func() {
		showAddBudgetDialog(win, client, currentUser.ID, func() {
			list.Refresh()
		})
	})

	// Conteneur principal
	content := container.NewBorder(nil, addButton, nil, nil, list)
	return content
}

// showAddBudgetDialog affiche une fenêtre modale pour ajouter un budget
func showAddBudgetDialog(win fyne.Window, client *ent.Client, userID int, refreshFunc func()) {
	ctx := context.Background()

	// Récupérer les catégories
	categories, err := client.Category.Query().All(ctx)
	if err != nil || len(categories) == 0 {
		dialog.ShowInformation("Erreur", "Aucune catégorie disponible.", win)
		return
	}

	// Champs du formulaire
	categorySelect := widget.NewSelect(make([]string, len(categories)), func(s string) {})
	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Montant")

	// Remplir le sélecteur de catégories
	categoryNames := make([]string, len(categories))
	for i, cat := range categories {
		categoryNames[i] = fmt.Sprintf("%s %s", cat.Icon, cat.Name)
	}
	categorySelect.Options = categoryNames

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Catégorie", Widget: categorySelect},
			{Text: "Montant", Widget: amountEntry},
		},
		OnSubmit: func() {
			selectedIndex := categorySelect.SelectedIndex()
			if selectedIndex == -1 {
				dialog.ShowInformation("Erreur", "Veuillez sélectionner une catégorie.", win)
				return
			}
			cat := categories[selectedIndex]

			var amount float64
			_, err := fmt.Sscanf(amountEntry.Text, "%f", &amount)
			if err != nil || amount <= 0 {
				dialog.ShowInformation("Erreur", "Veuillez entrer un montant valide.", win)
				return
			}

			// Créer le budget
			_, err = client.Budget.Create().
				SetAmount(amount).
				SetStartDate(time.Now()).
				SetEndDate(time.Now().AddDate(0, 1, 0)). // Budget mensuel
				SetCategoryID(cat.ID).
				SetUserID(userID).
				Save(ctx)
			if err != nil {
				dialog.ShowInformation("Erreur", fmt.Sprintf("Impossible de créer le budget: %v", err), win)
				return
			}

			dialog.ShowInformation("Succès", "Budget ajouté avec succès!", win)
			refreshFunc()
		},
	}

	dialog.ShowForm("Ajouter un Budget", "Créer", "Annuler", form.Items, func(b bool) {
		// Fermer la boîte de dialogue
	}, win)
}