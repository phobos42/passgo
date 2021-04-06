package ui

import (
	"fmt"
	"reflect"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const (
	newTitleLabel = "New Title:"
)

func editTitleMenu() {
	var v = thisView
	clearInfoBox()

	menu := tview.NewForm().
		AddInputField(newTitleLabel, "", 20, nil, nil).
		AddButton("Cancel", func() { switchToMain() }).
		AddButton("Submit Change", func() { editTitle() })
	menu.SetBorder(true).SetTitle("Edit Title Menu").SetBorderColor(v.colors.fg4).SetTitleColor(v.colors.fg0)

	v.menuForm = menu
	model := createModel(menu, 40, 10)
	v.pages.AddPage("modal", model, true, true)
	v.app.SetFocus(v.menuForm)
}

func editTitle() {
	var v = thisView
	//get form values
	newTitle := v.menuForm.GetFormItemByLabel(newTitleLabel).(*tview.InputField).GetText()

	currentNode := v.tree.GetCurrentNode()
	reference := currentNode.GetReference()
	switch t := reference.(type) {
	case **types.Container:
		(*t).Title = newTitle
	case **types.Entry:
		(*t).Title = newTitle
	case **types.Item:
		(*t).Title = newTitle
	default:
		v.infoBox.SetText("Cannot edit title of type: " + fmt.Sprint(reflect.TypeOf(t)))
		switchToMain()
		return
	}
	currentNode.SetText(newTitle)

	switchToMain()
}
