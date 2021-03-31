package ui

import (
	"fmt"
	"reflect"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func setupTree() {
	var v = thisView
	v.tree.SetBorder(true).SetTitle("Passgo")
	root := tview.NewTreeNode(v.dataRoot.Title)
	root.SetSelectedFunc(func() { containerSelected(root) }).SetReference(&v.dataRoot)
	v.tree.SetRoot(root).SetCurrentNode(root)

	for container := range v.dataRoot.Containers {
		fillTree(root, &v.dataRoot.Containers[container])
	}
	//root node is expanded by default
	v.tree.GetRoot().SetExpanded(true)

}

//recursive fill nodes
func fillTree(target *tview.TreeNode, reference interface{}) {

	var node *tview.TreeNode
	switch t := reference.(type) {
	case **types.Container:
		node = tview.NewTreeNode((*t).Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { containerSelected(node) })
		for entry := range (*t).Entries {
			fillTree(node, &(*t).Entries[entry])
		}
		for container := range (*t).Containers {
			fillTree(node, &(*t).Containers[container])
		}
	case **types.Entry:
		node = tview.NewTreeNode((*t).Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { entrySelected(node.GetReference(), node) })
		for item := range (*t).Items {
			fillTree(node, &(*t).Items[item])
		}
	case **types.Item:
		node = tview.NewTreeNode((*t).Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { itemSelected(node.GetReference()) })
	default:
		//unhandled reference type
		// v.infoBox.SetText(reflect.TypeOf(reference).String())
		thisView.infoBox.SetText(reflect.TypeOf(reference).String())
		return
	}
	target.AddChild(node)
}

//adds an item to both the ui tree and underlying data structure(defined in types)
func addItem(parentNode *tview.TreeNode, reference interface{}) string {
	// thisView.infoBox.SetText(reflect.TypeOf(parentNode.GetReference()).String())
	//add new reference to tree and datastructure

	ref := parentNode.GetReference()
	switch parent := ref.(type) {
	case **types.Container:
		//parent node is a container
		switch newObj := reference.(type) {
		//newObj is a new child
		case *types.Container:
			(*parent).Containers = append((*parent).Containers, newObj)

			fillTree(parentNode, &newObj)
		case *types.Entry:
			(*parent).Entries = append((*parent).Entries, newObj)
			fillTree(parentNode, &newObj)
		default:
			return "inner select bad" + reflect.TypeOf(newObj).String()
		}
	default:
		return "Selected Node was not a Container" + reflect.TypeOf(ref).String()
	}
	//thisView.infoBox.SetText(fmt.Sprintf("p:%p c:%p\n", &ref, &reference))
	return ""
}

//called when a container in the ui tree is selected
func containerSelected(n *tview.TreeNode) {
	n.SetExpanded(!n.IsExpanded())
}

//called when an item in the ui tree is selected
func itemSelected(reference interface{}) {
	switch reference.(type) {
	case **types.Item:
		//copy value from item for quick access
	default:
		fmt.Println("Bad Type for Item")
	}

}
func entrySelected(reference interface{}, n *tview.TreeNode) {
	switch reference.(type) {
	case **types.Entry:
		//check if node is already expanced
		if n.IsExpanded() {
			//open floating dialog with item in list
			createEntryMenu()
			n.SetExpanded(false)
		} else {
			n.SetExpanded(true)
		}
	default:
		fmt.Println("Bad Type for Entry")
	}

}
