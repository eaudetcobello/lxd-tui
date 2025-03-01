package lxd_dao

import (
	"fmt"

	lxd "github.com/canonical/lxd/client"
	"github.com/canonical/lxd/shared/api"
	. "github.com/eaudetcobello/lxd-tui/internal/dao/interfaces"
	. "github.com/eaudetcobello/lxd-tui/internal/logger"
)

type LXDProvider struct {
	Server lxd.InstanceServer
	Logger Logger
}

func NewLXDProvider(server lxd.InstanceServer, logger Logger) *LXDProvider {
	return &LXDProvider{Server: server, Logger: logger}
}

func ConnectLXDUnix(socketPath string, logger Logger) (*LXDProvider, error) {
	server, err := lxd.ConnectLXDUnix(socketPath, nil)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to LXD: %w", err)
	}

	return NewLXDProvider(server, logger), nil
}

func (c LXDProvider) getInstancesByType(instanceType api.InstanceType) ([]Instance, error) {
	apiInstances, err := c.Server.GetInstances(instanceType)
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
	c.Server.UseProject(projectName)
	return c
}

func (c LXDProvider) GetProjects() ([]Project, error) {
	allProjects := make([]Project, 0)
	projects, err := c.Server.GetProjects()
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

func (c LXDProvider) DeleteInstance(instanceName string, projectName string) error {
	op, err := c.Server.DeleteInstance(instanceName)
	if err != nil {
		return fmt.Errorf("Error deleting instance: %w", err)
	}
	err = op.Wait()
	if err != nil {
		return fmt.Errorf("Error waiting for operation: %w", err)
	}
	return nil
}

func (c LXDProvider) GetVMs() ([]Instance, error) {
	return c.getInstancesByType(api.InstanceTypeVM)
}

func (c LXDProvider) StopInstance(name string, s string) error {
	op, err := c.Server.UpdateInstanceState(name, api.InstanceStatePut{Action: "stop"}, "")
	if err != nil {
		return fmt.Errorf("Error stopping instance: %w", err)
	}
	err = op.Wait()
	if err != nil {
		return fmt.Errorf("Error waiting for operation: %w", err)
	}
	return nil
}
