package domain

import (
	lxd "github.com/canonical/lxd/shared/api"
)

// Container represents a LXD container.
type ProjectResource struct {
	Project lxd.Project
}

func (c ProjectResource) GetName() string {
	return c.Project.Name
}

func (c ProjectResource) GetDescription() string {
	return c.Project.Description
}

func (c ProjectResource) GetConfig() map[string]string {
	return c.Project.Config
}

func (c ProjectResource) GetType() ResourceType {
	return TypeProject
}

func NewProjectResource(project lxd.Project) ProjectResource {
	return ProjectResource{Project: project}
}
