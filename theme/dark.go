// theme/dark.go
package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// CustomDarkTheme est un thème sombre personnalisé pour l'application MonBudget
type CustomDarkTheme struct{}

var _ fyne.Theme = (*CustomDarkTheme)(nil)

// Color retourne la couleur associée à un nom de couleur de thème spécifique
func (c CustomDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 0x12, G: 0x12, B: 0x12, A: 0xff} // #121212 - Fond sombre
	case theme.ColorNameButton:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0xff} // #4FC3F7 - Bleu clair pour les boutons
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0xff} // #4FC3F7 - Couleur principale
	case theme.ColorNameFocus:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0xff} // #4FC3F7 - Couleur de focus
	case theme.ColorNameForeground:
		return color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff} // Blanc - Texte principal
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff} // Gris - Éléments désactivés
	case theme.ColorNameSuccess:
		return color.RGBA{R: 0x4c, G: 0xaf, B: 0x50, A: 0xff} // Vert - Succès
	case theme.ColorNameWarning:
		return color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 0xff} // Jaune - Avertissement
	case theme.ColorNameError:
		return color.RGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff} // Rouge - Erreur
	case theme.ColorNameHover:
		return color.RGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xff} // Gris foncé - Survol
	case theme.ColorNameHeaderBackground:
		return color.RGBA{R: 0x1e, G: 0x1e, B: 0x1e, A: 0xff} // #1e1e1e - Fond d'en-tête
	case theme.ColorNameInputBorder:
		return color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff} // Gris - Bordure d'entrée
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 0x1e, G: 0x1e, B: 0x1e, A: 0xff} // #1e1e1e - Fond d'entrée
	case theme.ColorNameSelection:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0x88} // #4FC3F7 avec transparence - Sélection
	default:
		// Pour les couleurs non définies, utiliser les valeurs par défaut du thème sombre
		// On retourne les couleurs par défaut pour le thème sombre
		return defaultDarkColor(name)
	}
}

// defaultDarkColor retourne les couleurs par défaut pour le thème sombre
func defaultDarkColor(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.Black
	case theme.ColorNameButton:
		return color.RGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xff}
	case theme.ColorNameDisabled:
		return color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}
	case theme.ColorNameError:
		return color.RGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameFocus:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0xff}
	case theme.ColorNameForeground:
		return color.White
	case theme.ColorNameHover:
		return color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	case theme.ColorNamePressed:
		return color.RGBA{R: 0x11, G: 0x11, B: 0x11, A: 0xff}
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x4f, G: 0xc3, B: 0xf7, A: 0xff}
	case theme.ColorNameScrollBar:
		return color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
	case theme.ColorNameSeparator:
		return color.RGBA{R: 0x33, G: 0x33, B: 0x33, A: 0xff}
	case theme.ColorNameShadow:
		return color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x33}
	case theme.ColorNameSuccess:
		return color.RGBA{R: 0x4c, G: 0xaf, B: 0x50, A: 0xff}
	case theme.ColorNameWarning:
		return color.RGBA{R: 0xff, G: 0xc1, B: 0x07, A: 0xff}
	default:
		return color.Black
	}
}

// Icon retourne la ressource d'icône associée à un nom d'icône de thème spécifique
func (c CustomDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// Utiliser les icônes par défaut
	return theme.DefaultTheme().Icon(name)
}

// Font retourne la ressource de police associée à un style de texte spécifique
func (c CustomDarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	// Utiliser les polices par défaut
	return theme.DefaultTheme().Font(style)
}

// Size retourne la taille associée à un nom de taille de thème spécifique
func (c CustomDarkTheme) Size(name fyne.ThemeSizeName) float32 {
	// Utiliser les tailles par défaut
	return theme.DefaultTheme().Size(name)
}