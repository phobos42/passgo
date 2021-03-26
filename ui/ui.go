package ui

import (
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const btnSize1 = 5

type View struct {
	app           *tview.Application
	tree          *tview.TreeView
	dataRoot      *types.Container
	infoBox       *tview.TextView
	utilityButton *tview.Button
	newButton     *tview.Button
	mainView      *tview.Flex
	pages         *tview.Pages
	menuForm      *tview.Form
}

func InitView(data *types.Container) *View {
	view := &View{
		app:           tview.NewApplication(),
		tree:          tview.NewTreeView(),
		dataRoot:      data,
		infoBox:       tview.NewTextView(),
		utilityButton: tview.NewButton("util"),
		newButton:     tview.NewButton("New"),
		mainView:      tview.NewFlex(),
		pages:         tview.NewPages(),
	}
	view.newButton.SetSelectedFunc(func() { createMenu(view) })
	view.utilityButton.SetSelectedFunc(func() { runUtil(view) })
	setupTree(view)
	return view
}

func ShowUI(v *View) {
	bottomBar := tview.NewFlex()
	bottomBar.SetDirection(tview.FlexColumn)
	bottomBar.AddItem(v.utilityButton, 0, btnSize1, false)
	bottomBar.AddItem(v.newButton, 0, btnSize1, false)
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

func RunUI(v *View) {
	if err := v.app.Run(); err != nil {
		panic(err)
	}
}

func runUtil(v *View) {
	//util button for testing
}

//switch to main page
func switchToMain(v *View) {
	v.pages.SwitchToPage("mainView")
}
