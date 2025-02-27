package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	domain "github.com/eaudetcobello/lxd-tui/internal/domain"
)

type model struct {
	cursor      int
	selected    map[int]*domain.Resource
	currentType domain.ResourceType
	resources   map[domain.ResourceType][]domain.Resource
}

func InitialModel(containers []domain.Resource, projects []domain.Resource) model {
	return model{
		cursor:      0,
		selected:    make(map[int]*domain.Resource),
		currentType: domain.TypeProject,
		resources: map[domain.ResourceType][]domain.Resource{
			domain.TypeContainer: containers,
			domain.TypeProject:   projects,
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor+1 < len(m.resources[m.currentType]) {
				m.cursor++
			}
		case "k", "up":
			if m.cursor-1 >= 0 {
				m.cursor--
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = &m.resources[m.currentType][m.cursor]
			}
		case "tab":
			switch m.currentType {
			case domain.TypeContainer:
				m.currentType = domain.TypeProject
			case domain.TypeProject:
				m.currentType = domain.TypeContainer
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	for i, choice := range m.resources[m.currentType] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		if choice.GetType() == domain.TypeContainer {
			container := choice.(domain.ContainerResource)
			s += fmt.Sprintf("%s [%s] %s - %s - %s\n", cursor, checked, choice.GetName(), container.GetStatus(), container.GetLastUsedAt())
			continue
		} else if choice.GetType() == domain.TypeProject {
			project := choice.(domain.ProjectResource)
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, project.GetName())
			continue
		}
	}

	s += "\nPress q to quit.\n"

	return s
}
