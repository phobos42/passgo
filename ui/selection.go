package uiparts

import (
	"github.com/rivo/tview"

	types "github.com/phobos42/passgo/utils"
)

type Selection struct {
	outerLayout   *tview.Flex
	foldersLayout *tview.Flex
	foldersList   *tview.List
	entriesLayout *tview.Flex
	entriesList   *tview.List
	fieldsLayout  *tview.Flex
	fieldsList    *tview.List
	data          *types.Allfolders
	app           *tview.Application
}

var selectionPart Selection

func initSelectionLayout() {

}

//CreateSelectionLayout
func CreateSelectionLayout(data *types.Allfolders, app *tview.Application) *tview.Flex {
	selectionPart.data = data
	selectionPart.app = app
	setupFolders()
	setupEntries()
	setupFields()
	setupOuterLayout()
	return selectionPart.outerLayout
}

func setupFolders() {
	selectionPart.foldersList = tview.NewList()
	selectionPart.foldersList.SetTitle("Folders")
	setlistproperties(selectionPart.foldersList)
	selectionPart.foldersLayout = tview.NewFlex()
	selectionPart.foldersLayout.AddItem(selectionPart.foldersList, 0, 1, true)
}
func setupEntries() {
	selectionPart.entriesList = tview.NewList()
	selectionPart.entriesList.SetTitle("Entries")
	setlistproperties(selectionPart.entriesList)
	selectionPart.entriesLayout = tview.NewFlex()
	selectionPart.entriesLayout.AddItem(selectionPart.entriesList, 0, 1, true)
}
func setupFields() {
	selectionPart.fieldsList = tview.NewList()
	selectionPart.fieldsList.SetTitle("Fields")
	setlistproperties(selectionPart.fieldsList)
	selectionPart.fieldsLayout = tview.NewFlex()
	selectionPart.fieldsLayout.AddItem(selectionPart.fieldsList, 0, 1, true)

}
func setupOuterLayout() {
	selectionPart.outerLayout = tview.NewFlex()
	selectionPart.outerLayout.SetDirection(tview.FlexRow)
	selectionPart.outerLayout.AddItem(selectionPart.foldersLayout, 0, 30, true)
	selectionPart.outerLayout.AddItem(selectionPart.entriesLayout, 0, 10, false)
	selectionPart.outerLayout.AddItem(selectionPart.fieldsLayout, 0, 10, false)
}
func setlistproperties(list *tview.List) {
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	list.SetTitleAlign(0)
}

func FillUI(folders types.Allfolders) {
	fillFolders(folders)

}

func fillFolders(folders types.Allfolders) {
	for folderName := range folders {
		selectionPart.foldersList.AddItem(folderName, "", 0, folderSelected)
	}

}
func fillEntries(entries types.Folder) {
	for entryName := range entries {
		selectionPart.entriesList.AddItem(entryName, "", 0, entrySelected)
	}
}
func fillFields(fields types.Entry) {
	for index := range fields {
		selectionPart.fieldsList.AddItem(fields[index].Title, "", 0, nil)
	}
}

func folderSelected() {
	index := selectionPart.foldersList.GetCurrentItem()
	name, _ := selectionPart.foldersList.GetItemText(index)
	selectionPart.entriesList.Clear()
	fillEntries((*selectionPart.data)[name])
	selectionPart.app.SetFocus(selectionPart.entriesLayout)
	selectionPart.outerLayout.ResizeItem(selectionPart.foldersLayout, 0, 10)

}
func entrySelected() {
	folderIndex := selectionPart.foldersList.GetCurrentItem()
	folderName, _ := selectionPart.foldersList.GetItemText(folderIndex)
	index := selectionPart.entriesList.GetCurrentItem()
	name, _ := selectionPart.entriesList.GetItemText(index)
	selectionPart.fieldsList.Clear()
	fillFields((*selectionPart.data)[folderName][name])
	selectionPart.app.SetFocus(selectionPart.fieldsLayout)
}
