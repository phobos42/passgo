package ui

import (
	"fmt"
	"reflect"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func setupTree(v *View) {
	v.tree.SetBorder(true).SetTitle("Passgo")
	root := tview.NewTreeNode(v.dataRoot.Title)
	root.SetSelectedFunc(func() { containerSelected(root) }).SetReference(*v.dataRoot)
	v.tree.SetRoot(root).SetCurrentNode(root)

	for container := range v.dataRoot.Containers {
		fillTree(root, v.dataRoot.Containers[container])
	}
	//root node is expanded by default
	v.tree.GetRoot().SetExpanded(true)

}

//recursive fill nodes
func fillTree(target *tview.TreeNode, reference interface{}) {

	var node *tview.TreeNode
	switch t := reference.(type) {
	case types.Container:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { containerSelected(node) })
		for entry := range t.Entries {
			fillTree(node, t.Entries[entry])
		}
		for container := range t.Containers {
			fillTree(node, t.Containers[container])
		}
	case types.Entry:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { entrySelected(node.GetReference(), node) })
		for item := range t.Items {
			fillTree(node, t.Items[item])
		}
	case types.Item:
		node = tview.NewTreeNode(t.Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { itemSelected(node.GetReference()) })
	default:
		//unhandled reference type
		// v.infoBox.SetText(reflect.TypeOf(reference).String())
		fmt.Println(reflect.TypeOf(reference).String())
		return
	}
	target.AddChild(node)
}

func addItem(parentNode *tview.TreeNode, reference interface{}) string {

	//add new reference to tree and datastructure
	ref := parentNode.GetReference()
	switch t := ref.(type) {
	case types.Container:
		switch obj := reference.(type) {
		case types.Container:
			t.Containers = append(t.Containers, obj)
			fillTree(parentNode, obj)
		case types.Entry:
			t.Entries = append(t.Entries, obj)
			fillTree(parentNode, obj)
		default:
			//cant happen
		}
	default:
		return "Selected Node was not a Container"
	}
	return ""
}

//called when a container is selected
func containerSelected(n *tview.TreeNode) {
	n.SetExpanded(!n.IsExpanded())
}

//called when an item is selected
func itemSelected(reference interface{}) {
	switch reference.(type) {
	case types.Item:
		//copy value from item for quick access
	default:
		fmt.Println("Bad Type for Item")
	}

}
func entrySelected(reference interface{}, n *tview.TreeNode) {
	switch reference.(type) {
	case types.Entry:
		//check if node is already expanced
		if n.IsExpanded() {
			//open floating dialog with item in list
			n.SetExpanded(false)
		} else {
			n.SetExpanded(true)
		}
	default:
		fmt.Println("Bad Type for Entry")
	}

}
