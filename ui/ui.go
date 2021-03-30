package ui

import (
	"os"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const btnSize1 = 5

type View struct {
	app        *tview.Application
	tree       *tview.TreeView
	dataRoot   *types.Container
	infoBox    *tview.TextView
	exitButton *tview.Button
	newButton  *tview.Button
	editButton *tview.Button
	mainView   *tview.Flex
	pages      *tview.Pages
	menuForm   *tview.Form
}

func InitView(data *types.Container) *View {
	view := &View{
		app:        tview.NewApplication(),
		tree:       tview.NewTreeView(),
		dataRoot:   data,
		infoBox:    tview.NewTextView(),
		exitButton: tview.NewButton("Exit"),
		newButton:  tview.NewButton("New"),
		editButton: tview.NewButton("Edit"),
		mainView:   tview.NewFlex(),
		pages:      tview.NewPages(),
	}
	view.newButton.SetSelectedFunc(func() { createMenu(view) })
	view.editButton.SetSelectedFunc(func() { editTitleMenu(view) })
	view.exitButton.SetSelectedFunc(func() { stopApp(view) })
	setupTree(view)
	return view
}

func ShowUI(v *View) {
	bottomBar := tview.NewFlex()
	bottomBar.SetDirection(tview.FlexColumn)
	bottomBar.AddItem(v.exitButton, 0, btnSize1, false)
	bottomBar.AddItem(v.newButton, 0, btnSize1, false)
	bottomBar.AddItem(v.editButton, 0, btnSize1, false)
	bottomBar.AddItem(v.infoBox, 0, 90, false)

	//Setup main view
	v.mainView.SetDirection(tview.FlexRow)
	v.mainView.AddItem(v.tree, 0, 99, true)
	v.mainView.AddItem(bottomBar, 1, 1, false)

	//Create pages
	v.pages.AddPage("mainView", v.mainView, true, true)

	v.app.SetRoot(v.pages, true)
	v.app.EnableMouse(true)
}
func createModel(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, false).
			AddItem(nil, 0, 1, false), width, 1, false).
		AddItem(nil, 0, 1, false)
}

func RunUI(v *View) {
	if err := v.app.Run(); err != nil {
		panic(err)
	}
}

func stopApp(v *View) {
	//util button for testing
	//try saving
	v.app.Stop()
	os.Exit(0)
}

//switch to main page
func switchToMain(v *View) {
	//remove any floating menu that exists
	v.pages.RemovePage("modal")
	//app must be stop to change focus back to tree
	v.app.Stop()
	v.app = tview.NewApplication()
	if err := v.app.SetRoot(v.pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	v.pages.SwitchToPage("mainView")
}
