package main

import (
    "context"
    "fmt"
    "log"
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    _ "modernc.org/sqlite"
    "expense-tracker/ent"
)

func main() {
    client, err := ent.Open("sqlite3", "file:expenses.db?_fk=1")
    if err != nil {
        log.Fatalf("failed opening sqlite connection: %v", err)
    }
    defer client.Close()

    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    // App Fyne
    myApp := app.New()
    myWindow := myApp.NewWindow("Expense Tracker")

    titleEntry := widget.NewEntry()
    titleEntry.SetPlaceHolder("Nom de la dépense")

    amountEntry := widget.NewEntry()
    amountEntry.SetPlaceHolder("Montant")

    addBtn := widget.NewButton("Ajouter", func() {
        expense, err := client.Expense.
            Create().
            SetTitle(titleEntry.Text).
            SetAmount(parseFloat(amountEntry.Text)).
            Save(ctx)
        if err != nil {
            log.Println("Erreur ajout:", err)
            return
        }
        fmt.Println("Nouvelle dépense ajoutée:", expense)
    })

    myWindow.SetContent(
        container.NewVBox(
            widget.NewLabel("Ajouter une dépense"),
            titleEntry,
            amountEntry,
            addBtn,
        ),
    )

    myWindow.ShowAndRun()
}

func parseFloat(s string) float64 {
    var f float64
    fmt.Sscanf(s, "%f", &f)
    return f
}
