package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	model model
}

func NewApp(InitialModel model) *App {
	return &App{
		model: InitialModel,
	}
}

func (a *App) Run() {
	p := tea.NewProgram(a.model)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
