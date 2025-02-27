package domain

import (
	"time"

	lxd "github.com/canonical/lxd/shared/api"
)

// Container represents a LXD container.
type ContainerResource struct {
	Container lxd.Container
}

func (c ContainerResource) GetName() string {
	return c.Container.Name
}

func (c ContainerResource) GetStatus() string {
	return c.Container.Status
}

func (c ContainerResource) GetLastUsedAt() time.Time {
	return c.Container.LastUsedAt
}

func (c ContainerResource) GetType() ResourceType {
	return TypeContainer
}

func NewContainerResource(container lxd.Container) ContainerResource {
	return ContainerResource{Container: container}
}
