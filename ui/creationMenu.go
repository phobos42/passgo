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

func createMenu() {
	var v = thisView
	menu := tview.NewForm().
		AddInputField(titleLabel, "", 20, nil, nil).
		AddDropDown(typeLabel, creationTypes, 0, nil).
		AddButton("Cancel", func() { switchToMain() }).
		AddButton("Create", func() { createNewItem() })
	menu.SetBorder(true).SetTitle("Menu")

	v.menuForm = menu
	model := createModel(menu, 40, 10)
	v.pages.AddPage("modal", model, true, true)
	v.app.SetFocus(v.menuForm)

}

func createNewItem() {
	var v = thisView
	//get form values
	title := v.menuForm.GetFormItemByLabel(titleLabel).(*tview.InputField).GetText()
	idx, _ := v.menuForm.GetFormItemByLabel(typeLabel).(*tview.DropDown).GetCurrentOption()
	var newthing interface{}
	switch idx {
	case 0:
		newthing = &types.Container{Title: title, Entries: []*types.Entry{}, Containers: []*types.Container{}}
	case 1:
		newthing = &types.Entry{Title: title, Items: []*types.Item{}}
	}

	// newthing := &types.Container{Title: title, Entries: []*types.Entry{}, Containers: []*types.Container{}}
	currentNode := v.tree.GetCurrentNode()
	currentNode.SetExpanded(true)

	// // //add to internal tree
	// container := currentNode.GetReference().(**types.Container)
	// (*container).Containers = append((*container).Containers, newthing)
	// lastidx := len((*container).Containers) - 1
	// v.infoBox.SetText(fmt.Sprintf("p:%p c:%p\n", (*container).Containers[lastidx], newthing))
	// fillTree(currentNode, &newthing)

	if r := addItem(currentNode, newthing); r != "" {
		v.infoBox.SetText(r)
	}
	switchToMain()
}
