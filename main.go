// main.go
package main

import (
    "fyne.io/fyne/v2/app"
	"expense-tracker/data"
	"expense-tracker/theme"
	"expense-tracker/ui"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "time"
    "fyne.io/fyne/v2/canvas"
)

func main() {
	
    myApp := app.New()
    myApp.Settings().SetTheme(&theme.CustomDarkTheme{})
    ShowSplashScreenWithLogo(myApp, func() {
        showMainWindow(myApp)
    })

    myApp.Run()
}

func showMainWindow(myApp fyne.App) {
    // Initialiser la base de données
	data.InitDB()
	
	// Appliquer le thème sombre personnalisé
	myApp.Settings().SetTheme(&theme.CustomDarkTheme{})

	// Créer la fenêtre principale
	myWindow := myApp.NewWindow("MonBudget")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Créer le menu principal
	menu := ui.MainMenu(myWindow)

	// Afficher le menu
	myWindow.SetContent(menu)
	myWindow.Show()
}
func ShowSplashScreenWithLogo(app fyne.App, onComplete func()) {
	splashWindow := app.NewWindow("MonBudget")
	splashWindow.SetFixedSize(true)
	splashWindow.Resize(fyne.NewSize(400, 600))
	splashWindow.CenterOnScreen()

	// Si vous avez un fichier logo.png
	logo := canvas.NewImageFromFile("assets/logo.png")
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(100, 100))

	title := widget.NewLabel(" 예산 MonBudget")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	version := widget.NewLabel("v1.0.0")
	version.Alignment = fyne.TextAlignCenter

	progress := widget.NewProgressBarInfinite()
	progress.Start()

	content := container.NewVBox(
		widget.NewSeparator(),
		logo, // Décommentez si vous avez un logo
		title,
		version,
		widget.NewSeparator(),
		progress,
		widget.NewSeparator(),
	)

	splashWindow.SetContent(container.NewCenter(content))
	splashWindow.Show()
	// Fermer après 3 secondes
	go func() {
		time.Sleep(3 * time.Second)
		onComplete()       // ouvre la fenêtre principale
        splashWindow.Close() // puis ferme le splash
	}()
}