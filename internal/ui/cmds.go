package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
)

func refreshAll(client interfaces.ClientProvider) tea.Cmd {
	return tea.Batch(
		loadContainers(client),
		loadProjects(client),
	)
}

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

func stopInstance(client interfaces.InstanceDAO, instanceName string) tea.Cmd {
	return func() tea.Msg {
		err := client.StopInstance(instanceName)
		return StopInstanceMsg{err: err}
	}
}

func deleteInstance(client interfaces.InstanceDAO, instanceName string) tea.Cmd {
	return func() tea.Msg {
		err := client.DeleteInstance(instanceName, "")

		return DeleteInstanceMsg{instanceName: instanceName, err: err}
	}
}

func deleteAndStopInstance(client interfaces.InstanceDAO, instanceName string) tea.Cmd {
	return tea.Sequence(
		stopInstance(client, instanceName),
		deleteInstance(client, instanceName),
	)
}

func refreshInstances(client interfaces.InstanceDAO) tea.Cmd {
	return func() tea.Msg {
		containers, err := client.GetInstances(interfaces.InstanceTypeAny, "")
		return InstancesLoadedMsg{instances: containers, err: err}
	}
}

func refreshProjects(client interfaces.ProjectDAO) tea.Cmd {
	return func() tea.Msg {
		projects, err := client.GetProjects()
		return ProjectsLoadedMsg{projects: projects, err: err}
	}
}
