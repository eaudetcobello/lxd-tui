package ui

import "github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"

type InstancesLoadedMsg struct {
	instances []interfaces.Instance
	err       error
}

type ProjectsLoadedMsg struct {
	projects []interfaces.Project
	err      error
}

type RefreshMsg struct {
	err error
}
