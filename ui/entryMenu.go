package ui

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

var creationItemTypes = []string{"text", "password"}

const valueLabel = "Value:"

func createEntryMenu() {
	var v = thisView
	clearInfoBox()
	currentNode := v.tree.GetCurrentNode()
	items := currentNode.GetChildren()

	outerMenu := tview.NewFlex().
		SetDirection(tview.FlexRow)

	itemContainer := tview.NewFlex().
		SetDirection(tview.FlexRow)

	fillEntriesMenu(currentNode, items, itemContainer)

	buttonFlex := tview.NewFlex()
	buttonFlex.SetDirection(tview.FlexColumn)
	buttonFlex.AddItem(tview.NewButton("Cancel").SetSelectedFunc(func() { switchToMain() }), btnSize2, 1, false)
	buttonFlex.AddItem(nil, 1, 1, false)
	buttonFlex.AddItem(tview.NewButton("New").SetSelectedFunc(func() { createNewItemMenu(currentNode, outerMenu) }), btnSize2, 1, false)
	buttonFlex.AddItem(nil, 1, 1, false)
	buttonFlex.AddItem(tview.NewButton("Done").SetSelectedFunc(func() { switchToMain() }), btnSize2, 1, false)

	outerMenu.AddItem(itemContainer, 0, 1, false)
	outerMenu.AddItem(buttonFlex, 1, 1, false)
	outerMenu.SetBorder(true).SetTitle("Entry Menu").SetBorderColor(v.colors.fg4).SetTitleColor(v.colors.fg0)

	v.menuFlex = outerMenu
	var menuWidth int
	if len(items) == 0 {
		menuWidth = 5
	} else {
		menuWidth = len(items)*4 + 3
	}
	model := createModel(outerMenu, 100, menuWidth)
	v.pages.AddPage("modal", model, true, true)
	v.app.SetFocus(v.mainView)
}

func fillEntriesMenu(currentNode *tview.TreeNode, items []*tview.TreeNode, flexContainer *tview.Flex) {

	for item := range items {
		itemForm := tview.NewForm()
		itemForm.SetHorizontal(true)

		if _, err := items[item].GetReference().(**types.Item); !err {
			thisView.infoBox.SetText("Recieved bad type for an item")
			return
		}
		//safe to cast after check
		ref := *items[item].GetReference().(**types.Item)
		itemType := ref.Type
		title := ref.Title
		value := ref.Value
		switch itemType {
		case "text":
			itemForm.AddInputField(title, value, 30, nil, nil).SetLabelColor(thisView.colors.aqua)
			itemForm.AddButton("Save", func() { ref.Value = (itemForm.GetFormItemByLabel(title).(*tview.InputField)).GetText() })
			itemForm.AddButton("Revert", func() { (itemForm.GetFormItemByLabel(title).(*tview.InputField)).SetText(ref.Value) })
			itemForm.AddButton("Copy", func() { writeToClipboard(itemForm.GetFormItemByLabel(title).(*tview.InputField).GetText()) })
			itemForm.AddButton("Delete", func() { createConfirmMenu(currentNode, items[item]) })
		case "password":
			itemForm.AddPasswordField(title, value, 30, '*', nil).SetLabelColor(thisView.colors.aqua)
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
			itemForm.AddButton("Delete", func() { createConfirmMenu(currentNode, items[item]) })
		default:
			thisView.infoBox.SetText("Unhandled Item type found...")
		}
		flexContainer.AddItem(itemForm, 0, 1, true)
	}
}

func createNewItemMenu(parentNode *tview.TreeNode, flexContainer *tview.Flex) {
	var v = thisView
	//make menu for form
	ref := parentNode.GetReference()
	if _, err := ref.(**types.Entry); !err {
		v.infoBox.SetText("Can only create Item inside entry")
		return
	}

	menu := tview.NewForm()
	menu.
		AddInputField(titleLabel, "", 20, nil, nil).
		AddDropDown(typeLabel, creationItemTypes, 0, nil).
		AddInputField(valueLabel, "", 20, nil, nil).
		AddButton("Cancel", func() { v.pages.RemovePage("modal2") }).
		AddButton("Create", func() {
			_, str := (menu.GetFormItemByLabel(typeLabel).(*tview.DropDown)).GetCurrentOption()
			createNewItem(parentNode,
				(menu.GetFormItemByLabel(titleLabel).(*tview.InputField)).GetText(),
				str,
				(menu.GetFormItemByLabel(valueLabel).(*tview.InputField)).GetText())
			//remove menu for item creation
			v.pages.RemovePage("modal2")
			//remake entry menu to include new value
			createEntryMenu()
		})
	menu.SetBorder(true).SetTitle("Menu").SetBorderColor(v.colors.fg4).SetTitleColor(v.colors.fg0)

	v.menuForm = menu
	model := createModel(menu, 40, 15)
	v.pages.AddPage("modal2", model, true, true)
	v.app.SetFocus(v.menuForm)

}
func createNewItem(parentNode *tview.TreeNode, title string, typestr string, value string) {
	//create new node
	newItem := &types.Item{Title: title, Type: typestr, Value: value}

	if r := addItem(parentNode, newItem); r != "" {
		thisView.infoBox.SetText(r)
	}
}
