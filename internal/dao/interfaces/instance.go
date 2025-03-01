package interfaces

type Instance struct {
	Name   string
	Status string
}

type InstanceType string

const (
	InstanceTypeContainer InstanceType = "container"
	InstanceTypeVM        InstanceType = "vm"
	InstanceTypeAny       InstanceType = "any"
)

type InstanceDAO interface {
	// GetInstances returns a list of instances
	GetInstances(instanceType InstanceType, projectName string) ([]Instance, error)
	DeleteInstance(instanceName string, projectName string) error
}
