// ui/categories.go
package ui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"expense-tracker/data"
	"expense-tracker/ent"
	
)

func createCategoryRow(win fyne.Window, client *ent.Client, cat *ent.Category, refreshFunc func()) *fyne.Container {
	// === GAUCHE : Nom de la catégorie ===
	nameLabel := widget.NewLabel(fmt.Sprintf("%s %s", cat.Icon, cat.Name))
	nameLabel.TextStyle = fyne.TextStyle{Bold: false} // Optionnel

	// === DROITE : Boutons d'action ===
	editBtn := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
		showEditCategoryDialog(win, client, cat, refreshFunc)
	})
	editBtn.Importance = widget.LowImportance
	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		showDeleteCategoryDialog(win, client, cat, refreshFunc)
	})
	deleteBtn.Importance = widget.LowImportance

	// Grouper les boutons dans un HBox
	buttons := container.NewHBox(editBtn, deleteBtn)

	// === LAYOUT : Border layout ===
	// Gauche = nom, Droite = boutons, Centre = vide
	row := container.NewBorder(
		nil,    // Haut
		nil,    // Bas
		nameLabel, // Gauche
		buttons,   // Droite
		widget.NewLabel(""), // Centre (vide)
	)

	return row
}

// CategoriesScreen crée l'écran de gestion des catégories avec un design moderne
func CategoriesScreen(win fyne.Window) fyne.CanvasObject {
	client := data.GetDB()
	ctx := context.Background()
	var refreshList func()
	// === HAUT : Titre de la page ===
	title := widget.NewLabel("📋 Gestion des Catégories")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	// === CENTRE : Liste des catégories ===
	listContainer := container.NewVBox()
	refreshList = func() {
		listContainer.Objects = nil
		categories, _ := client.Category.Query().All(ctx)
		for _, cat := range categories {
			row := createCategoryRow(win, client, cat, refreshList) // <-- Utilisation ici
			listContainer.Add(row)
		}
		listContainer.Refresh()
	}
	refreshList()

	scroll := container.NewScroll(listContainer)
	scroll.SetMinSize(fyne.NewSize(400, 400))

	// === BAS : Bouton "Ajouter" ===
	addButton := widget.NewButton("➕ Ajouter une catégorie", func() {
		showAddCategoryDialog(win, client, refreshList)
	})
	buttonBox := container.NewCenter(addButton)

	// === BORDER LAYOUT ===
	return container.NewBorder(title, buttonBox, nil, nil, scroll)
}

// showAddCategoryDialog affiche une fenêtre modale pour ajouter une catégorie
func showAddCategoryDialog(win fyne.Window, client *ent.Client, refreshFunc func()) {
	ctx := context.Background()

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Nom de la catégorie")
	
	iconEntry := widget.NewEntry()
	iconEntry.SetPlaceHolder("Icône (ex: 🍽️)")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nom", Widget: nameEntry},
			{Text: "Icône", Widget: iconEntry},
		},
		OnSubmit: func() {
			name := nameEntry.Text
			icon := iconEntry.Text
			
			if name == "" {
				dialog.ShowInformation("Erreur", "Veuillez entrer un nom de catégorie.", win)
				return
			}
			
			// Créer la catégorie
			_, err := client.Category.Create().
				SetName(name).
				SetIcon(icon).
				Save(ctx)
			if err != nil {
				dialog.ShowInformation("Erreur", fmt.Sprintf("Impossible de créer la catégorie: %v", err), win)
				return
			}

			dialog.ShowInformation("Succès", "Catégorie ajoutée avec succès!", win)
			refreshFunc()
		},
	}

	dialog.ShowForm("Ajouter une Catégorie", "Créer", "Annuler", form.Items, func(b bool) {
		// Fermer la boîte de dialogue
	}, win)
}

// showEditCategoryDialog affiche une fenêtre modale pour modifier une catégorie
func showEditCategoryDialog(win fyne.Window, client *ent.Client, cat *ent.Category, refreshFunc func()) {
	ctx := context.Background()

	nameEntry := widget.NewEntry()
	nameEntry.SetText(cat.Name)
	
	iconEntry := widget.NewEntry()
	iconEntry.SetText(cat.Icon)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nom", Widget: nameEntry},
			{Text: "Icône", Widget: iconEntry},
		},
		OnSubmit: func() {
			name := nameEntry.Text
			icon := iconEntry.Text
			
			if name == "" {
				dialog.ShowInformation("Erreur", "Veuillez entrer un nom de catégorie.", win)
				return
			}
			
			// Mettre à jour la catégorie
			_, err := client.Category.UpdateOneID(cat.ID).
				SetName(name).
				SetIcon(icon).
				Save(ctx)
			if err != nil {
				dialog.ShowInformation("Erreur", fmt.Sprintf("Impossible de modifier la catégorie: %v", err), win)
				return
			}

			dialog.ShowInformation("Succès", "Catégorie modifiée avec succès!", win)
			refreshFunc()
		},
	}

	dialog.ShowForm("Modifier une Catégorie", "Enregistrer", "Annuler", form.Items, func(b bool) {
		// Fermer la boîte de dialogue
	}, win)
}

// showDeleteCategoryDialog affiche une confirmation pour supprimer une catégorie
func showDeleteCategoryDialog(win fyne.Window, client *ent.Client, cat *ent.Category, refreshFunc func()) {
	ctx := context.Background()
	
	confirm := dialog.NewConfirm(
		"Supprimer Catégorie",
		fmt.Sprintf("Êtes-vous sûr de vouloir supprimer la catégorie \"%s %s\" ?\nToutes les dépenses associées seront également affectées.", cat.Icon, cat.Name),
		func(confirmed bool) {
			if confirmed {
				// Supprimer la catégorie
				err := client.Category.DeleteOneID(cat.ID).Exec(ctx)
				if err != nil {
					dialog.ShowInformation("Erreur", fmt.Sprintf("Impossible de supprimer la catégorie: %v", err), win)
					return
				}
				
				dialog.ShowInformation("Succès", "Catégorie supprimée avec succès!", win)
				refreshFunc()
			}
		},
		win,
	)
	confirm.Show()
}