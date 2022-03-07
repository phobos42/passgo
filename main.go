package main

import (
	"time"

	"github.com/phobos42/passgo/data"
	"github.com/phobos42/passgo/ui"
	utilTypes "github.com/phobos42/passgo/utils"
)

var application utilTypes.PassGo

func main() {
	ui.TheTime = time.Now().UnixNano()
	application.Folders = new(utilTypes.Container)
	data.RetreiveData(application.Folders)
	// exportJSON()

	ui.InitView(application.Folders)

	ui.ShowUI()
	ui.RunUI()
}
