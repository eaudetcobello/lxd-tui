package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
	lxd_dao "github.com/eaudetcobello/lxd-tui/internal/dao/lxd"
)

type model struct {
	cursor      int
	selected    map[int]*any
	currentView View

	apiClient lxd_dao.LXDProvider

	instances []interfaces.Instance
	projects  []interfaces.Project
}

func InitialModel(apiClient lxd_dao.LXDProvider) model {
	model := model{
		cursor:      0,
		selected:    make(map[int]*any),
		currentView: ViewInstances,
		apiClient:   apiClient,
	}
	return model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		loadContainers(m.apiClient),
		loadProjects(m.apiClient),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Handle container and project loading
	case InstancesLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.instances = msg.instances
		return m, nil
	case ProjectsLoadedMsg:
		if msg.err != nil {
			return m, nil
		}
		m.projects = msg.projects
	case RefreshMsg:
		// todo only refresh the current view
		return m, tea.Batch(
			loadContainers(m.apiClient),
			loadProjects(m.apiClient),
		)

	// Handle key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			switch m.currentView {
			case ViewInstances:
				if m.cursor+1 < len(m.instances) {
					m.cursor++
				}
			case ViewProjects:
				if m.cursor+1 < len(m.projects) {
					m.cursor++
				}
			}
		case "k", "up":
			if m.cursor-1 >= 0 {
				m.cursor--
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				var item any
				switch m.currentView {
				case ViewInstances:
					item = m.instances[m.cursor]
				case ViewProjects:
					item = m.projects[m.cursor]
				}
				m.selected[m.cursor] = &item
			}
		case "tab":
			switch m.currentView {
			case ViewInstances:
				m.currentView = ViewProjects
			case ViewProjects:
				m.currentView = ViewInstances
			}
		case "ctrl+d":
			switch m.currentView {
			case ViewInstances:
				if m.instances[m.cursor].Status == "Running" {
					err := m.apiClient.StopInstance(m.instances[m.cursor].Name, "")
					if err != nil {
						m.apiClient.Logger.Error(err, "Error stopping instance")
					}
				}
				go func() {
					err := m.apiClient.DeleteInstance(m.instances[m.cursor].Name, "")
					if err != nil {
						m.apiClient.Logger.Error(err, "Error deleting instance")
					}
				}()
			}

		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	switch m.currentView {
	case ViewInstances:
		s += "Containers\n"
		s += "---------\n"
		for i, container := range m.instances {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s - %s\n", cursor, checked, container.Name, container.Status)
		}

	case ViewProjects:
		s += "Projects\n"
		s += "-------\n"
		for i, project := range m.projects {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, project)
		}
	}

	s += "\nPress q to quit.\n"
	return s
}
