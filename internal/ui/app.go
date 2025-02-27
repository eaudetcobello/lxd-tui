package ui

import (
	"fmt"

	lxd "github.com/canonical/lxd/client"
	tea "github.com/charmbracelet/bubbletea"
	domain "github.com/eaudetcobello/lxd-tui/internal/domain"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	server, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		fmt.Println("Error running program:", err)
	}

	containers, err := server.UseProject("rockcraft").GetContainers()
	if err != nil {
		fmt.Println("Error running program:", err)
	}

	projects, err := server.GetProjects()
	if err != nil {
		fmt.Println("Error running program:", err)
	}
	projectsResources := make([]domain.Resource, 0)
	for _, project := range projects {
		projectsResources = append(projectsResources, domain.NewProjectResource(project))
	}

	containerResources := make([]domain.Resource, 0)
	for _, container := range containers {
		containerResources = append(containerResources, domain.NewContainerResource(container))
	}

	p := tea.NewProgram(InitialModel(containerResources, projectsResources))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
