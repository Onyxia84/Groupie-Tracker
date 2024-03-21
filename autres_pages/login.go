package autres_pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoginPage(w *fyne.Window) {
	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Login")
	login := container.New(layout.NewPaddedLayout(), loginEntry)

	mdpEntry := widget.NewPasswordEntry()
	mdpEntry.SetPlaceHolder("Mot de passe")
	mdp := container.New(layout.NewPaddedLayout(), mdpEntry)
	mdpEntry.Resize(fyne.NewSize(200, 30))

	btn := widget.NewButton("", func() {
		pagePrincipale(w)
	})
	btn_colour := canvas.NewRectangle(color.RGBA{0, 240, 0, 1})
	btn_icon := widget.NewIcon(theme.ConfirmIcon())
	btn_icon.Resize(fyne.NewSize(30, 30))
	btn_icon.Move(fyne.NewPos(30, -20))
	btn_icon.Refresh()
	login_btn := container.New(layout.NewPaddedLayout(), btn, btn_icon, btn_colour)

	title := canvas.NewText("Groupie Tracker", color.RGBA{255, 255, 255, 1})
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true

	containerPage := container.NewGridWithRows(3,
		title,
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewGridWithRows(3,
				login,
				mdp,
				login_btn,
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
	v := *w
	v.SetContent(containerPage)
	w = &v
}
