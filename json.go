package main

import (
	"encoding/json"
	"io/ioutil"

	types "github.com/phobos42/passgo/utils"
)

func ingestJSON() {
	var err error
	var jsonB []byte

	jsonB, err = ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	//var myfolders allfolders
	err = json.Unmarshal(jsonB, &application.folders)
	if err != nil {
		panic(err)
	}
}
func exportJSON() {
	var err error
	var jsonB []byte
	jsonB, err = json.MarshalIndent(application.folders, "", "\t")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("testout.json", jsonB, 0644)
	if err != nil {
		panic(err)
	}

}
func createJSON() {
	randomfield := types.Field{"stuff", "myvalue", false}
	randomfield2 := types.Field{"asdf", "pass", true}
	entries := types.Entry{
		0: randomfield,
		1: randomfield2,
	}

	folder1 := types.Folder{
		"place1": entries,
		"place2": entries,
	}
	application.folders = types.Allfolders{
		"folder1": folder1,
		"folder2": folder1,
	}
}
