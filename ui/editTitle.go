package ui

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const (
	newTitleLabel = "New Title:"
)

func editTitleMenu(v *View) {
	menu := tview.NewForm().
		AddInputField(newTitleLabel, "", 20, nil, nil).
		AddButton("Cancel", func() { switchToMain(v) }).
		AddButton("Submit Change", func() { editTitle(v) })
	menu.SetBorder(true).SetTitle("Edit Title Menu")

	v.menuForm = menu
	model := createModel(menu, 40, 10)
	v.pages.AddPage("modal", model, true, true)
	v.app.SetFocus(v.menuForm)
}

func editTitle(v *View) {
	//get form values
	newTitle := v.menuForm.GetFormItemByLabel(newTitleLabel).(*tview.InputField).GetText()

	currentNode := v.tree.GetCurrentNode()
	reference := currentNode.GetReference()
	switch t := reference.(type) {
	case types.Container:
		t.Title = newTitle
	case types.Entry:
		t.Title = newTitle
	default:
		v.infoBox.SetText("Edit Items in Entry menu")
		switchToMain(v)
		return
	}
	currentNode.SetText(newTitle)

	switchToMain(v)
}
