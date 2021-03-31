package ui

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func createEntryMenu() {
	var v = thisView
	currentNode := v.tree.GetCurrentNode()
	items := currentNode.GetChildren()

	outerMenu := tview.NewFlex().
		SetDirection(tview.FlexRow)

	for item := range items {
		itemForm := tview.NewForm()
		itemForm.SetHorizontal(true)

		if _, err := items[item].GetReference().(**types.Item); !err {
			v.infoBox.SetText("Recieved bad type for an item")
			return
		}
		//safe to cast after check
		ref := *items[item].GetReference().(**types.Item)
		itemType := ref.Type
		title := ref.Title
		value := ref.Value
		switch itemType {
		case "text":
			itemForm.AddInputField(title, value, 30, nil, nil)
			itemForm.AddButton("Save", func() { ref.Value = (itemForm.GetFormItemByLabel(title).(*tview.InputField)).GetText() })
			itemForm.AddButton("Revert", func() { (itemForm.GetFormItemByLabel(title).(*tview.InputField)).SetText(ref.Value) })
			itemForm.AddButton("Copy", func() { writeToClipboard(itemForm.GetFormItemByLabel(title).(*tview.InputField).GetText()) })
		case "password":
			itemForm.AddPasswordField(title, value, 30, '*', nil)
			itemForm.AddCheckbox("Show", false, func(checked bool) {
				if checked {
					itemForm.GetFormItemByLabel(title).(*tview.InputField).SetMaskCharacter(0)
				} else {
					itemForm.GetFormItemByLabel(title).(*tview.InputField).SetMaskCharacter('*')
				}
			})
			itemForm.AddButton("Save", func() { ref.Value = (itemForm.GetFormItemByLabel(title).(*tview.InputField)).GetText() })
			itemForm.AddButton("Revert", func() { (itemForm.GetFormItemByLabel(title).(*tview.InputField)).SetText(ref.Value) })
			itemForm.AddButton("Copy", func() { writeToClipboard(itemForm.GetFormItemByLabel(title).(*tview.InputField).GetText()) })
		default:
			v.infoBox.SetText("Unhandled Item type found...")
		}

		outerMenu.AddItem(itemForm, 0, 1, true)

	}

	buttonFlex := tview.NewFlex()
	buttonFlex.SetDirection(tview.FlexColumn)
	buttonFlex.AddItem(tview.NewButton("Cancel").SetSelectedFunc(func() { switchToMain() }), btnSize2, 1, false)
	//buttonFlex.AddItem(nil, 1, 1, false)
	//buttonFlex.AddItem(tview.NewButton("Submit Changes").SetSelectedFunc(func() { switchToMain() }), btnSize3, 1, false)

	outerMenu.AddItem(buttonFlex, 1, 1, false)
	outerMenu.SetBorder(true).SetTitle("Entry Menu")

	v.menuFlex = outerMenu
	model := createModel(outerMenu, 100, len(items)*4+1)
	v.pages.AddPage("modal", model, true, true)
	v.app.SetFocus(v.menuFlex)
}
