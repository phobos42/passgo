package uiparts

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func createEntriesPart() *types.EntriesLayoutParts {
	var ret types.EntriesLayoutParts
	ret.ButtonLayout = createButtonLayout(newEntryButton, editEntryButton, deleteEntryButton)
	ret.PromptLayout = createPromptLayout()
	ret.BottomRowLayout = createBottomRowLayout(ret.ButtonLayout, ret.PromptLayout)
	ret.InnerList = createInnerList()

	ret.EnclosingLayout = tview.NewFlex()
	ret.EnclosingLayout.SetDirection(tview.FlexRow)
	setLayoutProperties("Entries", ret.EnclosingLayout)
	ret.EnclosingLayout.AddItem(ret.InnerList, 0, 9, false)
	ret.EnclosingLayout.AddItem(ret.BottomRowLayout, 1, 1, false)
	return &ret
}

func fillEntries(entries types.Folder) {
	for entryName := range entries {
		uiPart.elp.InnerList.AddItem(entryName, "", 0, entrySelected)
	}
}

func entrySelected() {
	folderIndex := uiPart.flp.InnerList.GetCurrentItem()
	folderName, _ := uiPart.flp.InnerList.GetItemText(folderIndex)
	index := uiPart.elp.InnerList.GetCurrentItem()
	name, _ := uiPart.elp.InnerList.GetItemText(index)
	uiPart.fieldlp.InnerList.Clear()
	fillFields((*uiPart.data)[folderName][name])
	//uiPart.app.SetFocus(uiPart.fieldsLayout)
}

func newEntryButton() {

}
func editEntryButton() {

}
func deleteEntryButton() {

}
