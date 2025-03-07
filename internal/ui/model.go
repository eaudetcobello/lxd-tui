package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
	"github.com/eaudetcobello/lxd-tui/internal/dao/lxd"
)

type model struct {
	cursor      int
	selected    map[int]*any
	currentView View

	apiClient lxd.LXDProvider

	instances []interfaces.Instance
	projects  []interfaces.Project

	error string
}

func InitialModel(apiClient lxd.LXDProvider) model {
	model := model{
		cursor:      0,
		selected:    make(map[int]*any),
		currentView: ViewInstances,
		apiClient:   apiClient,
		error:       "",
	}
	return model
}

func (m model) getCurrentListLength() int {
	switch m.currentView {
	case ViewInstances:
		return len(m.instances)
	case ViewProjects:
		return len(m.projects)
	}
	return 0
}

func (m model) Init() tea.Cmd {
	return refreshAll(m.apiClient)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

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
		return m, refreshAll(m.apiClient)
	case StopInstanceMsg:
		if msg.err != nil {
			m.error = fmt.Sprintf("Error stopping instance: %v", msg.err)
			return m, nil
		}
		return m, nil
	case DeleteInstanceMsg:
		if msg.err != nil {
			m.error = fmt.Sprintf("Error deleting instance: %v", msg.err)
			return m, nil
		}

		m.cursor = 0

		return m, refreshAll(m.apiClient)

	// Handle key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < m.getCurrentListLength()-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = m.getCurrentListLength() - 1
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
		case "r":
			return m, refreshAll(m.apiClient)
		case "ctrl+d":
			switch m.currentView {
			case ViewInstances:
				if m.instances[m.cursor].Status == "Running" {
					var cmds []tea.Cmd

					for i := range m.selected {
						cmds = append(cmds, deleteAndStopInstance(m.apiClient, m.instances[i].Name))
					}

					m.selected = make(map[int]*any)

					return m, tea.Batch(cmds...)
				}
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
