package uiparts

//package level vars

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	types "github.com/phobos42/passgo/utils"
)

const buttonWidth = 10

type Selection struct {
	flp         *types.FoldersLayoutParts
	elp         *types.EntriesLayoutParts
	fieldlp     *types.FieldsLayoutParts
	outerLayout *tview.Flex
	data        *types.Allfolders
	app         *tview.Application
}

var uiPart = &Selection{}

func initSelectionLayout() {

}

//CreateSelectionLayout
func CreateSelectionLayout(data *types.Allfolders, app *tview.Application) *tview.Flex {
	uiPart.data = data
	uiPart.app = app

	uiPart.flp = createFoldersPart()
	uiPart.elp = createEntriesPart()
	uiPart.fieldlp = createFieldsPart()

	setupOuterLayout()
	return uiPart.outerLayout
}

func createButtonLayout(newButton func(), editButton func(), deleteButton func()) *tview.Flex {
	newTButton := tview.NewButton("New").SetSelectedFunc(newButton)
	editTButton := tview.NewButton("Edit").SetSelectedFunc(editButton)
	deleteTButton := tview.NewButton("Delete").SetSelectedFunc(deleteButton)

	buttonLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	buttonLayout.AddItem(newTButton, buttonWidth, 1, false)
	buttonLayout.AddItem(editTButton, buttonWidth, 1, false)
	buttonLayout.AddItem(deleteTButton, buttonWidth, 1, false)

	return buttonLayout
}

func createBottomRowLayout(buttonLayout *tview.Flex, promptLayout *tview.Flex) *tview.Flex {
	bottomRowLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	bottomRowLayout.AddItem(buttonLayout, 30, 2, false)
	bottomRowLayout.AddItem(promptLayout, 60, 8, false)

	return bottomRowLayout
}

func createPromptLayout() *tview.Flex {
	prompt := tview.NewTextView().SetText("Prompt")
	inputField := tview.NewInputField()

	promptLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	promptLayout.AddItem(prompt, 20, 1, false)
	promptLayout.AddItem(inputField, 40, 1, false)

	return promptLayout
}

func createInnerList() *tview.List {
	list := tview.NewList()
	setlistproperties(list)
	return list
}

func setupOuterLayout() {
	uiPart.outerLayout = tview.NewFlex()
	uiPart.outerLayout.SetDirection(tview.FlexRow)
	uiPart.outerLayout.AddItem(uiPart.flp.EnclosingLayout, 0, 30, true)
	uiPart.outerLayout.AddItem(uiPart.elp.EnclosingLayout, 0, 10, false)
	uiPart.outerLayout.AddItem(uiPart.fieldlp.EnclosingLayout, 0, 10, false)
}
func setlistproperties(list *tview.List) {
	list.ShowSecondaryText(false)
	list.SetTitleAlign(0)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetSelectedBackgroundColor(tcell.ColorDefault)
	list.SetSelectedTextColor(tcell.ColorBlue)
	list.SetSelectedFocusOnly(false)
}
func setLayoutProperties(title string, layout *tview.Flex) {
	layout.SetTitle(title)
	layout.SetBorder(true)
	layout.SetTitleAlign(0)
}

func FillUI(folders types.Allfolders) {
	fillFolders(folders)

}
