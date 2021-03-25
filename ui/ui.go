package ui

import (
	"fmt"
	"reflect"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

type View struct {
	app      *tview.Application
	tree     *tview.TreeView
	dataRoot *types.Container
	infoBox  *tview.TextView
}

func InitView(data *types.Container) *View {
	view := &View{
		app:      tview.NewApplication(),
		tree:     tview.NewTreeView(),
		dataRoot: data,
		infoBox:  tview.NewTextView(),
	}
	setupTree(view)
	return view
}

func setupTree(v *View) {
	v.tree.SetBorder(true).SetTitle("Passgo")
	root := tview.NewTreeNode(v.dataRoot.Title)
	v.tree.SetRoot(root).SetCurrentNode(root).
		SetSelectedFunc(func(n *tview.TreeNode) {
			n.SetExpanded(!n.IsExpanded())
		})
	for container := range v.dataRoot.Containers {
		fillTree(root, v.dataRoot.Containers[container], v)
	}
	v.tree.GetRoot().SetExpanded(true)
	// v.tree.GetRoot().SetExpanded(true)
}

//recursive fill
func fillTree(target *tview.TreeNode, reference interface{}, v *View) {

	var node *tview.TreeNode
	switch t := reference.(type) {
	case types.Container:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)
		for entry := range t.Entries {
			fillTree(node, t.Entries[entry], v)
		}
		for container := range t.Containers {
			fillTree(node, t.Containers[container], v)
		}
		target.AddChild(node)

	case types.Entry:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)

		for item := range t.Items {
			fillTree(node, t.Items[item], v)
		}
		target.AddChild(node)
	case types.Item:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)

		node.SetSelectedFunc(itemSelected)
		target.AddChild(node)
	default:
		v.infoBox.SetText(reflect.TypeOf(reference).String())
		fmt.Println(reflect.TypeOf(reference).String())
		return
	}
}

func itemSelected() {

}

func ShowUI(v *View) {
	mainView := tview.NewFlex()
	mainView.SetDirection(tview.FlexRow)
	mainView.AddItem(v.tree, 0, 99, true)
	mainView.AddItem(v.infoBox, 1, 1, false)
	v.app.SetRoot(mainView, true)
	v.app.EnableMouse(true)
}

func RunUI(v *View) {
	if err := v.app.Run(); err != nil {
		panic(err)
	}
}
