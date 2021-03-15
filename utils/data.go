package types

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

// func addFolder(title string) {
// 	//check if folder exists
// 	if application.folders["title"] != nil {
// 		//bad
// 	}
// 	application.folders[title] = folder{}

// }
// func addEntry(title string) {
// 	//check if folder exists
// 	//check if entry exists in folder
// }
// func addField(title string, value string, hide bool) {
// 	//check if folder exists
// 	//check if entry exists in folder
// 	//add to entries
// }
