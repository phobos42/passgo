package ui

import (
	"fmt"
	"reflect"

	types "github.com/phobos42/passgo/utils"
	"github.com/rivo/tview"
)

func setupTree() {
	var v = thisView
	v.tree.SetBorder(true).SetTitle("Passgo").SetBorderColor(v.colors.fg4).SetTitleColor(v.colors.fg0)
	root := tview.NewTreeNode(v.dataRoot.Title)
	root.SetColor(v.colors.green)

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
		node.SetColor(thisView.colors.blue)
		for entry := range (*t).Entries {
			fillTree(node, &(*t).Entries[entry])
		}
		for container := range (*t).Containers {
			fillTree(node, &(*t).Containers[container])
		}
	case **types.Entry:
		node = tview.NewTreeNode((*t).Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { entrySelected(node.GetReference(), node) })
		node.SetColor(thisView.colors.yellow)
		for item := range (*t).Items {
			fillTree(node, &(*t).Items[item])
		}
	case **types.Item:
		node = tview.NewTreeNode((*t).Title).SetReference(t).SetExpanded(false)
		node.SetSelectedFunc(func() { itemSelected(node.GetReference()) })
		node.SetColor(thisView.colors.aqua)
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
	//add new reference to tree and data structure
	ref := parentNode.GetReference()
	switch parent := ref.(type) {
	//parent node is a container
	case **types.Container:
		switch newObj := reference.(type) {
		//reference must be either Container or Entry
		case *types.Container:
			(*parent).Containers = append((*parent).Containers, newObj)

			fillTree(parentNode, &newObj)
		case *types.Entry:
			(*parent).Entries = append((*parent).Entries, newObj)
			fillTree(parentNode, &newObj)
		default:
			return "New node has bad type" + reflect.TypeOf(newObj).String()
		}
	//parent is an entry -> reference is an item
	case **types.Entry:
		if _, err := reference.(*types.Item); !err {
			return "Can only create Item inside entry"
		}
		item := reference.(*types.Item)
		(*parent).Items = append((*parent).Items, item)
		fillTree(parentNode, &item)
	default:
		return "Selected Node was not a Container: " + reflect.TypeOf(ref).String()
	}
	return ""
}

//called when a container in the ui tree is selected
func containerSelected(n *tview.TreeNode) {
	n.SetExpanded(!n.IsExpanded())
}

//called when an item in the ui tree is selected
func itemSelected(reference interface{}) {
	switch ref := reference.(type) {
	case **types.Item:
		writeToClipboard((**ref).Value)
	default:
		fmt.Println("Bad Type for Item")
	}

}
func entrySelected(reference interface{}, n *tview.TreeNode) {
	switch reference.(type) {
	case **types.Entry:
		createEntryMenu()
		n.SetExpanded(true)
		//check if node is already expanced
		// if n.IsExpanded() {
		// 	//open floating dialog with item in list
		// 	createEntryMenu()
		// 	n.SetExpanded(false)
		// } else {
		// 	n.SetExpanded(true)
		// }
	default:
		fmt.Println("Bad Type for Entry")
	}

}

func deleteChild(parentNode *tview.TreeNode, targetNode *tview.TreeNode) {
	//only need to remove top level reference to delete all children
	parentNode.RemoveChild(targetNode)
	parent_ref := parentNode.GetReference()
	target_ref := targetNode.GetReference()
	//parent type
	switch p_t := parent_ref.(type) {
	case **types.Container:
		//target type
		switch t_t := target_ref.(type) {
		case **types.Container:
			//remove container from container
			i := indexOfContainer((*p_t).Containers, t_t)
			if i != -1 {
				copy((**p_t).Containers[i:], (**p_t).Containers[i+1:])
				(**p_t).Containers = (**p_t).Containers[:len((**p_t).Containers)-1]
			}

		case **types.Entry:
			//remove entry from container
			i := indexOfEntry((*p_t).Entries, t_t)
			if i != -1 {
				copy((**p_t).Entries[i:], (**p_t).Entries[i+1:])
				(**p_t).Entries = (**p_t).Entries[:len((**p_t).Entries)-1]
			}
		}
	case **types.Entry:
		switch t_t := target_ref.(type) {
		case **types.Item:
			//remove item from entry
			i := indexOfItem((*p_t).Items, t_t)
			if i != -1 {
				copy((**p_t).Items[i:], (**p_t).Items[i+1:])
				(**p_t).Items = (**p_t).Items[:len((**p_t).Items)-1]
			}
		}
	}
}

func indexOfContainer(data []*types.Container, target **types.Container) int {
	for k, v := range data {
		if *target == v {
			return k
		}
	}
	return -1
}
func indexOfEntry(data []*types.Entry, target **types.Entry) int {
	for k, v := range data {
		if *target == v {
			return k
		}
	}
	return -1
}
func indexOfItem(data []*types.Item, target **types.Item) int {
	for k, v := range data {
		if *target == v {
			return k
		}
	}
	return -1
}
