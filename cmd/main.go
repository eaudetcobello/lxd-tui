package main

import (
	lxd_dao "github.com/eaudetcobello/lxd-tui/internal/dao/lxd"
	"github.com/eaudetcobello/lxd-tui/internal/ui"
)

func main() {
	apiClient, err := lxd_dao.NewLXDClient("")
	if err != nil {
		panic(err)
	}

	initialModel := ui.InitialModel(*apiClient)

	app := ui.NewApp(initialModel)

	app.Run()
}
