package uiparts

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func createFoldersPart() *types.FoldersLayoutParts {
	var ret types.FoldersLayoutParts
	ret.ButtonLayout = createButtonLayout(newFolderButton, editFolderButton, deleteFolderButton)
	ret.PromptLayout = createPromptLayout()
	ret.BottomRowLayout = createBottomRowLayout(ret.ButtonLayout, ret.PromptLayout)
	ret.InnerList = createInnerList()

	ret.EnclosingLayout = tview.NewFlex()
	ret.EnclosingLayout.SetDirection(tview.FlexRow)
	setLayoutProperties("Folders", ret.EnclosingLayout)
	ret.EnclosingLayout.AddItem(ret.InnerList, 0, 9, false)
	ret.EnclosingLayout.AddItem(ret.BottomRowLayout, 1, 1, false)
	return &ret
}

func fillFolders(folders types.Allfolders) {
	for folderName := range folders {
		uiPart.flp.InnerList.AddItem(folderName, "", 0, folderSelected)
	}

}

func folderSelected() {
	index := uiPart.flp.InnerList.GetCurrentItem()
	name, _ := uiPart.flp.InnerList.GetItemText(index)
	uiPart.elp.InnerList.Clear()
	fillEntries((*uiPart.data)[name])
	//uiPart.app.SetFocus(uiPart.entriesLayout)
	//uiPart.outerLayout.ResizeItem(uiPart.foldersLayout, 0, 10)
}

func newFolderButton() {

}
func editFolderButton() {

}
func deleteFolderButton() {

}
