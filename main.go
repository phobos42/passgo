package main

import (
	types "github.com/phobos42/passgo/utils"
)

type passgo struct {
	ui      userInterface
	folders types.Allfolders
}

var application passgo

func main() {
	ingestJSON()

	// fmt.Println(string(application.folders["folder1"]["place1"][0].Title))
	// exportJSON()
	initUI()
	showUI()
	runUI()

}

// app := tview.NewApplication()
// passwordForm := uiparts.CreateGeneratorForm(app)

//give data to selection layout
// selectionLayout := uiparts.Texpandable()

// masterlayout := tview.NewFlex().
// 	SetDirection(tview.FlexColumn).
// 	//AddItem(selectionLayout, 0, 50, true).
// 	AddItem(selectionLayout, 0, 50, true).
// 	AddItem(passwordForm, 0, 50, false)

// if err := app.SetRoot(masterlayout, true).EnableMouse(true).Run(); err != nil {
// 	panic(err)
// }

// folderlist = tview.NewList()
// folderlist.AddItem("text1", "", 0, selctedfunc)
// folderlist.AddItem("text2", "", 0, selctedfunc)
// folderlist.AddItem("text3", "", 0, selctedfunc)
// folderlist.AddItem("text4", "", 0, selctedfunc)

// testOutput = tview.NewTextView()
// testOutput.SetText("emtpy")

// entries = tview.NewList()
// entries.ShowSecondaryText(false)
// entries.SetBorderPadding(0, 0, 2, 2)
// entries.AddItem("item 1", "", '0', func() {
// 	expandSelection()
// })
// entries.AddItem("item 2", "", '1', func() {})
// cellone := tview.NewTableCell()
// table := tview.NewTable()
// exampleInput := tview.NewInputField()
// exampleButton := tview.NewButton("button")

// examplelayout := tview.NewFlex().
// 	SetDirection(tview.FlexColumn).
// 	AddItem(exampleInput, 0, 80, true).
// 	AddItem(exampleButton, 0, 20, false)

// largerlayout := tview.NewFlex().
// 	SetDirection(tview.FlexRow).
// 	AddItem(folderlist, 0, 5, false).
// 	AddItem(testOutput, 0, 5, false)
