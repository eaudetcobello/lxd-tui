package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
)

func loadContainers(client interfaces.InstanceDAO) tea.Cmd {
	return func() tea.Msg {
		containers, err := client.GetInstances(interfaces.InstanceTypeAny, "")
		return InstancesLoadedMsg{instances: containers, err: err}
	}
}

func loadProjects(client interfaces.ProjectDAO) tea.Cmd {
	return func() tea.Msg {
		projects, err := client.GetProjects()
		return ProjectsLoadedMsg{projects: projects, err: err}
	}
}
