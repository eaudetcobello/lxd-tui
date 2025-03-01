package main

import (
	"fmt"

	lxd_dao "github.com/eaudetcobello/lxd-tui/internal/dao/lxd"
	"github.com/eaudetcobello/lxd-tui/internal/logger"
	"github.com/eaudetcobello/lxd-tui/internal/ui"
)

func main() {
	logger, err := logger.NewFileLogger("lxd-tui.log")
	if err != nil {
		fmt.Printf("Error creating logger: %v", err)
	}

	lxdProvider, err := lxd_dao.ConnectLXDUnix("", logger)
	if err != nil {
		fmt.Printf("Error connecting to LXD: %v", err)
	}

	lxdProvider.Logger.Info("--- LXD TUI started ---")

	initialModel := ui.InitialModel(*lxdProvider)

	app := ui.NewApp(initialModel)

	app.Run()
}
