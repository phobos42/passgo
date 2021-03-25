package ui

import (
	"github.com/phobos42/passgo/json"
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
}

func InitView(data *types.Container) *View {
	view := &View{
		app:           tview.NewApplication(),
		tree:          tview.NewTreeView(),
		dataRoot:      data,
		infoBox:       tview.NewTextView(),
		utilityButton: tview.NewButton("util"),
	}
	view.utilityButton.SetSelectedFunc(func() { runUtil(view) })
	setupTree(view)
	return view
}

func ShowUI(v *View) {
	bottomBar := tview.NewFlex()
	bottomBar.SetDirection(tview.FlexColumn)
	bottomBar.AddItem(v.utilityButton, 0, btnSize1, false)
	bottomBar.AddItem(v.infoBox, 0, 90, false)

	mainView := tview.NewFlex()
	mainView.SetDirection(tview.FlexRow)
	mainView.AddItem(v.tree, 0, 99, true)
	mainView.AddItem(bottomBar, 1, 1, false)
	v.app.SetRoot(mainView, true)
	v.app.EnableMouse(true)
}

func RunUI(v *View) {
	if err := v.app.Run(); err != nil {
		panic(err)
	}
}

func runUtil(v *View) {
	//can I get a reference to a single node?

	//need to add new thing to current node.

	//new container or new entry
	newthing := types.Container{Title: "New Container", Entries: []types.Entry{}, Containers: []types.Container{}}
	//newthing := types.Entry{Title: "New Entry", Items: []types.Item{}}

	currentNode := v.tree.GetCurrentNode().SetExpanded(true)
	//add container or entry to selected node.
	v.infoBox.SetText(addItem(currentNode, newthing))

	json.ExportJSON(v.dataRoot)
	v.app.SetFocus(v.tree)
}
