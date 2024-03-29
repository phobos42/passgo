package utilTypes

type PassGo struct {
	Folders *Container
}

type Container struct {
	Title      string
	Entries    []*Entry
	Containers []*Container
}

type Entry struct {
	Title string
	Items []*Item
}
type Item struct {
	Type  string
	Title string
	Value string
}
