package uiparts

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func createFieldsPart() *types.FieldsLayoutParts {
	var ret types.FieldsLayoutParts
	ret.ButtonLayout = createButtonLayout(newFieldButton, editFieldButton, deleteFieldButton)
	ret.PromptLayout = createPromptLayout()
	ret.BottomRowLayout = createBottomRowLayout(ret.ButtonLayout, ret.PromptLayout)
	ret.InnerList = createInnerList()

	ret.EnclosingLayout = tview.NewFlex()
	ret.EnclosingLayout.SetDirection(tview.FlexRow)
	setLayoutProperties("Fields", ret.EnclosingLayout)
	ret.EnclosingLayout.AddItem(ret.InnerList, 0, 9, false)
	ret.EnclosingLayout.AddItem(ret.BottomRowLayout, 1, 1, false)
	return &ret
}

func fillFields(fields types.Entry) {
	for index := range fields {
		uiPart.fieldlp.InnerList.AddItem(fields[index].Title, "", 0, nil)
	}
}

func newFieldButton() {

}
func editFieldButton() {

}
func deleteFieldButton() {

}
