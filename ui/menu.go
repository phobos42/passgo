package ui

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const (
	titleLabel = "Title:"
	typeLabel  = "Type"
)

var creationTypes = []string{"Container", "Entry"}

func createMenu(v *View) {

	menu := tview.NewForm().
		AddInputField(titleLabel, "", 20, nil, nil).
		AddDropDown(typeLabel, creationTypes, 0, nil).
		AddButton("Cancel", func() { switchToMain(v) }).
		AddButton("Create", func() { createNewItem(v) })
	menu.SetBorder(true).SetTitle("Menu")

	v.menuForm = menu
	model := createModel(menu, 40, 10)
	v.pages.AddPage("modal", model, true, true)

}

func createModel(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, false).
			AddItem(nil, 0, 1, false), width, 1, false).
		AddItem(nil, 0, 1, false)
}

func createNewItem(v *View) {
	//get form values
	title := v.menuForm.GetFormItemByLabel(titleLabel).(*tview.InputField).GetText()
	idx, _ := v.menuForm.GetFormItemByLabel(typeLabel).(*tview.DropDown).GetCurrentOption()
	var newthing interface{}
	switch idx {
	case 0:
		newthing = types.Container{Title: title, Entries: []types.Entry{}, Containers: []types.Container{}}
	case 1:
		newthing = types.Entry{Title: title, Items: []types.Item{}}
	}

	currentNode := v.tree.GetCurrentNode().SetExpanded(true)
	//v.infoBox.SetText(addItem(currentNode, newthing))
	if r := addItem(currentNode, newthing); r != "" {
		v.infoBox.SetText(r)
	}
	switchToMain(v)
}
