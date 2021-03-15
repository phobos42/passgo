package main

import (
	uiparts "github.com/phobos42/passgo/ui"
	"github.com/rivo/tview"
)

type userInterface struct {
	app             *tview.Application
	masterlayout    *tview.Flex
	selectionLayout *tview.Flex
	passwordForm    *tview.Form
}

func initUI() {
	application.ui.app = tview.NewApplication()

	//keybindings

	//configs

	//createlayout
	application.ui.passwordForm = uiparts.CreateGeneratorForm(application.ui.app)
	createSelectionLayout()

	createMasterLayout()

	uiparts.FillUI(application.folders)
}

func showUI() {
	application.ui.app.SetRoot(application.ui.masterlayout, true)
	// application.ui.app.SetFocus()
	application.ui.app.EnableMouse(true)
	// application.ui.app.Draw()
}

func runUI() {
	if err := application.ui.app.Run(); err != nil {
		panic(err)
	}
}
func createSelectionLayout() {
	thing := uiparts.CreateSelectionLayout(&application.folders, application.ui.app)
	application.ui.selectionLayout = thing
}
func createMasterLayout() {
	application.ui.masterlayout = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(application.ui.selectionLayout, 0, 2, true).
		AddItem(application.ui.passwordForm, 0, 1, false)

}
