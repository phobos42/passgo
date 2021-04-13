package ui

import (
	"github.com/rivo/tview"
)

func createConfirmMenu(parentNode *tview.TreeNode, targetNode *tview.TreeNode) {
	var v = thisView
	//make menu for form

	menu := tview.NewForm()
	menu.AddButton("No", func() { v.pages.RemovePage("modal3") }).
		AddButton("Yes", func() {
			parentNode.RemoveChild(targetNode)
			deleteChild(parentNode, targetNode)
			//remove menu for item creation
			v.pages.RemovePage("modal3")
			//remake entry menu to include new value if it was open
			if v.pages.GetPageCount() == 2 {
				createEntryMenu()
			}
		})
	menu.SetBorder(true).SetTitle("Are You Sure?").SetBorderColor(v.colors.fg4).SetTitleColor(v.colors.fg0)

	v.menuForm = menu
	model := createModel(menu, 20, 5)
	v.pages.AddPage("modal3", model, true, true)
	v.app.SetFocus(v.menuForm)

}
