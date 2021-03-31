package json

import (
	"encoding/json"
	"io/ioutil"

	types "github.com/phobos42/passgo/utils"
)

//ingest json from file into data structure
func IngestJSON(root *types.Container) {
	var err error
	var jsonB []byte

	jsonB, err = ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	//var myfolders allfolders
	err = json.Unmarshal(jsonB, &root)
	if err != nil {
		panic(err)
	}
}

//export current data structure as json file
func ExportJSON(root *types.Container) {
	var err error
	var jsonB []byte

	jsonB, err = json.MarshalIndent(root, "", "\t")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("test2.json", jsonB, 0644)
	if err != nil {
		panic(err)
	}

}

//creates basic json structure
func createJSON() *types.Container {
	item1 := types.Item{"type", "title", "value"}

	entry1 := types.Entry{"entry title", []*types.Item{&item1, &item1}}

	container1 := types.Container{"container title", []*types.Entry{&entry1, &entry1}, []*types.Container{}}

	return &container1
}
