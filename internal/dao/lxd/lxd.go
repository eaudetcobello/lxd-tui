package lxd_dao

import (
	"fmt"

	lxd "github.com/canonical/lxd/client"
	"github.com/canonical/lxd/shared/api"
	. "github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
)

type LXDProvider struct {
	server lxd.InstanceServer
}

func NewLXDClient(socketPath string) (*LXDProvider, error) {
	server, err := lxd.ConnectLXDUnix(socketPath, nil)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to LXD: %w", err)
	}
	return &LXDProvider{server: server}, nil
}

func (c LXDProvider) getInstancesByType(instanceType api.InstanceType) ([]Instance, error) {
	apiInstances, err := c.server.GetInstances(instanceType)
	if err != nil {
		return nil, fmt.Errorf("Error getting instances: %w", err)
	}
	domainInstances := make([]Instance, 0)
	for _, instance := range apiInstances {
		domainInstances = append(domainInstances, Instance{Name: instance.Name, Status: instance.Status})
	}
	return domainInstances, nil
}

func (c LXDProvider) UseProject(projectName string) LXDProvider {
	c.server.UseProject(projectName)
	return c
}

func (c LXDProvider) GetProjects() ([]Project, error) {
	allProjects := make([]Project, 0)
	projects, err := c.server.GetProjects()
	if err != nil {
		return nil, fmt.Errorf("Error getting projects: %w", err)
	}
	for _, project := range projects {
		allProjects = append(allProjects, Project{Name: project.Name})
	}
	return allProjects, nil
}

func (c LXDProvider) GetInstances(instanceType InstanceType, projectName string) ([]Instance, error) {
	switch instanceType {
	case InstanceTypeAny:
		return c.getInstancesByType(api.InstanceTypeAny)
	case InstanceTypeContainer:
		return c.getInstancesByType(api.InstanceTypeContainer)
	case InstanceTypeVM:
		return c.getInstancesByType(api.InstanceTypeVM)
	default:
		return nil, fmt.Errorf("Unknown instance type: %s", instanceType)
	}
}

func (c LXDProvider) GetVMs() ([]Instance, error) {
	return c.getInstancesByType(api.InstanceTypeVM)
}
