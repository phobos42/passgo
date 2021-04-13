package ui

import (
	"os"

	"github.com/atotto/clipboard"
	"github.com/phobos42/passgo/json"
	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

const (
	btnSize1 = 5
	btnSize2 = 10
	btnSize3 = 16
)

type View struct {
	app          *tview.Application
	tree         *tview.TreeView
	dataRoot     *types.Container
	infoBox      *tview.TextView
	exitButton   *tview.Button
	newButton    *tview.Button
	editButton   *tview.Button
	deleteButton *tview.Button
	mainView     *tview.Flex
	pages        *tview.Pages
	menuForm     *tview.Form
	menuFlex     *tview.Flex
	colors       *mycolors
}

var currentTarget *tview.TreeNode
var thisView *View

//InitView creates an instance of View
func InitView(data *types.Container) *View {
	view := &View{
		app:          tview.NewApplication(),
		tree:         tview.NewTreeView(),
		dataRoot:     data,
		infoBox:      tview.NewTextView(),
		exitButton:   tview.NewButton("Exit"),
		newButton:    tview.NewButton("New"),
		editButton:   tview.NewButton("Edit Title"),
		deleteButton: tview.NewButton("Delete"),
		mainView:     tview.NewFlex(),
		pages:        tview.NewPages(),
		colors:       newcolors(),
	}
	thisView = view
	view.newButton.SetSelectedFunc(func() { createMenu() })
	colorButton(view.newButton)
	view.editButton.SetSelectedFunc(func() { editTitleMenu() })
	colorButton(view.editButton)
	view.exitButton.SetSelectedFunc(func() { stopApp() })

	view.deleteButton.SetSelectedFunc(func() {
		currentTarget = thisView.tree.GetCurrentNode()
		targetNode := thisView.tree.GetCurrentNode()
		thisView.tree.GetRoot().Walk(findParent)
		//currentTarget gets updated to be parent node by findParent
		if targetNode == thisView.tree.GetRoot() {
			thisView.infoBox.SetText("cannot delete root node")
		} else if currentTarget != targetNode {
			createConfirmMenu(currentTarget, targetNode)
		} else {
			thisView.infoBox.SetText("target to delete not found")
		}
		switchToMain()
	})

	colorButton(view.deleteButton)
	colorButton(view.exitButton)
	setupTree()

	return view
}

//ShowUI is called by main to create the initial tui layout
func ShowUI() {
	var v = thisView
	bottomBar := tview.NewFlex()
	bottomBar.SetDirection(tview.FlexColumn)
	bottomBar.AddItem(v.exitButton, 0, btnSize1, false)
	bottomBar.AddItem(nil, 1, 1, false)
	bottomBar.AddItem(v.newButton, 0, btnSize1, false)
	bottomBar.AddItem(nil, 1, 1, false)
	bottomBar.AddItem(v.editButton, 0, btnSize2, false)
	bottomBar.AddItem(nil, 1, 1, false)
	bottomBar.AddItem(v.deleteButton, 0, btnSize2, false)
	bottomBar.AddItem(nil, 1, 1, false)
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

//RunUI is called by main for the initial TUI startup
func RunUI() {
	var v = thisView
	if err := v.app.Run(); err != nil {
		panic(err)
	}
}

//stopApp will save data and exit the app
func stopApp() {
	var v = thisView
	//util button for testing
	//try saving
	json.ExportJSON(v.dataRoot)
	v.app.Stop()
	os.Exit(0)
}

//switch to switch to main ui page with tree focused
func switchToMain() {
	var v = thisView
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

func writeToClipboard(str string) {
	if err := clipboard.WriteAll(str); err != nil {
		thisView.infoBox.SetText(err.Error())
	}
}
func clearInfoBox() {
	thisView.infoBox.SetText("")
}

func colorButton(btn *tview.Button) {
	btn.SetBackgroundColor(thisView.colors.bg2)
	btn.SetLabelColor(thisView.colors.fg0)
}

func findParent(node *tview.TreeNode, parent *tview.TreeNode) bool {
	//check if currentTarget is a child of node
	if node == currentTarget {
		thisView.infoBox.SetText("Found")
		currentTarget = parent
		return false
	} else {

		return true
	}
}
