package main

import (
	"github.com/phobos42/passgo/json"
	"github.com/phobos42/passgo/ui"
	types "github.com/phobos42/passgo/utils"
)

type passgo struct {
	folders *types.Container
}

var application passgo

func main() {
	application.folders = new(types.Container)
	json.IngestJSON(application.folders)
	// exportJSON()

	view := ui.InitView(application.folders)

	ui.ShowUI(view)
	ui.RunUI(view)
}
