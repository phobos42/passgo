package types

import "github.com/rivo/tview"

//map name of folder to folder
type Allfolders map[string]Folder

//map name of entry to entry
type Folder map[string]Entry

//map index of entry to field
type Entry map[int]Field

// data contained in each field
type Field struct {
	Title string
	Value string
	Hide  bool
}

// folder ui struct
type FoldersLayoutParts struct {
	EnclosingLayout *tview.Flex
	InnerList       *tview.List
	BottomRowLayout *tview.Flex
	ButtonLayout    *tview.Flex
	PromptLayout    *tview.Flex
}

type EntriesLayoutParts struct {
	EnclosingLayout *tview.Flex
	InnerList       *tview.List
	BottomRowLayout *tview.Flex
	ButtonLayout    *tview.Flex
	PromptLayout    *tview.Flex
}
type FieldsLayoutParts struct {
	EnclosingLayout *tview.Flex
	InnerList       *tview.List
	BottomRowLayout *tview.Flex
	ButtonLayout    *tview.Flex
	PromptLayout    *tview.Flex
}
